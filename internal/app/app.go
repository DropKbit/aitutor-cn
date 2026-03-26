package app

import (
	"fmt"
	"strings"

	"github.com/DropKbit/aitutor-cn/internal/lesson"
	"github.com/DropKbit/aitutor-cn/internal/progress"
	"github.com/DropKbit/aitutor-cn/internal/ui"
	"github.com/DropKbit/aitutor-cn/pkg/types"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// AppModel is the root Bubbletea model.
type AppModel struct {
	width       int
	height      int
	layout      ui.Layout
	header      ui.HeaderModel
	footer      ui.FooterModel
	sidebarOpen bool
	ready       bool
	showWelcome bool
	showHelp    bool
	version     string
	anim        neuralNet

	lessons     []types.LessonDef
	lessonIdx   int
	lessonModel lesson.Model
	sidebar     ui.SidebarModel
	tracker     *progress.Tracker
}

func NewAppModel(version string) AppModel {
	return AppModel{
		header:      ui.NewHeaderModel(),
		footer:      ui.NewFooterModel(),
		sidebarOpen: false,
		showWelcome: true,
		version:     version,
	}
}

func (m AppModel) Init() tea.Cmd {
	return animTick()
}

func (m *AppModel) loadLessons() {
	m.lessons = lesson.All()
	if len(m.lessons) > 0 {
		m.header.Total = len(m.lessons)
		m.sidebar = ui.NewSidebarModel()
		m.sidebar.Lessons = m.lessons
		m.tracker = progress.NewTracker(len(m.lessons))
		m.sidebar.Completed = m.tracker.CompletedMap()

		// Resume from last lesson
		startIdx := m.tracker.LastLessonIdx()
		if startIdx >= len(m.lessons) {
			startIdx = 0
		}
		m.selectLesson(startIdx)
	}
}

func (m *AppModel) selectLesson(idx int) {
	if idx < 0 || idx >= len(m.lessons) {
		return
	}
	m.lessonIdx = idx
	def := m.lessons[idx]
	m.header.Tier = def.Tier
	m.header.LessonTitle = def.Title
	m.header.Current = idx + 1
	m.sidebar.Active = idx
	m.lessonModel = lesson.New(def, m.layout.ContentWidth-2, m.layout.ContentHeight-2)
	m.lessonModel.IsLast = idx == len(m.lessons)-1

	if m.tracker != nil {
		m.tracker.SetLastLesson(idx)
	}
}

func (m *AppModel) markLessonComplete() {
	if m.tracker != nil && m.lessonIdx < len(m.lessons) {
		m.tracker.CompleteLesson(m.lessons[m.lessonIdx].ID)
		m.sidebar.Completed = m.tracker.CompletedMap()
	}
}

func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case animTickMsg:
		if m.showWelcome {
			m.anim.advance()
			return m, animTick()
		}
		return m, nil
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.layout = ui.ComputeLayout(m.width, m.height, m.sidebarOpen)
		m.header.Width = m.width
		m.footer.Width = m.width

		if !m.ready {
			m.ready = true
			m.loadLessons()
		} else {
			lm, cmd := m.lessonModel.Update(tea.WindowSizeMsg{
				Width:  m.layout.ContentWidth - 2,
				Height: m.layout.ContentHeight - 2,
			})
			m.lessonModel = lm
			return m, cmd
		}
		return m, nil

	case tea.KeyMsg:
		// Welcome screen: any key dismisses
		if m.showWelcome {
			if key.Matches(msg, Keys.Quit) {
				return m, tea.Quit
			}
			m.showWelcome = false
			return m, nil
		}

		// Help overlay: any key dismisses
		if m.showHelp {
			m.showHelp = false
			return m, nil
		}

		switch {
		case key.Matches(msg, Keys.Quit):
			return m, tea.Quit
		case key.Matches(msg, Keys.Help):
			m.showHelp = true
			return m, nil
		case key.Matches(msg, Keys.Tab):
			m.sidebarOpen = !m.sidebarOpen
			m.layout = ui.ComputeLayout(m.width, m.height, m.sidebarOpen)
			return m, nil
		case key.Matches(msg, Keys.Next):
			if m.lessonIdx < len(m.lessons)-1 {
				m.selectLesson(m.lessonIdx + 1)
			}
			return m, nil
		case key.Matches(msg, Keys.Prev):
			if m.lessonIdx > 0 {
				m.selectLesson(m.lessonIdx - 1)
			}
			return m, nil
		case key.Matches(msg, Keys.AdvancePhase):
			// Right arrow always advances the phase
			prevPhase := m.lessonModel.Phase
			m.lessonModel.Advance()
			if prevPhase != lesson.PhaseComplete && m.lessonModel.Phase == lesson.PhaseComplete {
				m.markLessonComplete()
			}
			return m, nil
		case key.Matches(msg, Keys.Advance):
			phase := m.lessonModel.Phase
			if phase == lesson.PhaseTheory {
				m.lessonModel.Advance()
				return m, nil
			}
			if phase == lesson.PhaseComplete {
				// Already complete, do nothing on Enter
				return m, nil
			}
			// Fall through to forward to lesson model (viz/quiz)
		case key.Matches(msg, Keys.Back):
			m.lessonModel.GoBack()
			return m, nil
		}
	}

	// Forward to lesson model (handles viz/quiz interactions)
	prevPhase := m.lessonModel.Phase
	var cmd tea.Cmd
	m.lessonModel, cmd = m.lessonModel.Update(msg)

	// Check if lesson just completed
	if prevPhase != lesson.PhaseComplete && m.lessonModel.Phase == lesson.PhaseComplete {
		m.markLessonComplete()
	}

	return m, cmd
}

func (m AppModel) View() string {
	if !m.ready {
		return "正在初始化..."
	}

	if m.showWelcome {
		return m.viewWelcome()
	}

	if m.showHelp {
		return m.viewHelp()
	}

	// Show course completion screen when all lessons done and on last lesson's complete phase
	if m.tracker != nil && m.tracker.CompletedCount() >= len(m.lessons) &&
		m.lessonModel.Phase == lesson.PhaseComplete {
		return m.viewCourseComplete()
	}

	// Update footer hints based on lesson phase
	switch m.lessonModel.Phase {
	case lesson.PhaseTheory:
		m.footer.Bindings = []ui.KeyHint{
			{Key: "q", Desc: "退出"}, {Key: "Tab", Desc: "侧边栏"}, {Key: "n/p", Desc: "上一课/下一课"},
			{Key: "→/Enter", Desc: "下一阶段"}, {Key: "↑/↓", Desc: "滚动"}, {Key: "?", Desc: "帮助"},
		}
	case lesson.PhaseViz:
		m.footer.Bindings = []ui.KeyHint{
			{Key: "q", Desc: "退出"}, {Key: "Tab", Desc: "侧边栏"}, {Key: "n/p", Desc: "上一课/下一课"},
			{Key: "←/→", Desc: "上一阶段/下一阶段"}, {Key: "Enter/Space", Desc: "交互"}, {Key: "?", Desc: "帮助"},
		}
	case lesson.PhaseQuiz:
		m.footer.Bindings = []ui.KeyHint{
			{Key: "q", Desc: "退出"}, {Key: "Tab", Desc: "侧边栏"}, {Key: "n/p", Desc: "上一课/下一课"},
			{Key: "←", Desc: "上一阶段"}, {Key: "1-4", Desc: "答题"}, {Key: "?", Desc: "帮助"},
		}
	case lesson.PhaseComplete:
		m.footer.Bindings = []ui.KeyHint{
			{Key: "q", Desc: "退出"}, {Key: "Tab", Desc: "侧边栏"}, {Key: "n", Desc: "下一课"},
			{Key: "←", Desc: "上一阶段"}, {Key: "?", Desc: "帮助"},
		}
	}

	// Progress bar in header
	completedCount := 0
	if m.tracker != nil {
		completedCount = m.tracker.CompletedCount()
	}
	progressStr := progress.Bar(completedCount, len(m.lessons), 30)

	header := m.header.ViewWithProgress(progressStr)

	// Content
	contentWidth := m.layout.ContentWidth
	contentHeight := m.layout.ContentHeight
	content := ui.ContentStyle.
		Width(contentWidth).
		Height(contentHeight).
		Render(m.lessonModel.View())

	// Sidebar
	var body string
	if m.layout.SidebarOpen {
		m.sidebar.Width = m.layout.SidebarWidth
		m.sidebar.Height = contentHeight
		sidebar := m.sidebar.View()
		body = lipgloss.JoinHorizontal(lipgloss.Top, sidebar, content)
	} else {
		body = content
	}

	footer := m.footer.View()

	return lipgloss.JoinVertical(lipgloss.Left, header, body, footer)
}

func (m AppModel) viewWelcome() string {
	accent := lipgloss.NewStyle().Foreground(ui.ColorAccent).Bold(true)
	bright := lipgloss.NewStyle().Foreground(ui.ColorBright)
	dim := lipgloss.NewStyle().Foreground(ui.ColorMuted)
	green := lipgloss.NewStyle().Foreground(ui.ColorBeginner)
	yellow := lipgloss.NewStyle().Foreground(ui.ColorIntermediate)
	red := lipgloss.NewStyle().Foreground(ui.ColorAdvanced)

	logo := accent.Render("AITutor-ZH")

	var lines []string
	// Only show animation if terminal is tall enough (animation adds ~8 lines)
	if m.height >= 35 {
		lines = append(lines, m.anim.View())
	}
	lines = append(lines, logo)
	lines = append(lines, "")
	tagline := []rune("AI 编程概念交互式中文教程")
	visibleLen := m.anim.frame * 2
	if visibleLen > len(tagline) {
		visibleLen = len(tagline)
	}
	lines = append(lines, bright.Render("  "+string(tagline[:visibleLen])))
	lines = append(lines, "")
	lines = append(lines, dim.Render("  通过动手课程学习 AI 辅助开发的核心概念。"))
	lines = append(lines, dim.Render("  每节课都包含理论、交互式可视化和测验。"))
	lines = append(lines, "")
	lines = append(lines, fmt.Sprintf("  %s  初级   %s",
		green.Render("*"), dim.Render("上下文窗口、工具、提示词")))
	lines = append(lines, fmt.Sprintf("  %s  中级   %s",
		yellow.Render("*"), dim.Render("AGENTS.md、Hook、记忆、模式")))
	lines = append(lines, fmt.Sprintf("  %s  高级   %s",
		red.Render("*"), dim.Render("MCP、技能、子代理、worktree")))
	lines = append(lines, "")

	completedCount := 0
	if m.tracker != nil {
		completedCount = m.tracker.CompletedCount()
	}
	if completedCount > 0 {
		lines = append(lines, green.Render(fmt.Sprintf("  学习进度：已完成 %d/%d 课", completedCount, len(m.lessons))))
		lines = append(lines, "")
	}

	lines = append(lines, accent.Render("  按任意键开始"))
	lines = append(lines, dim.Render("  按 q 退出"))
	lines = append(lines, "")
	lines = append(lines, dim.Render("  "+m.version))
	lines = append(lines, "")
	lines = append(lines, dim.Render("  项目地址 → github.com/DropKbit/aitutor-cn"))
	lines = append(lines, "")
	lines = append(lines, dim.Render("  内容由社区贡献，部分可能经过 AI 辅助整理。"))
	lines = append(lines, dim.Render("  内容可能存在疏漏，不能替代专业培训。"))
	lines = append(lines, dim.Render("  欢迎提交修正与补充。"))

	content := strings.Join(lines, "\n")

	return lipgloss.Place(m.width, m.height,
		lipgloss.Center, lipgloss.Center,
		content)
}

func (m AppModel) viewHelp() string {
	accent := lipgloss.NewStyle().Foreground(ui.ColorAccent).Bold(true)
	bright := lipgloss.NewStyle().Foreground(ui.ColorBright).Bold(true)
	dim := lipgloss.NewStyle().Foreground(ui.ColorMuted)
	keyStyle := lipgloss.NewStyle().Foreground(ui.ColorHighlight).Bold(true).Width(16)

	var lines []string
	lines = append(lines, accent.Render("  帮助"))
	lines = append(lines, "")
	lines = append(lines, bright.Render("  导航"))
	lines = append(lines, fmt.Sprintf("  %s %s", keyStyle.Render("Tab"), dim.Render("切换侧边栏")))
	lines = append(lines, fmt.Sprintf("  %s %s", keyStyle.Render("n / p"), dim.Render("下一课 / 上一课")))
	lines = append(lines, fmt.Sprintf("  %s %s", keyStyle.Render("Up/Down  j/k"), dim.Render("滚动 / 导航")))
	lines = append(lines, fmt.Sprintf("  %s %s", keyStyle.Render("q  Ctrl+C"), dim.Render("退出")))
	lines = append(lines, "")
	lines = append(lines, bright.Render("  课程阶段"))
	lines = append(lines, fmt.Sprintf("  %s %s", keyStyle.Render("→  / Enter"), dim.Render("进入下一阶段")))
	lines = append(lines, fmt.Sprintf("  %s %s", keyStyle.Render("←  / Bksp"), dim.Render("返回上一阶段")))
	lines = append(lines, "")
	lines = append(lines, bright.Render("  可视化"))
	lines = append(lines, fmt.Sprintf("  %s %s", keyStyle.Render("Enter / Space"), dim.Render("与可视化交互")))
	lines = append(lines, fmt.Sprintf("  %s %s", keyStyle.Render("r"), dim.Render("重置可视化")))
	lines = append(lines, "")
	lines = append(lines, bright.Render("  测验"))
	lines = append(lines, fmt.Sprintf("  %s %s", keyStyle.Render("1-4"), dim.Render("选择答案（选择题）")))
	lines = append(lines, fmt.Sprintf("  %s %s", keyStyle.Render("Enter"), dim.Render("提交答案")))
	lines = append(lines, "")
	lines = append(lines, bright.Render("  每节课流程：理论 -> 可视化 -> 测验"))
	lines = append(lines, "")
	lines = append(lines, dim.Render("  按任意键关闭"))

	content := strings.Join(lines, "\n")

	box := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(ui.ColorAccent).
		Padding(1, 2).
		Render(content)

	return lipgloss.Place(m.width, m.height,
		lipgloss.Center, lipgloss.Center,
		box)
}

func (m AppModel) viewCourseComplete() string {
	accent := lipgloss.NewStyle().Foreground(ui.ColorAccent).Bold(true)
	bright := lipgloss.NewStyle().Foreground(ui.ColorBright).Bold(true)
	dim := lipgloss.NewStyle().Foreground(ui.ColorMuted)
	green := lipgloss.NewStyle().Foreground(ui.ColorBeginner).Bold(true)
	link := lipgloss.NewStyle().Foreground(ui.ColorHighlight).Underline(true)

	var lines []string
	lines = append(lines, "")
	lines = append(lines, green.Render("  恭喜你！"))
	lines = append(lines, "")
	lines = append(lines, bright.Render(fmt.Sprintf("  你已经完成全部 %d 节课程。", len(m.lessons))))
	lines = append(lines, "")
	lines = append(lines, dim.Render("  现在你已经掌握 AI 辅助开发中的核心概念："))
	lines = append(lines, dim.Render("  上下文窗口、工具、MCP、子代理、"))
	lines = append(lines, dim.Render("  批量执行等关键主题。"))
	lines = append(lines, "")
	lines = append(lines, accent.Render("  ── 接下来做什么？ ──"))
	lines = append(lines, "")
	lines = append(lines, dim.Render("  把这些概念用到真实项目中。尝试在自己的"))
	lines = append(lines, dim.Render("  工作流里使用 AI 编程助手，观察这些模式"))
	lines = append(lines, dim.Render("  如何在实际开发中发挥作用。"))
	lines = append(lines, "")
	lines = append(lines, accent.Render("  ── 贡献 ──"))
	lines = append(lines, "")
	lines = append(lines, dim.Render("  如果你发现缺漏或错误，欢迎一起完善。"))
	lines = append(lines, dim.Render("  提交 issue 或 PR："))
	lines = append(lines, "")
	lines = append(lines, "  "+link.Render("github.com/DropKbit/aitutor-cn"))
	lines = append(lines, "")
	lines = append(lines, dim.Render("  无论是新课程想法、缺陷修复，还是更好的"))
	lines = append(lines, dim.Render("  解释方式，都欢迎贡献。"))
	lines = append(lines, "")
	lines = append(lines, dim.Render("  按 p 回顾课程  |  按 q 退出"))

	content := strings.Join(lines, "\n")

	box := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(ui.ColorBeginner).
		Padding(1, 2).
		Render(content)

	return lipgloss.Place(m.width, m.height,
		lipgloss.Center, lipgloss.Center,
		box)
}
