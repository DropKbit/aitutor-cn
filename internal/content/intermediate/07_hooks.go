package intermediate

import (
	"github.com/DropKbit/aitutor-cn/internal/lesson"
	"github.com/DropKbit/aitutor-cn/internal/viz"
	"github.com/DropKbit/aitutor-cn/pkg/types"
)

func init() {
	lesson.Register(types.LessonDef{
		ID:         7,
		Title:      "钩子（Hooks）",
		Tier:       types.Intermediate,
		Summary:    "理解 AI 助手动作生命周期中的 Hook 机制",
		SourceFile: "internal/content/intermediate/07_hooks.go",
		VizBuilder: func(w, h int) interface{} { return viz.NewLifecycleModel(w, h) },
		Theory: []types.TheoryBlock{
			{Kind: types.Heading, Content: "Hook"},
			{Kind: types.Paragraph, Content: "Hook 是用户定义的 shell 命令，会在 AI 助手的特定事件触发时自动执行。它可以帮助你自定义行为、执行策略、并把 AI 工具接入你现有的工作流。"},
			{Kind: types.Heading, Content: "Hook 类型"},
			{Kind: types.Bullet, Content: "PreToolUse：工具执行前运行（可以阻止工具继续执行）\nPostToolUse：工具执行后运行\nNotification：当 AI 想通知你时运行\nSessionStart：新会话开始时运行\nPromptSubmit：当你发送消息时运行"},
			{Kind: types.Heading, Content: "Hook 如何工作"},
			{Kind: types.Code, Content: "  User sends message\n       │\n       ▼\n  ┌────────────────────┐\n  │ PromptSubmit hook  │\n  └────────────────────┘\n       │\n       ▼\n  AI decides to use tool\n       │\n       ▼\n  ┌────────────────────┐\n  │ PreToolUse hook    │  ← can BLOCK the tool\n  └────────────────────┘\n       │\n       ▼\n  Tool executes\n       │\n       ▼\n  ┌────────────────────┐\n  │ PostToolUse hook   │\n  └────────────────────┘"},
			{Kind: types.Heading, Content: "示例：编辑后自动格式化"},
			{Kind: types.Code, Content: "  // 示例：Claude Code hooks（.claude/hooks.json）\n  // 其他工具也有类似生命周期事件，只是语法不同\n  {\n    \"hooks\": {\n      \"PostToolUse\": [{\n        \"matcher\": \"Edit\",\n        \"command\": \"./scripts/format-on-save.sh\"\n      }]\n    }\n  }"},
			{Kind: types.Callout, Content: "Hook 很适合执行团队规范，例如每次编辑后自动运行 linter，或阻止对受保护文件的写入。"},
		},
		Questions: []types.QuizQuestion{
			{
				Kind:        types.MultipleChoice,
				Prompt:      "哪种 Hook 可以阻止工具继续执行？",
				Choices:     []string{"PostToolUse", "Notification", "PreToolUse", "SessionStart"},
				CorrectIdx:  2,
				Explanation: "PreToolUse 会在工具执行前触发，因此可以直接阻止工具运行。",
			},
			{
				Kind:        types.MultipleChoice,
				Prompt:      "PostToolUse Hook 的一个实际用途是什么？",
				Choices:     []string{"阻止危险命令", "编辑后自动格式化文件", "启动新会话", "发送通知"},
				CorrectIdx:  1,
				Explanation: "PostToolUse 在工具执行后触发，非常适合在 Edit 之后自动格式化代码。",
			},
		},
	})
}
