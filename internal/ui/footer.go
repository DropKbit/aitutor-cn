package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// FooterModel renders the bottom bar with key hints.
type FooterModel struct {
	Width    int
	Bindings []KeyHint
}

// KeyHint is a key-description pair for the footer.
type KeyHint struct {
	Key  string
	Desc string
}

var DefaultBindings = []KeyHint{
	{Key: "q", Desc: "退出"},
	{Key: "Tab", Desc: "侧边栏"},
	{Key: "n/p", Desc: "上一课/下一课"},
	{Key: "Enter", Desc: "前进"},
	{Key: "?", Desc: "帮助"},
}

func NewFooterModel() FooterModel {
	return FooterModel{Bindings: DefaultBindings}
}

func (f FooterModel) View() string {
	var parts []string
	for _, b := range f.Bindings {
		key := FooterKeyStyle.Render(b.Key)
		desc := FooterDescStyle.Render(b.Desc)
		parts = append(parts, fmt.Sprintf("%s %s", key, desc))
	}
	content := strings.Join(parts, lipgloss.NewStyle().Foreground(ColorDim).Render("  │  "))
	return FooterStyle.Width(f.Width).Render(content)
}
