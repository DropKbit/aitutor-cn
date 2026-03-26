package beginner

import (
	"github.com/DropKbit/aitutor-cn/internal/lesson"
	"github.com/DropKbit/aitutor-cn/internal/viz"
	"github.com/DropKbit/aitutor-cn/pkg/types"
)

func init() {
	lesson.Register(types.LessonDef{
		ID:         4,
		Title:      "提示词工程",
		Tier:       types.Beginner,
		Summary:    "为 AI 助手编写高质量指令",
		SourceFile: "internal/content/beginner/04_prompt_engineering.go",
		VizBuilder: func(w, h int) interface{} { return viz.NewPromptImproveModel(w, h) },
		Theory: []types.TheoryBlock{
			{Kind: types.Heading, Content: "提示词工程"},
			{Kind: types.Paragraph, Content: "你从 AI 助手那里得到的结果，很大程度上取决于你如何提问。提示词工程就是设计高质量指令、从而获得更好结果的实践。"},
			{Kind: types.Heading, Content: "好提示词的原则"},
			{Kind: types.Bullet, Content: "具体明确：例如“修复有效 token 仍返回 401 的登录问题”，而不是“修 bug”\n提供上下文：说明相关文件、报错信息和期望行为\n说明目标：解释你想达到什么结果，而不仅是改哪里\n设置约束：例如“不要修改公开 API”或“沿用现有模式”"},
			{Kind: types.Heading, Content: "前后对比"},
			{Kind: types.Code, Content: "  ✗ 不好：\"make it work\"\n  ✓ 更好：\"The UserService.GetByEmail method returns nil\n           instead of an error when the database query\n           fails. Fix it to return a wrapped error.\""},
			{Kind: types.Code, Content: "  ✗ 不好：\"add tests\"\n  ✓ 更好：\"Add unit tests for the ParseConfig function\n           covering: valid YAML, missing required fields,\n           and invalid port numbers.\""},
			{Kind: types.Heading, Content: "迭代优化"},
			{Kind: types.Paragraph, Content: "你不需要一开始就写出完美提示词。先给出清晰请求，再根据结果继续细化。AI 会记住当前对话上下文，所以你可以自然地纠偏和补充。"},
			{Kind: types.Callout, Content: "最好的提示词，通常就是你会给资深同事提供的那类信息：问题是什么、你试过什么、什么结果算成功？"},
			{Kind: types.Callout, Content: "延伸阅读：Prompt Engineering Guide — https://www.promptingguide.ai"},
		},
		Questions: []types.QuizQuestion{
			{
				Kind:        types.MultipleChoice,
				Prompt:      "以下哪条更适合用于修复 bug 的提示词？",
				Choices:     []string{"\"fix the bug\"", "\"make it work\"", "\"Fix the 401 error in LoginHandler when valid JWT tokens are rejected\"", "\"debug the code\""},
				CorrectIdx:  2,
				Explanation: "带有明确上下文的信息（问题是什么、出现在哪里、何时发生）远比模糊请求有效。",
			},
			{
				Kind:        types.MultipleChoice,
				Prompt:      "在提示词工程中，你不应该怎么做？",
				Choices:     []string{"提供错误上下文", "说明你的约束条件", "期望第一次就完美无误", "提及相关文件路径"},
				CorrectIdx:  2,
				Explanation: "提示词工程是一个迭代过程：先表达清楚，再看结果，最后持续优化。",
			},
		},
	})
}
