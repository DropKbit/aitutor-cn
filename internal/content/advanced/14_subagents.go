package advanced

import (
	"github.com/DropKbit/aitutor-cn/internal/lesson"
	"github.com/DropKbit/aitutor-cn/internal/viz"
	"github.com/DropKbit/aitutor-cn/pkg/types"
)

func init() {
	lesson.Register(types.LessonDef{
		ID:         14,
		Title:      "子代理（Subagents）",
		Tier:       types.Advanced,
		Summary:    "理解如何用专门化子进程并行处理任务",
		SourceFile: "internal/content/advanced/14_subagents.go",
		VizBuilder: func(w, h int) interface{} { return viz.NewFanoutModel(w, h) },
		Theory: []types.TheoryBlock{
			{Kind: types.Heading, Content: "子代理（Subagents）"},
			{Kind: types.Paragraph, Content: "子代理是由主代理派生出来的专门 AI 进程，用来并行处理特定子任务。它们能让多个独立问题同时推进，从而显著加快复杂任务的处理速度。"},
			{Kind: types.Heading, Content: "子代理如何工作"},
			{Kind: types.Code, Content: "  ┌─────────────┐\n  │  Main Agent  │\n  │  (orchestr.) │\n  └──────┬──────┘\n    ┌────┼────┐\n    ▼    ▼    ▼\n  ┌───┐┌───┐┌───┐\n  │ A ││ B ││ C │  ← parallel execution\n  └─┬─┘└─┬─┘└─┬─┘\n    └────┼────┘\n         ▼\n  ┌─────────────┐\n  │  Main Agent  │\n  │  (combines)  │\n  └─────────────┘"},
			{Kind: types.Heading, Content: "子代理类型"},
			{Kind: types.Bullet, Content: "Explore：快速探索代码库，只读\nGeneral-purpose：拥有完整工具访问能力，适合复杂任务\nPlan：进行架构规划，只读\nSpecialized：为特定工作流定制的代理，例如 code-review、test-runner"},
			{Kind: types.Heading, Content: "何时使用子代理"},
			{Kind: types.Bullet, Content: "需要回答多个彼此独立的研究问题时\n要在不同区域并行搜索文件时\n需要同时实现互不相关的改动时\n希望一边改代码一边跑测试时"},
			{Kind: types.Heading, Content: "隔离性"},
			{Kind: types.Paragraph, Content: "子代理可以运行在隔离的 git worktree 中，也就是仓库的独立副本。这样它们就能在不影响主工作区的前提下进行改动，从而避免多个代理同时编辑文件时发生冲突。"},
			{Kind: types.Callout, Content: "可以把子代理想象成你能瞬间拉起的团队成员：每个成员拿到清晰任务，独立工作，最后再把结果汇报回来。"},
		},
		Questions: []types.QuizQuestion{
			{
				Kind:        types.MultipleChoice,
				Prompt:      "如果你想快速探索代码库，最合适的子代理类型是什么？",
				Choices:     []string{"General-purpose", "Plan", "Explore", "Specialized"},
				CorrectIdx:  2,
				Explanation: "Explore 代理是专门用来搜索和浏览代码库的快速只读代理。",
			},
			{
				Kind:        types.MultipleChoice,
				Prompt:      "子代理在并行工作时通常如何避免合并冲突？",
				Choices:     []string{"它们会锁文件", "它们会使用隔离的 git worktree", "它们会轮流工作", "做不到，冲突不可避免"},
				CorrectIdx:  1,
				Explanation: "子代理可以运行在彼此隔离的 git worktree 中，从而把改动分散到不同副本里，减少直接冲突。",
			},
		},
	})
}
