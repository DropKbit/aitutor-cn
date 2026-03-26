package beginner

import (
	"github.com/DropKbit/aitutor-cn/internal/lesson"
	"github.com/DropKbit/aitutor-cn/internal/viz"
	"github.com/DropKbit/aitutor-cn/pkg/types"
)

func init() {
	lesson.Register(types.LessonDef{
		ID:         2,
		Title:      "上下文窗口",
		Tier:       types.Beginner,
		Summary:    "理解 AI 模型如何看到并处理信息",
		SourceFile: "internal/content/beginner/02_context_window.go",
		VizBuilder: func(w, h int) interface{} { return viz.NewBucketModel(w, h) },
		Theory: []types.TheoryBlock{
			{Kind: types.Heading, Content: "什么是上下文窗口"},
			{Kind: types.Paragraph, Content: "每个 AI 模型都有一个“上下文窗口”，也就是它一次能够处理的固定文本量，通常以 token 计量。你可以把它理解成模型的工作记忆。模型理解任务所需的一切内容，都必须放进这个窗口中。"},
			{Kind: types.Heading, Content: "上下文窗口里都装了什么？"},
			{Kind: types.Bullet, Content: "系统提示词：定义模型行为的基础指令\n项目配置文件：会自动加载的工具级说明，如 AGENTS.md、CLAUDE.md、.cursorrules、copilot-instructions.md\n会话历史：当前会话中的所有先前消息\n工具结果：读文件、搜索、执行命令后的输出\n文件内容：你正在处理的代码文件"},
			{Kind: types.Callout, Content: "延伸阅读：什么是 token？可以试试交互式 tokenizer — https://platform.openai.com/tokenizer"},
			{Kind: types.Heading, Content: "Token 上限"},
			{Kind: types.Paragraph, Content: "不同模型的上下文窗口大小不同。例如 Claude 提供 200K tokens，GPT-4o 为 128K，Gemini 最高可达 1M。虽然已经很大，但并不是无限的。大型代码库、长会话和冗长工具输出都可能把它填满。"},
			{Kind: types.Code, Content: "  示例：200,000 tokens 大约相当于\n  • 150,000 个英文单词\n  • 30,000 行代码\n  • 约 100 个平均大小的源码文件"},
			{Kind: types.Heading, Content: "上下文窗口管理"},
			{Kind: types.Paragraph, Content: "优秀的 AI 工具会自动替你管理上下文窗口。它们会压缩旧消息、总结工具结果，并优先保留最相关的信息。但理解这项限制，仍然能帮助你更高效地使用它们。"},
			{Kind: types.Callout, Content: "如果你发现在长对话里 AI 助手“忘记了”之前的内容，通常是因为那些信息被压缩了，或者已经从上下文窗口中移除。"},
			{Kind: types.Heading, Content: "隐藏成本：MCP 工具定义"},
			{Kind: types.Paragraph, Content: "每启用一个 MCP server，它的工具定义就会进入上下文窗口，哪怕你还没调用任何工具。每个工具定义大约消耗 200 tokens。有些 server 非常臃肿，仅 GitHub 一个就会暴露 34 个工具。"},
			{Kind: types.Code, Content: "  已启用的 MCP 服务：        上下文成本：\n  ─────────────────────       ──────────────\n  GitHub（34 个工具）         6,800 tokens\n  Slack（18 个工具）          3,600 tokens\n  Jira（22 个工具）           4,400 tokens\n  Linear（16 个工具）         3,200 tokens\n  ─────────────────────       ──────────────\n  总计                        18,000 tokens\n                              （不用也会白白占用！）"},
			{Kind: types.Callout, Content: "只启用你真正需要的 MCP server，并在启用的 server 中关闭不需要的单个工具。GitHub 虽有 34 个工具，但你可能只会用到其中 5 个。每个不用的工具定义都会悄悄浪费约 200 tokens。可以在可视化中亲自试试看。"},
			{Kind: types.Heading, Content: "高效使用上下文的策略"},
			{Kind: types.Bullet, Content: "尽量具体：明确请求比模糊请求更省上下文\n使用项目配置文件：持久指令无需反复重复\n拆分大任务：小而聚焦的会话通常更有效\n按路径引用文件：让 AI 只读取真正需要的内容\n关闭不用的 MCP server：工具定义会悄悄吃掉上下文\n关闭 server 中不用的单个工具：只保留真正会用到的"},
		},
		Questions: []types.QuizQuestion{
			{
				Kind:        types.MultipleChoice,
				Prompt:      "什么是上下文窗口？",
				Choices:     []string{"显示右键菜单的图形界面窗口", "AI 一次能处理的固定文本量", "查看变量的调试工具", "你输入命令的终端窗口"},
				CorrectIdx:  1,
				Explanation: "上下文窗口就是模型的工作记忆，所有输入内容都必须放在它的 token 限额之内。",
			},
			{
				Kind:        types.MultipleChoice,
				Prompt:      "当会话内容超过上下文窗口时，会发生什么？",
				Choices:     []string{"AI 会崩溃", "旧消息会被压缩或移除", "上下文窗口会自动扩容", "AI 会要求你重新开始"},
				CorrectIdx:  1,
				Explanation: "当上下文窗口被填满时，AI 工具会压缩或移除旧消息，为新内容腾出空间。",
			},
		},
	})
}
