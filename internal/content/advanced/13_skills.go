package advanced

import (
	"github.com/DropKbit/aitutor-cn/internal/lesson"
	"github.com/DropKbit/aitutor-cn/internal/viz"
	"github.com/DropKbit/aitutor-cn/pkg/types"
)

func init() {
	lesson.Register(types.LessonDef{
		ID:         13,
		Title:      "技能（Skills）",
		Tier:       types.Advanced,
		Summary:    "理解可复用工作流与专门知识的封装方式",
		SourceFile: "internal/content/advanced/13_skills.go",
		VizBuilder: func(w, h int) interface{} { return viz.NewSkillLoadModel(w, h) },
		Theory: []types.TheoryBlock{
			{Kind: types.Heading, Content: "技能（Skills）"},
			{Kind: types.Paragraph, Content: "技能是可复用、可组合的知识与工作流包，用来扩展 AI 助手的能力。它们提供特定领域的经验、分步流程，以及针对常见任务的约束与护栏。"},
			{Kind: types.Heading, Content: "技能如何工作"},
			{Kind: types.Paragraph, Content: "技能是按需加载的，它们不会一直驻留在上下文窗口中。当某个任务匹配到技能的适用领域时，技能才会被调用，其指令内容才会被加载进上下文。这样既节省上下文空间，又能在需要时提供深度专业能力。"},
			{Kind: types.Code, Content: "  User: \"Create a new MCP server\"\n         │\n         ▼\n  ┌──────────────────┐\n  │ Skill Detection  │\n  │ \"mcp-builder\"    │\n  └────────┬─────────┘\n           ▼\n  ┌──────────────────┐\n  │  Load Skill      │\n  │  Instructions    │\n  └────────┬─────────┘\n           ▼\n  ┌──────────────────┐\n  │ Follow Workflow  │\n  │ Step by Step     │\n  └──────────────────┘"},
			{Kind: types.Heading, Content: "技能类型"},
			{Kind: types.Bullet, Content: "刚性技能：必须严格遵守，例如 TDD、调试工作流\n柔性技能：原则可按上下文调整，例如设计模式\n流程技能：定义“如何处理任务”，例如 brainstorming、planning\n实现技能：指导具体执行，例如 frontend-design、mcp-builder"},
			{Kind: types.Heading, Content: "技能组合"},
			{Kind: types.Paragraph, Content: "技能之间还可以互相调用。例如，一个“构建功能”的工作流，可能会先调用 brainstorming，再调用 planning，然后再进入 test-driven development；每个技能都为当前阶段提供专门能力。"},
			{Kind: types.Callout, Content: "技能把原本依赖口口相传的经验，变成可重复执行的工作流。与其记住复杂流程，不如把它们编码为技能，并在每次任务中稳定执行。"},
		},
		Questions: []types.QuizQuestion{
			{
				Kind:        types.MultipleChoice,
				Prompt:      "为什么技能要按需加载，而不是始终放在上下文里？",
				Choices:     []string{"因为加载太慢", "为了节省上下文窗口空间", "因为它们需要特殊权限", "因为它们还是实验功能"},
				CorrectIdx:  1,
				Explanation: "按需加载可以保持上下文窗口高效，只在真正需要时才引入深度专业知识。",
			},
			{
				Kind:        types.MultipleChoice,
				Prompt:      "以下哪类技能必须严格执行，不能随意改写？",
				Choices:     []string{"柔性技能", "模式技能", "刚性技能", "实现技能"},
				CorrectIdx:  2,
				Explanation: "像 TDD 和调试流程这样的刚性技能，价值本身就在于纪律性，因此必须严格遵守。",
			},
		},
	})
}
