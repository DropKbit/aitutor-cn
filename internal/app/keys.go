package app

import "github.com/charmbracelet/bubbles/key"

// KeyMap defines all keybindings for the application.
type KeyMap struct {
	Quit         key.Binding
	Tab          key.Binding
	Next         key.Binding
	Prev         key.Binding
	Advance      key.Binding
	AdvancePhase key.Binding
	Back         key.Binding
	Help         key.Binding
	Up           key.Binding
	Down         key.Binding
	Select       key.Binding
	Space        key.Binding
	Reset        key.Binding
}

var Keys = KeyMap{
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "退出"),
	),
	Tab: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("Tab", "切换侧边栏"),
	),
	Next: key.NewBinding(
		key.WithKeys("n"),
		key.WithHelp("n", "下一课"),
	),
	Prev: key.NewBinding(
		key.WithKeys("p"),
		key.WithHelp("p", "上一课"),
	),
	Advance: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("Enter", "进入下一阶段"),
	),
	AdvancePhase: key.NewBinding(
		key.WithKeys("right"),
		key.WithHelp("→", "下一阶段"),
	),
	Back: key.NewBinding(
		key.WithKeys("backspace", "left"),
		key.WithHelp("←/Bksp", "上一阶段"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "帮助"),
	),
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "向上滚动"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "向下滚动"),
	),
	Select: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("Enter", "选择"),
	),
	Space: key.NewBinding(
		key.WithKeys(" "),
		key.WithHelp("Space", "交互"),
	),
	Reset: key.NewBinding(
		key.WithKeys("r"),
		key.WithHelp("r", "重置"),
	),
}
