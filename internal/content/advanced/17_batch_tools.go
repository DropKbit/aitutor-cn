package advanced

import (
	"github.com/DropKbit/aitutor-cn/internal/lesson"
	"github.com/DropKbit/aitutor-cn/internal/viz"
	"github.com/DropKbit/aitutor-cn/pkg/types"
)

func init() {
	lesson.Register(types.LessonDef{
		ID:         17,
		Title:      "批量工具调用",
		Tier:       types.Advanced,
		Summary:    "理解按工具定义执行策略与并行批处理",
		SourceFile: "internal/content/advanced/17_batch_tools.go",
		VizBuilder: func(w, h int) interface{} { return viz.NewBatchToolModel(w, h) },
		Theory: []types.TheoryBlock{
			{Kind: types.Heading, Content: "批量工具调用"},
			{Kind: types.Paragraph, Content: "AI 助手可以在一次响应中同时调用多个工具。但并不是所有工具都适合并行执行，有些工具带有副作用，必须独占运行。现代 AI 系统通常允许每个工具声明自己的批处理策略。"},
			{Kind: types.Heading, Content: "按工具定义批处理策略"},
			{Kind: types.Paragraph, Content: "每个工具都可以声明自己是否允许被批量执行，也就是与其他工具一起运行；或者是否必须单独顺序执行。像文件读取器、搜索器这类只读工具通常可批量执行；而文件编辑器、shell 命令这类带副作用的工具，则往往必须串行。"},
			{Kind: types.Code, Content: "  Example tool set:\n  Tool         Policy        Why\n  ────         ──────        ───\n  file read    ⚡ batchable   Read-only, no side effects\n  search       ⚡ batchable   Read-only search\n  grep         ⚡ batchable   Read-only search\n  file edit    🔒 sequential  Modifies files\n  file write   🔒 sequential  Creates/overwrites files\n  shell        🔒 sequential  Arbitrary side effects"},
			{Kind: types.Heading, Content: "批处理如何工作"},
			{Kind: types.Paragraph, Content: "AI 会把连续出现的、允许批量执行的工具调用打包成一个 round trip。当遇到只能顺序执行的工具时，它就先把当前批次发出去，然后单独运行该工具。这样既能最大化并行度，也能遵守安全约束。"},
			{Kind: types.Code, Content: "  Tool calls in order:         Execution plan:\n  ───────────────────         ───────────────\n  Read(go.mod)      ─┐\n  Read(main.go)     ─┤ batch   → Round 1 (3 parallel)\n  Grep(\"TODO\")      ─┘\n  Edit(main.go)     ─── alone  → Round 2 (1 alone)\n  Bash(go build)    ─── alone  → Round 3 (1 alone)\n  Bash(go test)     ─── alone  → Round 4 (1 alone)\n\n  Result: 4 round trips instead of 6"},
			{Kind: types.Heading, Content: "为什么不能把一切都批量执行？"},
			{Kind: types.Bullet, Content: "同一文件上的 Edit + Edit 可能发生冲突\n某些 Bash 命令依赖之前 Edit 的结果\nWrite 可能会创建一个随后被其他工具读取的新文件\n顺序工具需要依赖前一步的输出结果"},
			{Kind: types.Heading, Content: "如何最大化并行度"},
			{Kind: types.Paragraph, Content: "理解批处理策略之后，你就能更有意识地组织请求，以获得更快速度。一个常见技巧是把读取操作前置：先让 AI 把所有信息收集齐，再开始修改。这样天然就能把可批量执行的读取操作归到一起。"},
			{Kind: types.Code, Content: "  Slow (interleaved):           Fast (reads first):\n  ──────────────────           ──────────────────\n  Read A → Edit A               Read A ─┐\n  Read B → Edit B               Read B ─┤ 1 round trip\n  Read C → Edit C               Read C ─┘\n  = 6 round trips               Edit A → Edit B → Edit C\n                                = 4 round trips"},
			{Kind: types.Callout, Content: "可以在可视化里亲自试一试：切换各个工具的批处理策略，观察执行计划如何变化。你会看到，把读取操作聚在一起能明显减少 round trip。"},
		},
		Questions: []types.QuizQuestion{
			{
				Kind:        types.MultipleChoice,
				Prompt:      "为什么 Read 和 Grep 这类工具通常会被标记为可批量执行？",
				Choices:     []string{"因为它们比其他工具更快", "因为它们是只读操作，没有副作用", "因为它们使用更少 token", "因为它们返回结果总是更小"},
				CorrectIdx:  1,
				Explanation: "只读工具没有副作用，因此并行执行通常是安全的，它们之间不会相互干扰。",
			},
			{
				Kind:        types.MultipleChoice,
				Prompt:      "当 AI 在一个批次中遇到只能顺序执行的工具时，会发生什么？",
				Choices:     []string{"它会跳过这个工具", "它会把这个工具改成可批量执行", "它会先刷出当前批次，然后单独运行该工具", "它会等待用户确认"},
				CorrectIdx:  2,
				Explanation: "只能顺序执行的工具会迫使 AI 先执行掉当前待发批次，然后把该工具单独运行，再继续后续步骤。",
			},
		},
	})
}
