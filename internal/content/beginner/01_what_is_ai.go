package beginner

import (
	"github.com/DropKbit/aitutor-cn/internal/lesson"
	"github.com/DropKbit/aitutor-cn/internal/viz"
	"github.com/DropKbit/aitutor-cn/pkg/types"
)

func init() {
	lesson.Register(types.LessonDef{
		ID:         1,
		Title:      "什么是 AI 编程助手？",
		Tier:       types.Beginner,
		Summary:    "理解 AI 助手如何帮助你编写和修改代码",
		SourceFile: "internal/content/beginner/01_what_is_ai.go",
		VizBuilder: func(w, h int) interface{} { return viz.NewAgentLoopModel(w, h) },
		Theory: []types.TheoryBlock{
			{Kind: types.Heading, Content: "什么是 AI 编程助手？"},
			{Kind: types.Paragraph, Content: "AI 编程助手是由大语言模型（LLM）驱动的工具，能够帮助开发者编写、理解和修改代码。与传统自动补全或 linter 不同，它能理解自然语言指令，并在较高层次上对代码进行推理。"},
			{Kind: types.Heading, Content: "代理循环（Agent Loop）"},
			{Kind: types.Paragraph, Content: "现代 AI 编程助手通常运行在“代理循环”中：先读取上下文，再推理下一步要做什么，随后执行动作（如编辑文件或运行命令），最后观察结果。这个循环会一直持续，直到任务完成。"},
			{Kind: types.Code, Content: "  ┌──────────────────┐\n  │    用户请求      │\n  └────────┬─────────┘\n           ▼\n  ┌─────────────────┐\n  │   读取上下文    │◄──────┐\n  └────────┬────────┘       │\n           ▼                │\n  ┌─────────────────┐       │\n  │    推理与规划   │       │\n  └────────┬────────┘       │\n           ▼                │\n  ┌─────────────────┐       │\n  │    执行动作     │       │\n  │   （工具调用）  │       │\n  └────────┬────────┘       │\n           ▼                │\n  ┌─────────────────┐       │\n  │    观察结果     │───────┘\n  └────────┬────────┘\n           ▼\n  ┌─────────────────┐\n  │      响应       │\n  └─────────────────┘"},
			{Kind: types.Heading, Content: "核心能力"},
			{Kind: types.Bullet, Content: "代码生成：根据自然语言描述编写新代码\n代码修改：对现有代码进行精确更改\n代码解释：理解并讲解复杂代码库\n缺陷修复：定位并修复代码问题\n重构：在保持行为不变的前提下优化结构\n测试编写：为代码生成测试"},
			{Kind: types.Heading, Content: "它是工具，不是魔法"},
			{Kind: types.Paragraph, Content: "AI 助手的作用是增强你的能力。只有在你提供清晰上下文、认真审查输出，并在它偏离方向时及时引导时，它才会发挥最佳效果。理解其工作方式，也正是本教程的目标，会显著提升你的使用效率。"},
			{Kind: types.Callout, Content: "高效的开发者不只是“会用”AI 助手，他们还理解它在底层是如何工作的。这正是本教程要教你的内容。"},
			{Kind: types.Callout, Content: "延伸阅读：Large Language Models — https://en.wikipedia.org/wiki/Large_language_model"},
		},
		Questions: []types.QuizQuestion{
			{
				Kind:        types.MultipleChoice,
				Prompt:      "AI 编程助手中的“代理循环”指的是什么？",
				Choices:     []string{"一种编程语言特性", "读取上下文 → 推理 → 行动 → 观察的循环", "一种无限循环 bug", "一种用户界面模式"},
				CorrectIdx:  1,
				Explanation: "代理循环是 AI 助手的核心流程：读取上下文、推理下一步、执行动作并观察结果。",
			},
			{
				Kind:        types.MultipleChoice,
				Prompt:      "以下哪一项不是 AI 编程助手的核心能力？",
				Choices:     []string{"代码生成", "缺陷修复", "完全取代开发者", "重构"},
				CorrectIdx:  2,
				Explanation: "AI 助手的作用是增强开发者能力，它是工具，而不是替代者。",
			},
		},
	})
}
