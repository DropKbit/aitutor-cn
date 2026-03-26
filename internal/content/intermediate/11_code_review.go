package intermediate

import (
	"github.com/DropKbit/aitutor-cn/internal/lesson"
	"github.com/DropKbit/aitutor-cn/internal/viz"
	"github.com/DropKbit/aitutor-cn/pkg/types"
)

func init() {
	lesson.Register(types.LessonDef{
		ID:         11,
		Title:      "AI 代码审查",
		Tier:       types.Intermediate,
		Summary:    "识别 AI 生成代码中的常见错误",
		SourceFile: "internal/content/intermediate/11_code_review.go",
		VizBuilder: func(w, h int) interface{} { return viz.NewBugHunterModel(w, h) },
		Theory: []types.TheoryBlock{
			{Kind: types.Heading, Content: "自信陷阱"},
			{Kind: types.Paragraph, Content: "AI 生成的代码常常看起来非常“像那么回事”：排版整洁、变量名合理、逻辑似乎也讲得通。但这种 polished 外表会掩盖很多微妙错误，而这些错误在人类写的粗糙草稿中反而更容易被发现。代码可能“看起来对”，却在边界条件上失败、调用了不存在的 API，或者藏着反向逻辑。"},
			{Kind: types.Heading, Content: "AI 常见错误"},
			{Kind: types.Bullet, Content: "循环和切片中的 off-by-one 错误，尤其是边界条件\n错误处理不当：吞错、错误类型不对、遗漏清理逻辑\n臆造出当前依赖版本并不存在的方法或 API\n使用过时模式：废弃 API、旧库版本、已移除特性\n微妙逻辑错误：运算符优先级错误、漏掉 nil 检查、条件反转\n并发处理错误：缺少同步、竞态条件、死锁"},
			{Kind: types.Heading, Content: "Bug 发现清单"},
			{Kind: types.Bullet, Content: "逐行阅读，而不只是看整体结构\n确认所有函数/方法名都真实存在于依赖中\n检查循环边界与 off-by-one 条件\n沿着错误路径推演：每一步失败后会发生什么？\n查找缺失的 nil/null/空值检查\n确认代码与你当前语言版本和库版本兼容"},
			{Kind: types.Heading, Content: "验证策略"},
			{Kind: types.Paragraph, Content: "可以要求 AI 解释它自己写的代码：“逐步带我过一遍这个函数。”如果解释与代码实际行为对不上，那通常就是 bug。接受生成代码之前先写测试；提交之前用边界输入实际跑一遍。"},
			{Kind: types.Code, Content: "  Verification workflow:\n  1. Read every line of the generated code\n  2. Check that imported methods actually exist\n  3. Run tests with edge-case inputs\n  4. Check error paths and nil handling\n  5. Only then: commit the code"},
			{Kind: types.Heading, Content: "示例：找出 bug"},
			{Kind: types.Code, Content: "  func isValidAge(age int) bool {\n      if age < 0 || age > 150 {\n          return true   // ← Bug! Returns true for INVALID ages\n      }\n      return false\n  }"},
			{Kind: types.Paragraph, Content: "这个函数在年龄超出范围时返回 true，这与 isValid 这个名字表达的含义正好相反。代码读起来很自然，也能通过编译，但逻辑完全反了。AI 非常容易犯这种“看起来顺、实际反”的错误。"},
			{Kind: types.Heading, Content: "建立代码审查习惯"},
			{Kind: types.Bullet, Content: "把 AI 输出当成初级开发者提交的 PR，用同样严格的标准去审查\n没有逐行读过的 AI 代码，绝不要直接提交\n每次 AI 生成改动后，都要运行测试套件\n如果你无法解释某一行在做什么，就不要保留它\n也可以把 AI 当第二审查者：“帮我找出这段代码里的 bug”"},
			{Kind: types.Callout, Content: "最危险的 AI bug 不是会崩溃的那些，而是会悄悄产出错误结果的那些。少一个边界检查、条件反转、并发竞态：它们都能顺利编译，往往只有在线上才会暴露。"},
		},
		Questions: []types.QuizQuestion{
			{
				Kind:        types.MultipleChoice,
				Prompt:      "为什么 AI 生成的 bug 往往比普通 bug 更难发现？",
				Choices:     []string{"AI bug 一定会触发编译错误", "AI 写的代码格式很差", "AI 代码看起来很 polished、很权威，从而掩盖了细微错误", "AI bug 只会在线上出现"},
				CorrectIdx:  2,
				Explanation: "AI 常会产出排版整洁、看起来很合理的代码，这种“表面正确”会让人很容易略过其中细微的逻辑错误。",
			},
			{
				Kind:        types.Ordering,
				Prompt:      "请将以下代码审查步骤按最有效的顺序排列：",
				Choices:     []string{"逐行阅读生成的代码", "确认导入的方法真实存在", "使用边界输入运行测试", "提交代码"},
				CorrectIdx:  0,
				Explanation: "先完整阅读，再核对 API，随后跑边界测试，最后在验证通过后再提交，这是最稳妥的顺序。",
			},
			{
				Kind:        types.MultipleChoice,
				Prompt:      "哪类 AI 代码错误最危险，因为它可能只会在线上暴露？",
				Choices:     []string{"语法错误", "缺少 import 语句", "并发代码中的竞态条件", "变量名写错"},
				CorrectIdx:  2,
				Explanation: "竞态条件往往可以顺利编译，也可能通过大多数测试，但会在并发压力下才显现，因此在线上环境特别危险。",
			},
			{
				Kind:        types.FillBlank,
				Prompt:      "你应该把 AI 生成的代码当成一位 _______ 开发者提交的 PR 来审查。",
				Answer:      "junior",
				Explanation: "用审查初级开发者 PR 的严格标准来审查 AI 输出，才能避免被其 polished 外表迷惑，及时发现隐藏错误。",
			},
		},
	})
}
