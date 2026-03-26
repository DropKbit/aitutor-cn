package intermediate

import (
	"github.com/DropKbit/aitutor-cn/internal/lesson"
	"github.com/DropKbit/aitutor-cn/internal/viz"
	"github.com/DropKbit/aitutor-cn/pkg/types"
)

func init() {
	lesson.Register(types.LessonDef{
		ID:         8,
		Title:      "记忆与持久化",
		Tier:       types.Intermediate,
		Summary:    "理解 AI 助手如何跨会话保留信息",
		SourceFile: "internal/content/intermediate/08_memory.go",
		VizBuilder: func(w, h int) interface{} { return viz.NewMemorySortModel(w, h) },
		Theory: []types.TheoryBlock{
			{Kind: types.Heading, Content: "记忆与持久化"},
			{Kind: types.Paragraph, Content: "默认情况下，每次 AI 会话都会从零开始，模型并不会记得之前的对话。但很多机制都可以让知识在会话之间保留下来。"},
			{Kind: types.Heading, Content: "记忆层级"},
			{Kind: types.Code, Content: "  ┌─────────────────────────────────────┐\n  │            会话记忆                 │\n  │     （会话上下文，临时存在）        │\n  ├─────────────────────────────────────┤\n  │          自动记忆文件               │\n  │      （工具管理的持久笔记）         │\n  ├─────────────────────────────────────┤\n  │          项目配置文件               │\n  │     （AGENTS.md、CLAUDE.md 等）     │\n  ├─────────────────────────────────────┤\n  │            用户设置                 │\n  │       （工具级用户偏好）            │\n  └─────────────────────────────────────┘"},
			{Kind: types.Heading, Content: "自动记忆"},
			{Kind: types.Paragraph, Content: "有些 AI 编程工具支持自动记忆功能，允许 AI 把笔记写入持久文件，并在后续会话中再次加载。它适合保存经过多次交互确认的内容，例如模式、约定、调试经验和用户偏好。"},
			{Kind: types.Heading, Content: "什么该记，什么不该记"},
			{Kind: types.Bullet, Content: "应该记：稳定模式、架构决策、用户偏好、反复出现的解决方案\n不该记：会话专属细节、进行中的工作、尚未验证的推测"},
			{Kind: types.Callout, Content: "在支持持久记忆的工具里，你可以直接告诉 AI 记住某件事，例如“始终使用 bun 而不是 npm”，它会把这个偏好保存到后续会话中。"},
		},
		Questions: []types.QuizQuestion{
			{
				Kind:        types.MultipleChoice,
				Prompt:      "哪一层记忆是最持久的？",
				Choices:     []string{"会话记忆", "自动记忆文件", "项目配置文件（如 AGENTS.md）", "对话历史"},
				CorrectIdx:  2,
				Explanation: "像 AGENTS.md 这样的项目配置文件是最持久的，因为它们通常纳入版本控制，并且会在每次会话中自动加载。",
			},
			{
				Kind:        types.MultipleChoice,
				Prompt:      "以下哪类信息不应该写入自动记忆？",
				Choices:     []string{"稳定的编码模式", "架构决策", "当前会话的任务细节", "用户偏好"},
				CorrectIdx:  2,
				Explanation: "当前任务、进行中的工作这类会话专属信息只是临时上下文，不适合写入持久记忆。",
			},
		},
	})
}
