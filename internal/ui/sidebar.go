package ui

import (
	"fmt"

	"github.com/DropKbit/aitutor-cn/pkg/types"
	"github.com/charmbracelet/lipgloss"
)

// SidebarModel renders the lesson list panel.
type SidebarModel struct {
	Width     int
	Height    int
	Lessons   []types.LessonDef
	Active    int
	Completed map[int]bool
}

func NewSidebarModel() SidebarModel {
	return SidebarModel{
		Completed: make(map[int]bool),
	}
}

func (s SidebarModel) View() string {
	var items []string

	currentTier := types.Tier(-1)
	for i, l := range s.Lessons {
		// Add tier header when tier changes
		if l.Tier != currentTier {
			currentTier = l.Tier
			tierColor := TierColor(int(currentTier))
			header := lipgloss.NewStyle().
				Bold(true).
				Foreground(tierColor).
				MarginTop(1).
				Render(fmt.Sprintf("── %s ──", currentTier))
			items = append(items, header)
		}

		style := SidebarItemStyle
		prefix := "  "
		suffix := ""

		if i == s.Active {
			style = SidebarActiveStyle
			prefix = "▸ "
		}

		if s.Completed[l.ID] {
			suffix = " ✓"
			if i != s.Active {
				style = SidebarCompletedStyle
			}
		}

		title := l.Title
		maxLen := s.Width - 6
		if maxLen > 0 && lipgloss.Width(title) > maxLen {
			title = truncateWidth(title, maxLen-1) + "…"
		}

		items = append(items, style.Render(fmt.Sprintf("%s%d. %s%s", prefix, l.ID, title, suffix)))
	}

	content := lipgloss.JoinVertical(lipgloss.Left, items...)
	return SidebarStyle.
		Width(s.Width).
		Height(s.Height).
		Render(content)
}

func truncateWidth(s string, maxWidth int) string {
	if maxWidth <= 0 {
		return ""
	}

	var out []rune
	width := 0
	for _, r := range s {
		rw := lipgloss.Width(string(r))
		if width+rw > maxWidth {
			break
		}
		out = append(out, r)
		width += rw
	}

	return string(out)
}
