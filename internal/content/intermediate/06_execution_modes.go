package intermediate

import (
	"github.com/DropKbit/aitutor-cn/internal/lesson"
	"github.com/DropKbit/aitutor-cn/internal/viz"
	"github.com/DropKbit/aitutor-cn/pkg/types"
)

func init() {
	lesson.Register(types.LessonDef{
		ID:         6,
		Title:      "执行模式",
		Tier:       types.Intermediate,
		Summary:    "理解规划模式与执行模式的区别",
		SourceFile: "internal/content/intermediate/06_execution_modes.go",
		VizBuilder: func(w, h int) interface{} { return viz.NewModePickerModel(w, h) },
		Theory: []types.TheoryBlock{
			{Kind: types.Heading, Content: "执行模式"},
			{Kind: types.Paragraph, Content: "AI 编程助手会根据任务类型运行在不同模式下。其中最常见的两种是规划模式（Plan Mode）和执行模式（Execution Mode），它们分别针对不同工作阶段做了优化。"},
			{Kind: types.Heading, Content: "规划模式"},
			{Kind: types.Paragraph, Content: "在规划模式下，AI 主要负责分析和规划，不直接修改代码。它会读取文件、探索代码库并输出结构化计划。这特别适合那些你想先审查方案、再决定是否实施的复杂任务。"},
			{Kind: types.Bullet, Content: "只读：不修改文件\n输出带步骤的结构化计划\n识别需要改动的文件和潜在风险\n适合架构决策和大型重构"},
			{Kind: types.Heading, Content: "执行模式"},
			{Kind: types.Paragraph, Content: "在执行模式下，AI 会主动进行修改，例如编辑文件、运行命令，并根据结果持续迭代。这是大多数明确任务的默认模式。"},
			{Kind: types.Bullet, Content: "拥有完整工具访问能力：读、写、执行\n会根据结果继续迭代，如测试失败或构建报错\n适合边界清晰、范围聚焦的任务"},
			{Kind: types.Heading, Content: "何时使用哪一种"},
			{Kind: types.Code, Content: "  Plan Mode:                    Execution Mode:\n  ─────────────                 ────────────────\n  \"How should we restructure    \"Add input validation\n   the auth system?\"             to the signup handler\"\n\n  \"What's the best approach     \"Fix the NPE in\n   for adding caching?\"          UserService.java\"\n\n  \"Review this PR's              \"Write tests for\n   architecture\"                 the new endpoint\""},
			{Kind: types.Callout, Content: "一种常见工作流是：先用规划模式设计方案并审阅计划，再切换到执行模式完成实现。"},
		},
		Questions: []types.QuizQuestion{
			{
				Kind:        types.MultipleChoice,
				Prompt:      "面对“我们应该如何重构认证系统？”这类问题时，应使用哪种模式？",
				Choices:     []string{"执行模式", "规划模式", "调试模式", "测试模式"},
				CorrectIdx:  1,
				Explanation: "架构类问题更适合规划模式，因为它会在不做修改的前提下进行只读探索并产出结构化方案。",
			},
			{
				Kind:        types.FillBlank,
				Prompt:      "在规划模式中，AI 运行在 ___-only 模式下（不修改文件）。空格里应填什么单词？",
				Answer:      "read",
				Explanation: "规划模式是 read-only，也就是只读分析和规划，不直接改动代码。",
			},
		},
	})
}
