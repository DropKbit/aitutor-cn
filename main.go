package main

import (
	"fmt"
	"os"

	"github.com/DropKbit/aitutor-cn/internal/app"
	tea "github.com/charmbracelet/bubbletea"

	// Register all lessons via init()
	_ "github.com/DropKbit/aitutor-cn/internal/content/advanced"
	_ "github.com/DropKbit/aitutor-cn/internal/content/beginner"
	_ "github.com/DropKbit/aitutor-cn/internal/content/intermediate"
)

var version = "dev"

func main() {
	p := tea.NewProgram(
		app.NewAppModel(version),
		tea.WithAltScreen(),
	)
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "错误：%v\n", err)
		os.Exit(1)
	}
}
