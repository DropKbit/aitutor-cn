package intermediate

import (
	"github.com/DropKbit/aitutor-cn/internal/lesson"
	"github.com/DropKbit/aitutor-cn/internal/viz"
	"github.com/DropKbit/aitutor-cn/pkg/types"
)

func init() {
	lesson.Register(types.LessonDef{
		ID:         9,
		Title:      "代理式循环",
		Tier:       types.Intermediate,
		Summary:    "理解 AI 代理如何通过循环迭代解决问题",
		SourceFile: "internal/content/intermediate/09_agentic_loop.go",
		VizBuilder: func(w, h int) interface{} { return viz.NewAgenticLoopModel(w, h) },
		Theory: []types.TheoryBlock{
			{Kind: types.Heading, Content: "代理式循环"},
			{Kind: types.Paragraph, Content: "每一个 AI 编程代理的核心都是一个循环。它不是一次性把代码生成完就结束，而是不断迭代：读取上下文、进行推理、执行动作、观察结果，然后再回到开头。这正是代理与普通代码生成器之间的根本区别。"},
			{Kind: types.Heading, Content: "核心循环：Read → Think → Act → Observe"},
			{Kind: types.Paragraph, Content: "几乎所有代理系统都遵循这一模式的某种变体："},
			{Kind: types.Code, Content: "  Read ──> Think ──> Act ──> Observe\n    ^                           │\n    └───────── loops back ──────┘\n\n  Read:     Gather context (files, errors, docs)\n  Think:    Reason about what to do next\n  Act:      Execute a tool (edit, run, search)\n  Observe:  Check the result, feed it back in"},
			{Kind: types.Paragraph, Content: "这个循环会一直持续，直到任务完成，或代理判断自己无法继续推进。关键在于：每一轮迭代的输出，都会成为下一轮的输入，因此整个过程具备自我纠错能力。"},
			{Kind: types.Heading, Content: "为什么循环如此重要"},
			{Kind: types.Paragraph, Content: "如果没有循环，AI 就只能做一次性响应：生成一段代码，然后“希望它能运行”。有了循环，代理就可以："},
			{Kind: types.Bullet, Content: "通过运行测试验证自己的工作\n在失败后进行自我修正\n在第一次尝试信息不足时继续补充上下文\n把复杂任务拆成更小步骤逐步完成\n在代码库与预期不一致时动态适应"},
			{Kind: types.Heading, Content: "真实示例：修复一个 bug"},
			{Kind: types.Code, Content: "  Iteration 1: Search for the bug\n    Read:    Grep for the error message\n    Think:   Found the handler, need to read it\n    Act:     Read the file\n    Observe: See missing error check → need to fix\n\n  Iteration 2: Apply the fix\n    Read:    Understand the code around the bug\n    Think:   Add error handling after the query\n    Act:     Edit the file\n    Observe: Fix applied → need to test\n\n  Iteration 3: Verify\n    Read:    Run tests\n    Think:   One test failed! Different code path.\n    Act:     Edit to handle that case too\n    Observe: Tests pass → done!"},
			{Kind: types.Heading, Content: "循环的不同变体"},
			{Kind: types.Paragraph, Content: "不同代理框架会用不同名字描述这些步骤，但底层模式是相同的："},
			{Kind: types.Code, Content: "  Pattern          Steps\n  ──────           ─────\n  OODA             Observe → Orient → Decide → Act\n  ReAct            Reason → Act → Observe\n  Plan-and-Execute Plan → Execute → Observe → Replan\n  General agent    Perceive → Decide → Execute → Evaluate"},
			{Kind: types.Callout, Content: "名字可以不同，但所有代理系统都遵循同一原则：收集信息、进行推理、执行动作、检查结果、然后重复。正是这个循环把语言模型变成了代理。"},
			{Kind: types.Callout, Content: "延伸阅读：ReAct pattern — https://arxiv.org/abs/2210.03629 | OODA loop — https://en.wikipedia.org/wiki/OODA_loop"},
			{Kind: types.Heading, Content: "谁来控制这个循环？"},
			{Kind: types.Bullet, Content: "停止条件：代理判断任务已完成，例如测试通过或用户确认\n最大迭代次数：安全阈值，用来防止无限循环\n错误处理：代理卡住时可以提前退出\n用户干预：人类可以随时重定向或终止代理\nToken 预算：上下文窗口限制了能容纳多少轮迭代"},
			{Kind: types.Heading, Content: "一次性响应 vs 代理式循环"},
			{Kind: types.Code, Content: "  Single-shot:              Agentic:\n  ──────────               ────────\n  Prompt → Response         Prompt → Loop ──┐\n  (hope it's right)               ↑         │\n                                  └─────────┘\n                            (verify and correct)\n\n  Speed: Fast               Speed: Slower per task\n  Accuracy: Variable        Accuracy: High (self-correcting)\n  Complexity: Simple only   Complexity: Handles hard tasks"},
			{Kind: types.Callout, Content: "试试可视化示例，按步骤走一遍真实调试场景。你会看到每一轮迭代如何建立在上一轮之上：失败不是终点，而是驱动下一步行动的信息。"},
		},
		Questions: []types.QuizQuestion{
			{
				Kind:        types.MultipleChoice,
				Prompt:      "代理式循环与一次性响应相比，最大的区别是什么？",
				Choices:     []string{"它会用掉更多 token", "它会迭代：观察结果并自我修正", "它写出来的代码总是更好", "它运行得更快"},
				CorrectIdx:  1,
				Explanation: "代理式循环会不断迭代：先行动，再观察结果，并把观察结果作为下一步输入。这种自我修正能力，正是代理与一次性生成的区别所在。",
			},
			{
				Kind:        types.MultipleChoice,
				Prompt:      "在典型的代理循环中，测试失败后会发生什么？",
				Choices:     []string{"代理直接放弃并报告失败", "失败结果会变成下一轮迭代的输入", "代理从头重新开始", "代理立即询问用户下一步怎么做"},
				CorrectIdx:  1,
				Explanation: "失败本身就是信息。测试输出会成为下一轮“Read”的输入，代理据此分析问题并尝试新的办法。",
			},
		},
	})
}
