package advanced

import (
	"github.com/DropKbit/aitutor-cn/internal/lesson"
	"github.com/DropKbit/aitutor-cn/internal/viz"
	"github.com/DropKbit/aitutor-cn/pkg/types"
)

func init() {
	lesson.Register(types.LessonDef{
		ID:         12,
		Title:      "MCP（模型上下文协议）",
		Tier:       types.Advanced,
		Summary:    "理解如何通过外部工具服务器扩展 AI 能力",
		SourceFile: "internal/content/advanced/12_mcp.go",
		VizBuilder: func(w, h int) interface{} { return viz.NewMCPCallerModel(w, h) },
		Theory: []types.TheoryBlock{
			{Kind: types.Heading, Content: "MCP（Model Context Protocol）"},
			{Kind: types.Paragraph, Content: "MCP 是一个开放协议，用来让 AI 助手连接外部工具服务器。它不要求把所有能力都硬编码进模型本身，而是允许通过插件式架构，把工具、资源和提示模板交给专门的 server 提供。"},
			{Kind: types.Heading, Content: "架构"},
			{Kind: types.Code, Content: "  ┌──────────┐     stdio/HTTP    ┌──────────────┐\n  │ AI 客户端 │◄════════════════►│  MCP Server  │\n  │          │                  │              │\n  └──────────┘                  │  ┌────────┐  │\n                                │  │ 工具 1 │  │\n                                │  ├────────┤  │\n                                │  │ 工具 2 │  │\n                                │  ├────────┤  │\n                                │  │ 资源   │  │\n                                │  └────────┘  │\n                                └──────────────┘"},
			{Kind: types.Heading, Content: "关键概念"},
			{Kind: types.Bullet, Content: "Tools：AI 可调用的函数，例如查数据库、发 Slack 消息\nResources：AI 可读取的数据，例如文档和 API schema\nPrompts：用于常见任务的可复用提示模板\nSampling：允许 server 请求客户端执行 LLM completion\nTransports：通信通道，本地常用 stdio，远程常用 Streamable HTTP"},
			{Kind: types.Heading, Content: "配置方式"},
			{Kind: types.Code, Content: "  // MCP 配置（路径因工具而异：.claude/mcp.json、.cursor/mcp.json 等）\n  {\n    \"mcpServers\": {\n      \"github\": {\n        \"command\": \"gh-mcp-server\",\n        \"args\": [\"--repo\", \"owner/repo\"]\n      },\n      \"database\": {\n        \"command\": \"db-mcp-server\",\n        \"args\": [\"--connection\", \"postgres://...\"]\n      }\n    }\n  }"},
			{Kind: types.Callout, Content: "MCP 把 AI 助手从封闭系统变成可扩展平台。任何开发者都能通过编写 MCP server，为 AI 增加新能力。"},
			{Kind: types.Callout, Content: "延伸阅读：MCP Specification — https://spec.modelcontextprotocol.io/ | MCP Introduction — https://modelcontextprotocol.io/introduction"},
		},
		Questions: []types.QuizQuestion{
			{
				Kind:        types.MultipleChoice,
				Prompt:      "MCP 是什么的缩写？",
				Choices:     []string{"Model Control Protocol", "Model Context Protocol", "Machine Code Pipeline", "Multi-Channel Processor"},
				CorrectIdx:  1,
				Explanation: "MCP 是 Model Context Protocol 的缩写，它是一个把 AI 连接到外部工具服务器的开放协议。",
			},
			{
				Kind:        types.MultipleChoice,
				Prompt:      "MCP 在本地 server 上通常使用哪种 transport？",
				Choices:     []string{"HTTP", "WebSocket", "stdio", "gRPC"},
				CorrectIdx:  2,
				Explanation: "MCP 在本地 server 上通常使用 stdio（进程间通信），远程 server 则常用 Streamable HTTP。",
			},
		},
	})
}
