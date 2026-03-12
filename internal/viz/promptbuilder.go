package viz

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/naorpeled/aitutor/internal/ui"
)

type techniqueChallenge struct {
	Scenario    string
	BadPrompt   string
	Options     []string
	CorrectIdx  int
	Technique   string
	Explanation string
}

// PromptBuilderModel is an interactive exercise matching coding scenarios to prompting techniques.
type PromptBuilderModel struct {
	width      int
	height     int
	challenges []techniqueChallenge
	current    int
	cursor     int
	answered   bool
	correct    bool
	score      int
}

func NewPromptBuilderModel(w, h int) Model {
	return &PromptBuilderModel{
		width:  w,
		height: h,
		challenges: []techniqueChallenge{
			{
				Scenario:  "You need the AI to write error handling that matches your project's style",
				BadPrompt: "\"Handle errors properly\"",
				Options: []string{
					"Give it a persona: \"You are a senior developer\"",
					"Show a negative example: \"DON'T: log.Fatal(err). DO: return fmt.Errorf(...)\"",
					"Ask it to think step by step",
					"Provide the function signature first",
				},
				CorrectIdx:  1,
				Technique:   "Negative Examples",
				Explanation: "Showing what you DON'T want alongside what you DO want is the most reliable way to control coding style. The AI matches the pattern you demonstrate.",
			},
			{
				Scenario:  "You want the AI to implement a complex algorithm you've already specified with tests",
				BadPrompt: "\"Implement a URL slug function\"",
				Options: []string{
					"Assign it a persona as an algorithm expert",
					"List all the edge cases to handle",
					"Provide the test cases first, then ask for code that passes them",
					"Ask it to scaffold the code first",
				},
				CorrectIdx:  2,
				Technique:   "Test-Driven Prompting",
				Explanation: "Providing tests first creates an executable specification with zero ambiguity. The AI generates code that satisfies the exact behavior you defined.",
			},
			{
				Scenario:  "You need the AI to design a caching system and you want it to evaluate trade-offs",
				BadPrompt: "\"Build a cache\"",
				Options: []string{
					"Show negative examples of bad caching code",
					"Provide test cases for the cache",
					"Give it a persona: \"You are a senior systems engineer specializing in distributed caching\"",
					"Provide the function signature upfront",
				},
				CorrectIdx:  2,
				Technique:   "Persona / Role Assignment",
				Explanation: "Setting a specific expert role focuses the AI's reasoning. A \"senior systems engineer\" will naturally consider trade-offs, concurrency, and eviction strategies that a generic prompt would miss.",
			},
			{
				Scenario:  "You want to add 10 similar API endpoints that follow the same pattern as existing ones",
				BadPrompt: "\"Add CRUD endpoints for Orders, Products, and Reviews\"",
				Options: []string{
					"Ask it to think through the design first",
					"Paste 2 existing endpoints as examples and say \"follow this pattern\"",
					"Assign it an expert persona",
					"Provide test cases for each endpoint",
				},
				CorrectIdx:  1,
				Technique:   "Few-Shot Examples",
				Explanation: "Showing 2-3 examples of existing code is the most reliable way to match your project's patterns. The AI mirrors your naming, error handling, and structure exactly.",
			},
			{
				Scenario:  "You need the AI to design a complex migration strategy before writing any code",
				BadPrompt: "\"Migrate our database schema\"",
				Options: []string{
					"Show examples of past migrations",
					"Provide test cases for the migration",
					"Ask it to reason through the approach: \"Compare strategies, evaluate risks, then implement\"",
					"Show negative examples of bad migrations",
				},
				CorrectIdx:  2,
				Technique:   "Chain-of-Thought",
				Explanation: "Asking the AI to reason step-by-step before coding produces better architectural decisions. It evaluates trade-offs and catches issues before writing a single line.",
			},
			{
				Scenario:  "You need to choose between 3 possible architectures for a real-time notification system",
				BadPrompt: "\"Build a notification system\"",
				Options: []string{
					"Provide test cases for the notification system",
					"Show existing notification code as examples",
					"Ask 3 imaginary experts to each propose an approach, critique each other, then pick the best",
					"Assign it a persona as a notifications specialist",
				},
				CorrectIdx:  2,
				Technique:   "Tree-of-Thought",
				Explanation: "Tree-of-thought explores multiple approaches in parallel, evaluates each, and picks the strongest. Perfect when there's no obvious right architecture.",
			},
			{
				Scenario:  "You need a large feature but want to review it in small, manageable pieces",
				BadPrompt: "\"Build the entire payment system\"",
				Options: []string{
					"Assign it an expert persona",
					"Show examples of similar systems",
					"Ask for type signatures and stubs first, then implement one function at a time",
					"Provide all the test cases upfront",
				},
				CorrectIdx:  2,
				Technique:   "Scaffold-Then-Detail",
				Explanation: "Asking for the skeleton first (types, signatures, doc comments) then filling in one function per prompt keeps each response focused and reviewable.",
			},
		},
	}
}

func (m *PromptBuilderModel) Init() tea.Cmd { return nil }

func (m *PromptBuilderModel) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.answered {
			if key.Matches(msg, key.NewBinding(key.WithKeys("enter", " "))) {
				m.current++
				m.cursor = 0
				m.answered = false
				if m.current >= len(m.challenges) {
					m.current = len(m.challenges)
				}
			}
			if key.Matches(msg, key.NewBinding(key.WithKeys("r"))) {
				m.reset()
			}
			return m, nil
		}

		if m.current >= len(m.challenges) {
			if key.Matches(msg, key.NewBinding(key.WithKeys("r"))) {
				m.reset()
			}
			return m, nil
		}

		switch {
		case key.Matches(msg, key.NewBinding(key.WithKeys("up", "k"))):
			if m.cursor > 0 {
				m.cursor--
			}
		case key.Matches(msg, key.NewBinding(key.WithKeys("down", "j"))):
			c := m.challenges[m.current]
			if m.cursor < len(c.Options)-1 {
				m.cursor++
			}
		case key.Matches(msg, key.NewBinding(key.WithKeys("1"))):
			m.cursor = 0
			m.submit()
		case key.Matches(msg, key.NewBinding(key.WithKeys("2"))):
			m.cursor = 1
			m.submit()
		case key.Matches(msg, key.NewBinding(key.WithKeys("3"))):
			m.cursor = 2
			m.submit()
		case key.Matches(msg, key.NewBinding(key.WithKeys("4"))):
			if len(m.challenges[m.current].Options) > 3 {
				m.cursor = 3
				m.submit()
			}
		case key.Matches(msg, key.NewBinding(key.WithKeys("enter", " "))):
			m.submit()
		case key.Matches(msg, key.NewBinding(key.WithKeys("r"))):
			m.reset()
		}
	}
	return m, nil
}

func (m *PromptBuilderModel) reset() {
	m.current = 0
	m.cursor = 0
	m.answered = false
	m.score = 0
}

func (m *PromptBuilderModel) submit() {
	c := m.challenges[m.current]
	m.answered = true
	m.correct = m.cursor == c.CorrectIdx
	if m.correct {
		m.score++
	}
}

func (m *PromptBuilderModel) View() string {
	accent := lipgloss.NewStyle().Foreground(ui.ColorAccent).Bold(true)
	bad := lipgloss.NewStyle().Foreground(ui.ColorIncorrect).Bold(true)
	good := lipgloss.NewStyle().Foreground(ui.ColorCorrect).Bold(true)
	highlight := lipgloss.NewStyle().Foreground(ui.ColorHighlight).Bold(true)
	dim := lipgloss.NewStyle().Foreground(ui.ColorMuted)
	text := lipgloss.NewStyle().Foreground(lipgloss.Color("#d1d5db"))
	explain := lipgloss.NewStyle().Foreground(ui.ColorIntermediate)

	var lines []string
	lines = append(lines, "")
	lines = append(lines, accent.Render("  Technique Matcher — Pick the Best Prompting Strategy"))
	lines = append(lines, dim.Render("  For each scenario, choose the most effective technique"))
	lines = append(lines, "")

	if m.current >= len(m.challenges) {
		lines = append(lines, good.Render(fmt.Sprintf("  Exercise Complete! Score: %d/%d", m.score, len(m.challenges))))
		lines = append(lines, "")
		if m.score == len(m.challenges) {
			lines = append(lines, good.Render("  Perfect! You can match techniques to scenarios."))
		} else if m.score >= 4 {
			lines = append(lines, text.Render("  Good work! Remember: persona for expertise, examples for patterns, tests for specs."))
		} else {
			lines = append(lines, text.Render("  Keep practicing — each technique shines in different situations."))
		}
		lines = append(lines, "", dim.Render("  [r] Try again"))
		return strings.Join(lines, "\n")
	}

	c := m.challenges[m.current]
	lines = append(lines, dim.Render(fmt.Sprintf("  Challenge %d of %d", m.current+1, len(m.challenges))))
	lines = append(lines, "")
	lines = append(lines, text.Render("  Scenario: "+c.Scenario))
	lines = append(lines, bad.Render("  Naive prompt: "+c.BadPrompt))
	lines = append(lines, "")
	lines = append(lines, highlight.Render("  Which technique would improve this most?"))
	lines = append(lines, "")

	for i, opt := range c.Options {
		prefix := fmt.Sprintf("  %d) ", i+1)
		style := text

		if m.answered {
			if i == c.CorrectIdx {
				prefix = good.Render(fmt.Sprintf("  %d) ✓ ", i+1))
				style = good
			} else if i == m.cursor && !m.correct {
				prefix = bad.Render(fmt.Sprintf("  %d) ✗ ", i+1))
				style = bad
			}
		} else if i == m.cursor {
			prefix = highlight.Render(fmt.Sprintf("  %d) ▸ ", i+1))
			style = highlight
		}

		wrapped := opt
		if len(wrapped) > m.width-10 {
			wrapped = wrapped[:m.width-13] + "..."
		}
		lines = append(lines, prefix+style.Render(wrapped))
	}

	if m.answered {
		lines = append(lines, "")
		if m.correct {
			lines = append(lines, good.Render("  ✓ Correct! — ")+accent.Render(c.Technique))
		} else {
			lines = append(lines, bad.Render("  ✗ Not quite — ")+accent.Render("Best technique: "+c.Technique))
		}
		lines = append(lines, explain.Render("  "+c.Explanation))
		lines = append(lines, "", highlight.Render("  Press Enter to continue"))
	}

	lines = append(lines, "", dim.Render("  [↑/↓] Navigate  [1-4] Select  [Enter] Confirm  [r] Restart"))

	return strings.Join(lines, "\n")
}
