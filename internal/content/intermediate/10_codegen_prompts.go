package intermediate

import (
	"github.com/naorpeled/aitutor/internal/lesson"
	"github.com/naorpeled/aitutor/internal/viz"
	"github.com/naorpeled/aitutor/pkg/types"
)

func init() {
	lesson.Register(types.LessonDef{
		ID:         10,
		Title:      "Advanced Prompting Techniques",
		Tier:       types.Intermediate,
		Summary:    "Seven techniques that dramatically improve AI-generated code",
		SourceFile: "internal/content/intermediate/10_codegen_prompts.go",
		VizBuilder: func(w, h int) interface{} { return viz.NewPromptBuilderModel(w, h) },
		Theory: []types.TheoryBlock{
			{Kind: types.Heading, Content: "Seven Techniques That Actually Work"},
			{Kind: types.Paragraph, Content: "Lesson 4 covered the basics: be specific, provide context. This lesson goes deeper with seven proven techniques that dramatically improve code generation. Each one is backed by real-world best practices from AI labs and developer communities."},
			{Kind: types.Heading, Content: "1. Persona / Role Assignment"},
			{Kind: types.Paragraph, Content: "Setting an expert role in your prompt focuses the AI's reasoning and brings domain-specific knowledge to bear. A \"senior systems engineer\" reasons differently about caching than a generic assistant. Even a single sentence makes a difference."},
			{Kind: types.Code, Content: "  \"You are a senior Go developer specializing in\n   concurrent systems. Implement a worker pool that\n   processes jobs from a channel. It must be\n   thread-safe, handle panics in workers, and support\n   configurable pool size.\""},
			{Kind: types.Heading, Content: "2. Few-Shot Examples"},
			{Kind: types.Paragraph, Content: "Show 2-3 examples of what you want, and the AI matches the pattern. This is the most reliable way to get code that follows your project's conventions — naming, error handling, structure, all of it."},
			{Kind: types.Code, Content: "  \"Here are two existing handlers in our codebase:\n   [paste handler A]\n   [paste handler B]\n   Following the exact same patterns (error handling,\n   response format, naming), create a handler for\n   DELETE /users/:id.\""},
			{Kind: types.Heading, Content: "3. Negative / Contrastive Examples"},
			{Kind: types.Paragraph, Content: "Show the AI what you DON'T want alongside what you DO want. Especially effective for error handling patterns, naming conventions, and coding style where the AI might default to common but undesired patterns."},
			{Kind: types.Code, Content: "  ✗ Don't want:  if err != nil { log.Fatal(err) }\n  ✓ Do want:      if err != nil {\n                      return fmt.Errorf(\"fetch user %d: %w\", id, err)\n                  }\n  Apply this error handling pattern throughout."},
			{Kind: types.Heading, Content: "4. Test-Driven Prompting"},
			{Kind: types.Paragraph, Content: "Provide the tests first, then ask for code that passes them. Tests act as an executable specification with zero ambiguity about expected behavior. This is one of the highest-impact techniques for getting correct implementations."},
			{Kind: types.Code, Content: "  \"Write a function that passes these tests:\n   assert slugify('Hello World') == 'hello-world'\n   assert slugify('  Spaces  ') == 'spaces'\n   assert slugify('Special!@#') == 'special'\n   assert slugify('UPPER') == 'upper'\""},
			{Kind: types.Heading, Content: "5. Chain-of-Thought"},
			{Kind: types.Paragraph, Content: "Ask the AI to reason through the problem before coding. This produces better architectural decisions and catches issues upfront. Especially valuable for design choices, complex algorithms, and trade-off analysis."},
			{Kind: types.Code, Content: "  \"Before writing code, think through:\n   1. Compare token bucket vs sliding window for\n      our use case (bursty traffic, per-user limits)\n   2. Explain which approach you'd choose and why\n   3. Then implement it\""},
			{Kind: types.Heading, Content: "6. Tree-of-Thought"},
			{Kind: types.Paragraph, Content: "An extension of chain-of-thought: instead of following one reasoning path, ask the AI to explore multiple approaches in parallel, evaluate each, and pick the best. This is powerful for architectural decisions where there's no obvious right answer."},
			{Kind: types.Code, Content: "  \"Imagine three expert developers each proposing\n   a different approach to this caching system.\n   Each expert writes one paragraph explaining their\n   design. Then they critique each other's approaches.\n   Finally, pick the strongest approach and implement it.\""},
			{Kind: types.Paragraph, Content: "Tree-of-Thought works because it forces the AI to consider alternatives before committing. Chain-of-thought follows one path; tree-of-thought branches, evaluates, and backtracks — like a senior engineer whiteboarding multiple options."},
			{Kind: types.Heading, Content: "7. Scaffold-Then-Detail"},
			{Kind: types.Paragraph, Content: "For large features, ask for the skeleton first — type definitions, function signatures, doc comments with stub implementations. Then fill in one function at a time in follow-up prompts. Each response stays focused and reviewable."},
			{Kind: types.Callout, Content: "Match the technique to the situation: Persona for expertise. Few-shot for pattern matching. Negative examples for style. Tests for specs. Chain-of-thought for reasoning. Tree-of-thought for design decisions. Scaffold for large features."},
			{Kind: types.Heading, Content: "Further Reading"},
			{Kind: types.Callout, Content: "Prompt Engineering Guide — https://www.promptingguide.ai"},
			{Kind: types.Callout, Content: "Chain-of-Thought Prompting (Wei et al., 2022) — https://arxiv.org/abs/2201.11903"},
			{Kind: types.Callout, Content: "Tree of Thoughts (Yao et al., 2023) — https://arxiv.org/abs/2305.10601"},
			{Kind: types.Callout, Content: "Few-Shot Prompting — https://www.promptingguide.ai/techniques/fewshot"},
		},
		Questions: []types.QuizQuestion{
			{
				Kind:        types.MultipleChoice,
				Prompt:      "You want the AI to match your project's error handling style. Which technique is most effective?",
				Choices:     []string{"Assign it an expert persona", "Show negative examples: what you DON'T want vs what you DO want", "Ask it to think step by step", "Provide test cases"},
				CorrectIdx:  1,
				Explanation: "Negative/contrastive examples are the most reliable way to control coding style. Show the pattern to avoid and the pattern to follow.",
			},
			{
				Kind:        types.MultipleChoice,
				Prompt:      "What is test-driven prompting?",
				Choices:     []string{"Writing tests after the AI generates code", "Asking the AI to write tests first", "Providing test cases first, then asking for code that passes them", "Running the AI's code against your test suite"},
				CorrectIdx:  2,
				Explanation: "Test-driven prompting gives the AI your tests upfront as an executable specification. The AI generates code that satisfies exactly the behavior you defined.",
			},
			{
				Kind:        types.FillBlank,
				Prompt:      "Setting an expert _______ in your prompt focuses the AI's reasoning for domain-specific tasks.",
				Answer:      "persona",
				Explanation: "A persona like 'senior systems engineer' or 'security specialist' brings domain-specific knowledge and reasoning patterns to the AI's output.",
			},
			{
				Kind:        types.MultipleChoice,
				Prompt:      "How does tree-of-thought differ from chain-of-thought?",
				Choices:     []string{"It uses a different AI model", "It explores multiple approaches in parallel, evaluates each, and picks the best", "It skips the reasoning step entirely", "It only works for mathematical problems"},
				CorrectIdx:  1,
				Explanation: "Chain-of-thought follows one reasoning path. Tree-of-thought branches into multiple approaches, evaluates and critiques each, then commits to the strongest — like whiteboarding multiple options.",
			},
		},
	})
}
