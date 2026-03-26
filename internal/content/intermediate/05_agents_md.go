package intermediate

import (
	"github.com/DropKbit/aitutor-cn/internal/lesson"
	"github.com/DropKbit/aitutor-cn/internal/viz"
	"github.com/DropKbit/aitutor-cn/pkg/types"
)

func init() {
	lesson.Register(types.LessonDef{
		ID:         5,
		Title:      "项目级 AI 配置文件",
		Tier:       types.Intermediate,
		Summary:    "理解面向项目的 AI 配置文件及其作用",
		SourceFile: "internal/content/intermediate/05_agents_md.go",
		VizBuilder: func(w, h int) interface{} { return viz.NewClaudeMDBuilderModel(w, h) },
		Theory: []types.TheoryBlock{
			{Kind: types.Heading, Content: "项目级 AI 配置文件"},
			{Kind: types.Paragraph, Content: "AI 编程工具会使用项目级配置文件来接收持久指令。这些文件会在每次会话开始时自动加载进上下文窗口，相当于项目的持久记忆。典型例子包括 AGENTS.md（跨工具标准）、CLAUDE.md（Claude Code）、.cursorrules（Cursor）和 copilot-instructions.md（GitHub Copilot）。"},
			{Kind: types.Heading, Content: "为什么重要：证据在哪里"},
			{Kind: types.Paragraph, Content: "研究表明，这类文件确实会产生可量化影响。一项覆盖 10 个仓库、124 个 PR 的研究发现，引入 AGENTS.md 与 28.64% 的运行时中位数下降、16.58% 的输出 token 减少相关，同时任务完成率保持不变（Girgis et al., 2026）。"},
			{Kind: types.Paragraph, Content: "Vercel 还做过一次内部评测，目标是 8 个 Next.js 16 的新 API，这些 API 不在模型训练数据中。没有配置文件时，代理通过率只有 53%；加入一个压缩到 8KB 的 AGENTS.md 文档索引后，通过率提升到了 100%。虽然评测范围较窄，但它清楚说明：结构良好的项目上下文非常有帮助，尤其是面对模型没见过的 API 时。"},
			{Kind: types.Callout, Content: "可以把这些配置文件当成给 AI 助手的入职文档，就像你会在第一天告诉新同事的那些事情。"},
			{Kind: types.Heading, Content: "层级与作用域"},
			{Kind: types.Code, Content: "  示例层级（Claude Code）：\n  ~/.claude/CLAUDE.md          ← 用户级（所有项目）\n  ~/project/.claude/CLAUDE.md  ← 个人项目配置（通常 git ignore）\n  ~/project/CLAUDE.md          ← 项目根目录（提交到仓库）\n  ~/project/src/CLAUDE.md      ← 目录级（局部生效）\n\n  跨工具标准：\n  ~/project/AGENTS.md          ← 多个 AI 工具都可识别"},
			{Kind: types.Paragraph, Content: "大多数工具都支持配置文件层级：用户级设置作用于所有项目，项目级设置作用于整个仓库，目录级设置只作用于局部目录。AGENTS.md 是跨工具标准，被多个 AI 编程工具识别，因此非常适合团队协作。"},
			{Kind: types.Heading, Content: "指令预算：少即是多"},
			{Kind: types.Paragraph, Content: "关键点在于：指令并不是越多越好。Gloaguen 等人在 2025 年的研究发现，过于全面的配置文件反而可能降低任务成功率，同时让推理成本增加 20% 以上。代理确实会忠实执行指令，但过多无关要求会触发更广泛的探索，最终把任务带偏。"},
			{Kind: types.Paragraph, Content: "前沿 LLM 大约能较稳定地遵循 150 到 200 条指令。配置文件中的每个 token 都会在每次请求时被加载，不管它与当前任务是否相关。过大的文件既浪费上下文，也会干扰代理；小而聚焦的文件，才能把更多容量留给真正的任务。"},
			{Kind: types.Heading, Content: "配置文件里应该写什么"},
			{Kind: types.Bullet, Content: "项目简介：用一句话说明这个项目是什么\n构建与测试命令：如何编译、测试、lint\n包管理器：仅在不是默认选项时说明，例如使用 bun 而不是 npm\n代码约定：命名模式、架构决策\n必须遵守/避免事项：项目中特别关键的规则"},
			{Kind: types.Heading, Content: "不应该写什么"},
			{Kind: types.Bullet, Content: "穷举式文件列表：路径很容易漂移，应该描述能力而非死记文件名\n所有你知道的事情：保持精简，必要时逐层展开\n自动生成内容：这类内容通常追求“全”，而不是“准”和“简”\n过时文档：旧信息会直接污染代理的上下文"},
			{Kind: types.Heading, Content: "渐进式披露"},
			{Kind: types.Paragraph, Content: "对于较大的项目，可以把 AGENTS.md 保持为简洁索引，把详细指导拆到单独文档里。Vercel 发现，把 40KB 文档压缩为 8KB 的索引，再配合检索指引，比把所有内容直接内联进去效果更好。"},
			{Kind: types.Code, Content: "  # AGENTS.md（尽量保持精简，约 150 行）\n\n  ## Build\n  - `make build` 用于编译\n  - `make test` 用于运行全部测试\n\n  ## Conventions\n  - 数据库字段使用 snake_case\n  - 所有 API handler 位于 internal/api/\n  - 永远不要提交 .env 文件\n\n  ## Detailed Docs\n  - 语言约定见 docs/TYPESCRIPT.md\n  - 接口模式见 docs/API.md"},
			{Kind: types.Callout, Content: "像维护生产代码一样维护配置文件：谨慎评审新增规则，清理过时条目，不要因为一次事故就反射性加规则。无人维护的配置文件，往往比没有配置文件更糟。"},
			{Kind: types.Heading, Content: "参考资料"},
			{Kind: types.Bullet, Content: "Girgis et al. (2026) — \"AGENTS.md Files and AI Coding Agent Efficiency\" — arxiv.org/abs/2601.20404\nGloaguen et al. (2025) — \"Evaluating AGENTS.md for Coding Agents\" — arxiv.org/abs/2602.11988\nVercel — \"AGENTS.md outperforms skills in our agent evals\" — vercel.com/blog/agents-md-outperforms-skills-in-our-agent-evals\nAI Hero — \"A Complete Guide to AGENTS.md\" — aihero.dev/a-complete-guide-to-agents-md"},
		},
		Questions: []types.QuizQuestion{
			{
				Kind:        types.MultipleChoice,
				Prompt:      "CLAUDE.md 和 AGENTS.md 的区别是什么？",
				Choices:     []string{"CLAUDE.md 给 Claude 用，AGENTS.md 给其他 AI 用", "CLAUDE.md 是工具专属文件，AGENTS.md 是跨工具标准", "二者没有区别", "AGENTS.md 已被弃用"},
				CorrectIdx:  1,
				Explanation: "CLAUDE.md 是 Claude Code 专用配置，并且可存在于多个层级；AGENTS.md 是多个 AI 编程助手都能识别的跨工具标准，适合纳入版本控制。",
			},
			{
				Kind:        types.MultipleChoice,
				Prompt:      "为什么过于全面的配置文件反而会伤害代理表现？",
				Choices:     []string{"因为会让 AI 崩溃", "因为无关指令会触发更广泛的探索，让任务跑偏", "因为会超出文件大小限制", "因为会与 system prompt 冲突"},
				CorrectIdx:  1,
				Explanation: "研究发现，代理会忠实执行指令，但无关要求会促使它进行更广泛的探索，降低成功率，并让推理成本增加 20% 以上。",
			},
		},
	})
}
