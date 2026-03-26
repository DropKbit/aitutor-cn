package advanced

import (
	"github.com/DropKbit/aitutor-cn/internal/lesson"
	"github.com/DropKbit/aitutor-cn/internal/viz"
	"github.com/DropKbit/aitutor-cn/pkg/types"
)

func init() {
	lesson.Register(types.LessonDef{
		ID:         15,
		Title:      "Git 工作树（Worktree）",
		Tier:       types.Advanced,
		Summary:    "理解并行开发中的隔离工作区机制",
		SourceFile: "internal/content/advanced/15_git_worktrees.go",
		VizBuilder: func(w, h int) interface{} { return viz.NewWorktreeSimModel(w, h) },
		Theory: []types.TheoryBlock{
			{Kind: types.Heading, Content: "Git Worktree"},
			{Kind: types.Paragraph, Content: "Git worktree 允许你同时检出同一个仓库的多个分支，并把它们放在不同目录中。对于 AI 辅助开发来说，这意味着真正的并行工作成为可能，而且更容易避免冲突。"},
			{Kind: types.Heading, Content: "Worktree 如何工作"},
			{Kind: types.Code, Content: "  ~/project/           ← main worktree (main branch)\n  ~/project-worktrees/\n    ├── feature-auth/   ← worktree (feature/auth branch)\n    ├── fix-bug-123/    ← worktree (fix/bug-123 branch)\n    └── refactor-api/   ← worktree (refactor/api branch)"},
			{Kind: types.Paragraph, Content: "每个 worktree 都是完整的工作副本，拥有自己的分支、暂存区和工作目录。它们共享同一份 .git 数据，因此创建起来很轻量，也非常快。"},
			{Kind: types.Heading, Content: "Worktree + AI 代理"},
			{Kind: types.Bullet, Content: "子代理可以拿到隔离的 worktree，从而减少合并冲突\n主工作区可以保持干净，而多个代理并行推进\n可以在合并前按 worktree 分别审查改动\n代理任务完成后可自动清理 worktree"},
			{Kind: types.Heading, Content: "常用命令"},
			{Kind: types.Code, Content: "  # 创建新的 worktree\n  git worktree add ../feature-x -b feature/x\n\n  # 列出所有 worktree\n  git worktree list\n\n  # 删除一个 worktree\n  git worktree remove ../feature-x"},
			{Kind: types.Heading, Content: "最佳实践"},
			{Kind: types.Bullet, Content: "把 worktree 放在兄弟目录中，而不是仓库内部\nworktree 名称尽量与分支名对应，便于识别\n分支合并后及时清理 worktree\n避免两个 worktree 指向同一个分支"},
			{Kind: types.Callout, Content: "worktree 是安全并行 AI 开发的关键基础设施，它为每个代理提供了独立沙箱。"},
			{Kind: types.Callout, Content: "延伸阅读：Git Worktrees — https://git-scm.com/docs/git-worktree"},
		},
		Questions: []types.QuizQuestion{
			{
				Kind:        types.MultipleChoice,
				Prompt:      "多个 git worktree 之间共享的是什么？",
				Choices:     []string{"工作目录", "暂存区", ".git 数据", "分支名称"},
				CorrectIdx:  2,
				Explanation: "worktree 共享同一份 .git 数据（对象库、refs 等），这也是它们足够轻量的原因。每个 worktree 仍有自己的工作目录、暂存区和分支。",
			},
			{
				Kind:        types.FillBlank,
				Prompt:      "创建新 git worktree 的命令是什么？（以 'git' 开头）",
				Answer:      "git worktree add",
				Explanation: "使用 'git worktree add <path> -b <branch>' 可以创建新的 worktree。",
			},
		},
	})
}
