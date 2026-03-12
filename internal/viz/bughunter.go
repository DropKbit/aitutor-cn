package viz

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/naorpeled/aitutor/internal/ui"
)

type bugLang int

const (
	langGo bugLang = iota
	langPython
	langJS
)

var bugLangNames = [3]string{"Go", "Python", "JavaScript"}

type bugVariant struct {
	Code        string
	Options     []string
	CorrectIdx  int
	Explanation string
}

type bugChallengeML struct {
	Title    string
	Variants [3]bugVariant // indexed by bugLang
}

// BugHunterModel is an interactive code review exercise where users spot bugs.
type BugHunterModel struct {
	width      int
	height     int
	challenges []bugChallengeML
	lang       bugLang
	picking    bool // true while choosing language
	langCursor int
	current    int
	cursor     int
	answered   bool
	correct    bool
	score      int
}

func bugChallenges() []bugChallengeML {
	return []bugChallengeML{
		{
			Title: "Find the off-by-one error",
			Variants: [3]bugVariant{
				{ // Go
					Code: "func getLastN(items []string, n int) []string {\n" +
						"    start := len(items) - n\n" +
						"    return items[start:len(items)]\n" +
						"}",
					Options:     []string{"No error — code is correct", "Should be items[start:len(items)-1]", "Missing bounds check: start could be negative", "Should use append instead of slicing"},
					CorrectIdx:  2,
					Explanation: "If n > len(items), start becomes negative, causing a panic. AI often generates slice code without bounds checking.",
				},
				{ // Python
					Code: "def get_last_n(items: list, n: int) -> list:\n" +
						"    start = len(items) - n\n" +
						"    return items[start:]",
					Options:     []string{"No error — code is correct", "Should be items[start:-1]", "Missing bounds check: start could be negative", "Should use items.copy() first"},
					CorrectIdx:  2,
					Explanation: "If n > len(items), start becomes negative and Python silently returns wrong results (wraps around). AI often skips bounds checking.",
				},
				{ // JS
					Code: "function getLastN(items, n) {\n" +
						"    const start = items.length - n;\n" +
						"    return items.slice(start);\n" +
						"}",
					Options:     []string{"No error — code is correct", "Should use splice instead of slice", "Missing bounds check: start could be negative", "Should use Array.from() first"},
					CorrectIdx:  2,
					Explanation: "If n > items.length, start is negative. While JS slice handles negatives, it changes semantics silently. AI often skips validation.",
				},
			},
		},
		{
			Title: "Spot the non-existent method",
			Variants: [3]bugVariant{
				{ // Go
					Code: "func containsAny(s string, chars []string) bool {\n" +
						"    for _, c := range chars {\n" +
						"        if strings.Includes(s, c) {\n" +
						"            return true\n" +
						"        }\n" +
						"    }\n" +
						"    return false\n" +
						"}",
					Options:     []string{"The loop variable should be index not value", "strings.Includes doesn't exist — it's strings.Contains", "Should return false inside the loop", "The function signature is wrong"},
					CorrectIdx:  1,
					Explanation: "strings.Includes doesn't exist in Go — the correct method is strings.Contains. AI frequently hallucinate plausible-sounding API names.",
				},
				{ // Python
					Code: "def contains_any(s: str, chars: list) -> bool:\n" +
						"    for c in chars:\n" +
						"        if s.contains(c):\n" +
						"            return True\n" +
						"    return False",
					Options:     []string{"The loop should use enumerate()", "s.contains() doesn't exist — use 'c in s' or s.__contains__(c)", "Should return False inside the loop", "The type hints are wrong"},
					CorrectIdx:  1,
					Explanation: "Python strings don't have a .contains() method. Use the 'in' operator: 'c in s'. AI mixes up APIs from different languages.",
				},
				{ // JS
					Code: "function containsAny(s, chars) {\n" +
						"    for (const c of chars) {\n" +
						"        if (s.contains(c)) {\n" +
						"            return true;\n" +
						"        }\n" +
						"    }\n" +
						"    return false;\n" +
						"}",
					Options:     []string{"Should use for...in instead of for...of", "s.contains() doesn't exist — it's s.includes()", "Should return false inside the loop", "const should be let"},
					CorrectIdx:  1,
					Explanation: "JavaScript strings don't have .contains() — the correct method is .includes(). AI often confuses method names across languages.",
				},
			},
		},
		{
			Title: "Find the error handling bug",
			Variants: [3]bugVariant{
				{ // Go
					Code: "func readConfig(path string) *Config {\n" +
						"    data, _ := os.ReadFile(path)\n" +
						"    var cfg Config\n" +
						"    json.Unmarshal(data, &cfg)\n" +
						"    return &cfg\n" +
						"}",
					Options:     []string{"Config should be returned by value not pointer", "json.Unmarshal needs a pointer to pointer", "Errors from ReadFile and Unmarshal are silently ignored", "The function should take io.Reader instead"},
					CorrectIdx:  2,
					Explanation: "Both os.ReadFile and json.Unmarshal return errors that are silently discarded. If the file doesn't exist, you get a zero-value Config with no indication of failure.",
				},
				{ // Python
					Code: "def read_config(path: str) -> dict:\n" +
						"    with open(path) as f:\n" +
						"        data = f.read()\n" +
						"    config = json.loads(data)\n" +
						"    return config",
					Options:     []string{"Should use pathlib instead of open()", "json.loads should be json.load(f)", "No try/except — FileNotFoundError and JSONDecodeError are unhandled", "Should return a dataclass not dict"},
					CorrectIdx:  2,
					Explanation: "If the file doesn't exist or contains invalid JSON, the function crashes with an unhandled exception. AI often generates the happy path without error handling.",
				},
				{ // JS
					Code: "function readConfig(path) {\n" +
						"    const data = fs.readFileSync(path);\n" +
						"    const config = JSON.parse(data);\n" +
						"    return config;\n" +
						"}",
					Options:     []string{"Should use async fs.readFile instead", "JSON.parse should be JSON.stringify", "No try/catch — file read and JSON parse errors are unhandled", "Should return a Promise"},
					CorrectIdx:  2,
					Explanation: "If the file doesn't exist or contains invalid JSON, the function throws an unhandled error. AI often writes the happy path without try/catch.",
				},
			},
		},
		{
			Title: "Spot the concurrency bug",
			Variants: [3]bugVariant{
				{ // Go
					Code: "var count int\n" +
						"\n" +
						"func handleRequest(w http.ResponseWriter, r *http.Request) {\n" +
						"    count++\n" +
						"    fmt.Fprintf(w, \"Request #%d\", count)\n" +
						"}",
					Options:     []string{"fmt.Fprintf should use Fprintln instead", "count should be a local variable", "count++ is not goroutine-safe — needs sync/atomic or mutex", "w should be flushed after writing"},
					CorrectIdx:  2,
					Explanation: "HTTP handlers run concurrently. Incrementing a shared int without synchronization is a data race. AI often generates code that looks correct single-threaded but fails under concurrency.",
				},
				{ // Python
					Code: "count = 0\n" +
						"\n" +
						"async def handle_request(request):\n" +
						"    global count\n" +
						"    count += 1\n" +
						"    return Response(f\"Request #{count}\")",
					Options:     []string{"Should use nonlocal instead of global", "count should be a local variable", "count += 1 is not safe with concurrent async tasks — needs a lock", "Response should be JSONResponse"},
					CorrectIdx:  2,
					Explanation: "With async frameworks, multiple coroutines can interleave at await points. count += 1 (read-modify-write) is not atomic and can produce wrong values under concurrent requests.",
				},
				{ // JS
					Code: "let count = 0;\n" +
						"\n" +
						"app.get('/api', (req, res) => {\n" +
						"    count++;\n" +
						"    res.send(`Request #${count}`);\n" +
						"});",
					Options:     []string{"Should use const instead of let", "count should be scoped inside the handler", "If using worker threads/cluster, count++ is not safe across processes", "res.send should be res.json"},
					CorrectIdx:  2,
					Explanation: "In a clustered Node.js app (or worker threads), each process has its own count variable. The counter would be wrong. AI generates single-process code that breaks in production deployments.",
				},
			},
		},
		{
			Title: "Find the subtle logic error",
			Variants: [3]bugVariant{
				{ // Go
					Code: "func isValidAge(age int) bool {\n" +
						"    if age < 0 || age > 150 {\n" +
						"        return true\n" +
						"    }\n" +
						"    return false\n" +
						"}",
					Options:     []string{"150 should be 120 for realistic age limit", "The condition operators should be && not ||", "The return values are swapped — returns true for invalid ages", "age should be int64 not int"},
					CorrectIdx:  2,
					Explanation: "The function returns true when age IS out of range — the exact opposite of what 'isValid' implies. AI can produce logic that reads naturally but has inverted conditions.",
				},
				{ // Python
					Code: "def is_valid_age(age: int) -> bool:\n" +
						"    if age < 0 or age > 150:\n" +
						"        return True\n" +
						"    return False",
					Options:     []string{"150 should be 120 for realistic age limit", "Should use 'and' instead of 'or'", "The return values are swapped — returns True for invalid ages", "Should raise ValueError instead"},
					CorrectIdx:  2,
					Explanation: "The function returns True when age IS out of range — the exact opposite of what 'is_valid' implies. AI can produce logic that reads naturally but has inverted conditions.",
				},
				{ // JS
					Code: "function isValidAge(age) {\n" +
						"    if (age < 0 || age > 150) {\n" +
						"        return true;\n" +
						"    }\n" +
						"    return false;\n" +
						"}",
					Options:     []string{"150 should be 120 for realistic age limit", "Should use && instead of ||", "The return values are swapped — returns true for invalid ages", "age should be checked for NaN"},
					CorrectIdx:  2,
					Explanation: "The function returns true when age IS out of range — the exact opposite of what 'isValid' implies. AI can produce logic that reads naturally but has inverted conditions.",
				},
			},
		},
	}
}

func NewBugHunterModel(w, h int) Model {
	return &BugHunterModel{
		width:      w,
		height:     h,
		challenges: bugChallenges(),
		picking:    true,
	}
}

func (m *BugHunterModel) Init() tea.Cmd { return nil }

func (m *BugHunterModel) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// Language picker phase
		if m.picking {
			switch {
			case key.Matches(msg, key.NewBinding(key.WithKeys("up", "k"))):
				if m.langCursor > 0 {
					m.langCursor--
				}
			case key.Matches(msg, key.NewBinding(key.WithKeys("down", "j"))):
				if m.langCursor < 2 {
					m.langCursor++
				}
			case key.Matches(msg, key.NewBinding(key.WithKeys("1"))):
				m.lang = langGo
				m.picking = false
			case key.Matches(msg, key.NewBinding(key.WithKeys("2"))):
				m.lang = langPython
				m.picking = false
			case key.Matches(msg, key.NewBinding(key.WithKeys("3"))):
				m.lang = langJS
				m.picking = false
			case key.Matches(msg, key.NewBinding(key.WithKeys("enter", " "))):
				m.lang = bugLang(m.langCursor)
				m.picking = false
			}
			return m, nil
		}

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

		v := m.challenges[m.current].Variants[m.lang]
		switch {
		case key.Matches(msg, key.NewBinding(key.WithKeys("up", "k"))):
			if m.cursor > 0 {
				m.cursor--
			}
		case key.Matches(msg, key.NewBinding(key.WithKeys("down", "j"))):
			if m.cursor < len(v.Options)-1 {
				m.cursor++
			}
		case key.Matches(msg, key.NewBinding(key.WithKeys("1"))):
			m.cursor = 0
			m.submitBug()
		case key.Matches(msg, key.NewBinding(key.WithKeys("2"))):
			m.cursor = 1
			m.submitBug()
		case key.Matches(msg, key.NewBinding(key.WithKeys("3"))):
			m.cursor = 2
			m.submitBug()
		case key.Matches(msg, key.NewBinding(key.WithKeys("4"))):
			if len(v.Options) > 3 {
				m.cursor = 3
				m.submitBug()
			}
		case key.Matches(msg, key.NewBinding(key.WithKeys("enter", " "))):
			m.submitBug()
		case key.Matches(msg, key.NewBinding(key.WithKeys("r"))):
			m.reset()
		}
	}
	return m, nil
}

func (m *BugHunterModel) reset() {
	m.current = 0
	m.cursor = 0
	m.answered = false
	m.score = 0
	m.picking = true
	m.langCursor = int(m.lang)
}

func (m *BugHunterModel) submitBug() {
	v := m.challenges[m.current].Variants[m.lang]
	m.answered = true
	m.correct = m.cursor == v.CorrectIdx
	if m.correct {
		m.score++
	}
}

func (m *BugHunterModel) View() string {
	accent := lipgloss.NewStyle().Foreground(ui.ColorAccent).Bold(true)
	bad := lipgloss.NewStyle().Foreground(ui.ColorIncorrect).Bold(true)
	good := lipgloss.NewStyle().Foreground(ui.ColorCorrect).Bold(true)
	highlight := lipgloss.NewStyle().Foreground(ui.ColorHighlight).Bold(true)
	dim := lipgloss.NewStyle().Foreground(ui.ColorMuted)
	text := lipgloss.NewStyle().Foreground(lipgloss.Color("#d1d5db"))
	explain := lipgloss.NewStyle().Foreground(ui.ColorIntermediate)
	codeStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#a5f3fc"))

	var lines []string
	lines = append(lines, "")
	lines = append(lines, accent.Render("  Bug Hunter — Review AI-Generated Code"))
	lines = append(lines, dim.Render("  Find the bug in each code snippet"))
	lines = append(lines, "")

	// Language picker
	if m.picking {
		lines = append(lines, highlight.Render("  Choose your programming language:"))
		lines = append(lines, "")
		for i, name := range bugLangNames {
			prefix := "    "
			style := text
			if i == m.langCursor {
				prefix = "  ▸ "
				style = highlight
			}
			lines = append(lines, prefix+style.Render(fmt.Sprintf("%d) %s", i+1, name)))
		}
		lines = append(lines, "", dim.Render("  [↑/↓] Navigate  [1-3] Select  [Enter] Confirm"))
		return strings.Join(lines, "\n")
	}

	langLabel := dim.Render(fmt.Sprintf("  Language: %s", bugLangNames[m.lang]))
	lines = append(lines, langLabel)
	lines = append(lines, "")

	if m.current >= len(m.challenges) {
		lines = append(lines, good.Render(fmt.Sprintf("  Exercise Complete! Score: %d/%d", m.score, len(m.challenges))))
		lines = append(lines, "")
		if m.score == len(m.challenges) {
			lines = append(lines, good.Render("  Perfect! You have a sharp eye for AI-generated bugs."))
		} else if m.score >= 3 {
			lines = append(lines, text.Render("  Good work! Remember to check bounds, error handling, and concurrency."))
		} else {
			lines = append(lines, text.Render("  Keep practicing — review AI code like a junior developer's PR."))
		}
		lines = append(lines, "", dim.Render("  [r] Try again (change language)"))
		return strings.Join(lines, "\n")
	}

	ch := m.challenges[m.current]
	v := ch.Variants[m.lang]
	lines = append(lines, dim.Render(fmt.Sprintf("  Challenge %d of %d", m.current+1, len(m.challenges))))
	lines = append(lines, highlight.Render("  "+ch.Title))
	lines = append(lines, "")

	// Code display
	codeLines := strings.Split(v.Code, "\n")
	for _, cl := range codeLines {
		lines = append(lines, codeStyle.Render("  "+cl))
	}
	lines = append(lines, "")
	lines = append(lines, highlight.Render("  What's the bug?"))
	lines = append(lines, "")

	for i, opt := range v.Options {
		prefix := fmt.Sprintf("  %d) ", i+1)
		style := text

		if m.answered {
			if i == v.CorrectIdx {
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
			lines = append(lines, good.Render("  ✓ Correct!"))
		} else {
			lines = append(lines, bad.Render("  ✗ Not quite"))
		}
		lines = append(lines, explain.Render("  "+v.Explanation))
		lines = append(lines, "", highlight.Render("  Press Enter to continue"))
	}

	lines = append(lines, "", dim.Render("  [↑/↓] Navigate  [1-4] Select  [Enter] Confirm  [r] Restart"))

	return strings.Join(lines, "\n")
}
