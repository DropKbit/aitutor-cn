package beginner

import (
	"github.com/DropKbit/aitutor-cn/internal/lesson"
	"github.com/DropKbit/aitutor-cn/internal/viz"
	"github.com/DropKbit/aitutor-cn/pkg/types"
)

func init() {
	lesson.Register(types.LessonDef{
		ID:         3,
		Title:      "工具",
		Tier:       types.Beginner,
		Summary:    "理解 AI 助手如何与代码库交互",
		SourceFile: "internal/content/beginner/03_tools.go",
		VizBuilder: func(w, h int) interface{} { return viz.NewToolFlowModel(w, h) },
		Theory: []types.TheoryBlock{
			{Kind: types.Heading, Content: "工具：AI 的手和眼"},
			{Kind: types.Paragraph, Content: "AI 模型本身只能思考和生成文本。想真正做事，比如读文件、搜代码、运行命令，就需要借助工具。工具本质上是 AI 可以调用的函数，用来与真实世界交互。"},
			{Kind: types.Heading, Content: "常见工具类别"},
			{Kind: types.Bullet, Content: "文件操作：Read、Write、Edit 文件\n搜索：Glob（按模式找文件）、Grep（搜索内容）\n执行：Bash（运行 shell 命令）\n导航：LSP（跳转定义、查找引用）"},
			{Kind: types.Heading, Content: "工具调用流程"},
			{Kind: types.Code, Content: "  1. AI 判断自己需要更多信息\n  2. AI 调用工具（例如 Read file.go）\n  3. 工具执行并返回结果\n  4. AI 处理结果\n  5. AI 继续推理或调用下一个工具"},
			{Kind: types.Heading, Content: "专用工具 vs 通用工具"},
			{Kind: types.Paragraph, Content: "像 Read、Edit、Grep 这样的专用工具，通常优先于 Bash 这类通用工具。因为它们更安全、意图更清晰，也更便于审查 AI 到底做了什么。"},
			{Kind: types.Code, Content: "  ✓ Read(\"src/main.go\")        — 意图清晰、安全、易审查\n  ✗ Bash(\"cat src/main.go\")    — 语义不透明，更难审查"},
			{Kind: types.Heading, Content: "权限模型"},
			{Kind: types.Paragraph, Content: "工具运行在权限系统之下。有些工具可以自动执行，例如读文件；另一些则需要显式批准，例如执行任意 shell 命令或写文件。这样能确保你始终掌握控制权。"},
			{Kind: types.Callout, Content: "可以把工具理解成 AI 的手和眼。没有工具，它只能“想”；有了工具，它才能探索、修改和构建。"},
			{Kind: types.Callout, Content: "延伸阅读：Language Server Protocol — https://microsoft.github.io/language-server-protocol/ | Glob patterns — https://en.wikipedia.org/wiki/Glob_(programming)"},
		},
		Questions: []types.QuizQuestion{
			{
				Kind:        types.MultipleChoice,
				Prompt:      "为什么专用工具（Read、Edit、Grep）通常比 Bash 更优先？",
				Choices:     []string{"因为它们更快", "因为更安全、意图更清晰、也更容易审查", "因为它们占用更少内存", "因为它们更新"},
				CorrectIdx:  1,
				Explanation: "专用工具能提供更好的安全性、更清晰的意图表达，也更方便审查 AI 的操作。",
			},
			{
				Kind:        types.Ordering,
				Prompt:      "请将工具调用流程按正确顺序排列：",
				Choices:     []string{"AI 判断自己需要信息", "AI 调用工具", "工具执行并返回结果", "AI 处理结果"},
				Explanation: "正确顺序是：判断 → 调用 → 执行 → 处理结果。",
			},
		},
	})
}
