package intermediate

import (
	"github.com/DropKbit/aitutor-cn/internal/lesson"
	"github.com/DropKbit/aitutor-cn/internal/viz"
	"github.com/DropKbit/aitutor-cn/pkg/types"
)

func init() {
	lesson.Register(types.LessonDef{
		ID:         10,
		Title:      "高级提示技巧",
		Tier:       types.Intermediate,
		Summary:    "七种能显著提升 AI 代码生成质量的技巧",
		SourceFile: "internal/content/intermediate/10_codegen_prompts.go",
		VizBuilder: func(w, h int) interface{} { return viz.NewPromptBuilderModel(w, h) },
		Theory: []types.TheoryBlock{
			{Kind: types.Heading, Content: "七个真正有效的技巧"},
			{Kind: types.Paragraph, Content: "第 4 课讲了基础原则：要具体、要给上下文。这一课会更进一步，介绍七种经过验证、能够显著提升代码生成效果的技巧。它们都来自 AI 实验室和开发者社区的真实最佳实践。"},
			{Kind: types.Heading, Content: "1. Persona / 角色设定"},
			{Kind: types.Paragraph, Content: "在提示词里指定专家角色，能够聚焦 AI 的推理方式，并调动特定领域知识。一个“资深系统工程师”在思考缓存问题时，和一个普通助手的推理方式并不相同。哪怕只加一句话，也可能带来明显差异。"},
			{Kind: types.Code, Content: "  \"You are a senior Go developer specializing in\n   concurrent systems. Implement a worker pool that\n   processes jobs from a channel. It must be\n   thread-safe, handle panics in workers, and support\n   configurable pool size.\""},
			{Kind: types.Heading, Content: "2. Few-Shot 示例"},
			{Kind: types.Paragraph, Content: "给 AI 看 2 到 3 个你想要的示例，它就更容易照着模式来。这是让生成代码符合项目约定的最可靠方法之一，无论是命名、错误处理还是结构风格。"},
			{Kind: types.Code, Content: "  \"Here are two existing handlers in our codebase:\n   [paste handler A]\n   [paste handler B]\n   Following the exact same patterns (error handling,\n   response format, naming), create a handler for\n   DELETE /users/:id.\""},
			{Kind: types.Heading, Content: "3. 负例 / 对比示例"},
			{Kind: types.Paragraph, Content: "不仅告诉 AI 你要什么，还要告诉它你不要什么。对于错误处理模式、命名规范和代码风格，这种对比方式尤其有效，因为 AI 很容易默认落入常见但不符合你要求的写法。"},
			{Kind: types.Code, Content: "  ✗ Don't want:  if err != nil { log.Fatal(err) }\n  ✓ Do want:      if err != nil {\n                      return fmt.Errorf(\"fetch user %d: %w\", id, err)\n                  }\n  Apply this error handling pattern throughout."},
			{Kind: types.Heading, Content: "4. Test-Driven Prompting"},
			{Kind: types.Paragraph, Content: "先给测试，再要求 AI 写出能通过测试的代码。测试本身就是可执行规格，它几乎不留行为歧义。这是获得正确实现的高收益技巧之一。"},
			{Kind: types.Code, Content: "  \"Write a function that passes these tests:\n   assert slugify('Hello World') == 'hello-world'\n   assert slugify('  Spaces  ') == 'spaces'\n   assert slugify('Special!@#') == 'special'\n   assert slugify('UPPER') == 'upper'\""},
			{Kind: types.Heading, Content: "5. Chain-of-Thought"},
			{Kind: types.Paragraph, Content: "让 AI 在写代码前先把问题想清楚。这能帮助它做出更好的架构判断，并在实现前就暴露潜在问题。对于设计选择、复杂算法和权衡分析尤其有价值。"},
			{Kind: types.Code, Content: "  \"Before writing code, think through:\n   1. Compare token bucket vs sliding window for\n      our use case (bursty traffic, per-user limits)\n   2. Explain which approach you'd choose and why\n   3. Then implement it\""},
			{Kind: types.Heading, Content: "6. Tree-of-Thought"},
			{Kind: types.Paragraph, Content: "这是 Chain-of-Thought 的扩展版：不是只沿一条推理路径前进，而是让 AI 并行探索多个方案、逐个评估，然后挑出最优解。对于没有唯一标准答案的架构决策尤其强大。"},
			{Kind: types.Code, Content: "  \"Imagine three expert developers each proposing\n   a different approach to this caching system.\n   Each expert writes one paragraph explaining their\n   design. Then they critique each other's approaches.\n   Finally, pick the strongest approach and implement it.\""},
			{Kind: types.Paragraph, Content: "Tree-of-Thought 之所以有效，是因为它强迫 AI 在做出承诺之前先考虑替代方案。Chain-of-Thought 只走一条路径；Tree-of-Thought 会分支、评估、回退，像资深工程师在白板上同时推演多个方案。"},
			{Kind: types.Heading, Content: "7. Scaffold-Then-Detail"},
			{Kind: types.Paragraph, Content: "面对大功能时，先让 AI 搭好骨架，例如类型定义、函数签名、带 stub 的文档注释；然后再通过后续提示逐个补全函数实现。这样每次响应都更聚焦，也更容易审查。"},
			{Kind: types.Callout, Content: "技巧要和场景匹配：Persona 用于专家视角，Few-shot 用于模式匹配，负例用于约束风格，测试用于规格定义，Chain-of-Thought 用于推理，Tree-of-Thought 用于设计决策，Scaffold 适合大型功能。"},
			{Kind: types.Heading, Content: "延伸阅读"},
			{Kind: types.Callout, Content: "Prompt Engineering Guide — https://www.promptingguide.ai"},
			{Kind: types.Callout, Content: "Chain-of-Thought Prompting (Wei et al., 2022) — https://arxiv.org/abs/2201.11903"},
			{Kind: types.Callout, Content: "Tree of Thoughts (Yao et al., 2023) — https://arxiv.org/abs/2305.10601"},
			{Kind: types.Callout, Content: "Few-Shot Prompting — https://www.promptingguide.ai/techniques/fewshot"},
		},
		Questions: []types.QuizQuestion{
			{
				Kind:        types.MultipleChoice,
				Prompt:      "如果你想让 AI 遵循你项目中的错误处理风格，哪种技巧最有效？",
				Choices:     []string{"给它设定专家人格", "给出负例对比：明确展示你不要的写法和你要的写法", "让它一步一步思考", "提供测试用例"},
				CorrectIdx:  1,
				Explanation: "负例/对比例子是控制编码风格最可靠的方式之一。你要明确告诉 AI 什么模式该避免，什么模式才应该遵循。",
			},
			{
				Kind:        types.MultipleChoice,
				Prompt:      "什么是 test-driven prompting？",
				Choices:     []string{"在 AI 生成代码后再去写测试", "先让 AI 自己写测试", "先提供测试用例，再要求 AI 写出能通过它们的代码", "把 AI 写的代码丢进你的测试套件里运行"},
				CorrectIdx:  2,
				Explanation: "Test-driven prompting 的核心是先把测试作为可执行规格给 AI，再让它生成满足这些行为约束的代码。",
			},
			{
				Kind:        types.FillBlank,
				Prompt:      "在提示词中设置一个专家 _______，可以让 AI 在特定领域任务中更聚焦地推理。",
				Answer:      "persona",
				Explanation: "像“资深系统工程师”或“安全专家”这类 persona，会把领域知识和特定推理方式带进 AI 的输出中。",
			},
			{
				Kind:        types.MultipleChoice,
				Prompt:      "Tree-of-Thought 和 Chain-of-Thought 的区别是什么？",
				Choices:     []string{"它使用不同的 AI 模型", "它会并行探索多个方案、逐个评估，再挑选最优方案", "它会跳过推理步骤", "它只适用于数学问题"},
				CorrectIdx:  1,
				Explanation: "Chain-of-Thought 只沿一条路径推理；Tree-of-Thought 会分叉出多个方案、逐个评估和批判，然后再选择最强的方案，类似白板讨论多个备选设计。",
			},
		},
	})
}
