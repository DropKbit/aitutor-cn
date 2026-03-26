package advanced

import (
	"github.com/DropKbit/aitutor-cn/internal/lesson"
	"github.com/DropKbit/aitutor-cn/internal/viz"
	"github.com/DropKbit/aitutor-cn/pkg/types"
)

func init() {
	lesson.Register(types.LessonDef{
		ID:         16,
		Title:      "工具搜索与延迟加载工具",
		Tier:       types.Advanced,
		Summary:    "理解如何通过延迟加载工具优化上下文使用",
		SourceFile: "internal/content/advanced/16_tool_search.go",
		VizBuilder: func(w, h int) interface{} { return viz.NewToolSearchModel(w, h) },
		Theory: []types.TheoryBlock{
			{Kind: types.Heading, Content: "工具搜索问题"},
			{Kind: types.Paragraph, Content: "AI 助手可能拥有上百个可用工具，包括内置工具、MCP server 工具和插件工具。如果把所有工具定义都直接塞进上下文窗口，就会浪费成千上万 token。解决办法就是：延迟加载工具，只在需要时再加载。"},
			{Kind: types.Heading, Content: "延迟加载工具如何工作"},
			{Kind: types.Paragraph, Content: "延迟工具会先被注册，但不会立即进入上下文。它们会出现在一份可用工具名摘要列表中。等 AI 真正需要某个工具时，再通过搜索机制发现并按需加载完整定义。"},
			{Kind: types.Code, Content: "  Example: available deferred tools\n  ─────────────────────────────────\n  mcp__slack__send_message\n  mcp__github__create_pr\n  mcp__database__query\n  NotebookEdit\n  WebSearch"},
			{Kind: types.Heading, Content: "Tool Search 查询方式"},
			{Kind: types.Bullet, Content: "关键字搜索：例如 \"slack message\" 可以找到 Slack 发消息相关工具\n直接选择：按确切工具名直接加载\n过滤搜索：先限定某个类别，再按相关性排序\n多选加载：一次列出多个工具名，批量载入"},
			{Kind: types.Heading, Content: "为什么这很重要"},
			{Kind: types.Paragraph, Content: "如果没有延迟加载，50 个 MCP 工具、每个约 200 tokens，就意味着有 10,000 tokens 被浪费在你可能永远不会使用的工具定义上。有了延迟加载，只有真正需要的工具才会进入上下文。"},
			{Kind: types.Code, Content: "  Without deferred tools:\n  ┌──────────────────────────────┐\n  │ System Prompt (8k)           │\n  │ ALL tool definitions (10k)   │ ← wasted\n  │ Conversation (5k)            │\n  │ ...remaining: 177k           │\n  └──────────────────────────────┘\n\n  With deferred tools:\n  ┌──────────────────────────────┐\n  │ System Prompt (8k)           │\n  │ Core tools only (2k)         │ ← efficient\n  │ Conversation (5k)            │\n  │ ...remaining: 185k           │ ← 8k more!\n  └──────────────────────────────┘"},
			{Kind: types.Callout, Content: "调用延迟工具之前，一定要先使用 ToolSearch 把它加载进来。直接调用尚未加载的延迟工具会失败。"},
		},
		Questions: []types.QuizQuestion{
			{
				Kind:        types.MultipleChoice,
				Prompt:      "为什么有些工具会被设计为延迟加载，而不是始终放在上下文里？",
				Choices:     []string{"因为它们不重要", "为了节省上下文窗口中的 token", "因为它们是实验功能", "为了提升安全性"},
				CorrectIdx:  1,
				Explanation: "延迟加载通过“只在需要时加载定义”，节省了上下文窗口空间。",
			},
			{
				Kind:        types.MultipleChoice,
				Prompt:      "如果你已经知道某个工具的确切名字，应该使用哪种 ToolSearch 查询方式？",
				Choices:     []string{"\"slack tools\"", "\"select:mcp__slack__send_message\"", "\"+find slack\"", "\"load slack\""},
				CorrectIdx:  1,
				Explanation: "知道确切工具名时，应使用 \"select:tool_name\" 进行直接选择。",
			},
		},
	})
}
