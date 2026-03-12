package intermediate

import (
	"github.com/naorpeled/aitutor/internal/lesson"
	"github.com/naorpeled/aitutor/internal/viz"
	"github.com/naorpeled/aitutor/pkg/types"
)

func init() {
	lesson.Register(types.LessonDef{
		ID:         11,
		Title:      "Reviewing AI-Generated Code",
		Tier:       types.Intermediate,
		Summary:    "Spotting common mistakes in AI-generated code",
		SourceFile: "internal/content/intermediate/11_code_review.go",
		VizBuilder: func(w, h int) interface{} { return viz.NewBugHunterModel(w, h) },
		Theory: []types.TheoryBlock{
			{Kind: types.Heading, Content: "The Confidence Trap"},
			{Kind: types.Paragraph, Content: "AI-generated code looks authoritative — clean formatting, reasonable variable names, plausible logic. This polish masks subtle errors that a rough human draft wouldn't have. The code passes a \"looks right\" test but may fail on edge cases, use non-existent APIs, or contain inverted logic."},
			{Kind: types.Heading, Content: "Common AI Mistakes"},
			{Kind: types.Bullet, Content: "Off-by-one errors in loops and slice operations, especially boundary conditions\nIncorrect error handling — swallowing errors, wrong error types, missing cleanup\nHallucinated methods or APIs that don't exist in your library version\nOutdated patterns — deprecated APIs, old library versions, removed features\nSubtle logic errors — wrong operator precedence, missed nil checks, inverted conditions\nIncorrect concurrency — missing synchronization, race conditions, deadlocks"},
			{Kind: types.Heading, Content: "The Bug Spotter's Checklist"},
			{Kind: types.Bullet, Content: "Read every line, not just the structure\nCheck all function/method names actually exist in your dependencies\nVerify loop bounds and off-by-one conditions\nTrace error paths: what happens when each operation fails?\nLook for missing nil/null/empty checks\nConfirm the code matches your language version and library versions"},
			{Kind: types.Heading, Content: "Verification Strategies"},
			{Kind: types.Paragraph, Content: "Ask the AI to explain its own code: \"Walk me through this function step by step.\" If the explanation doesn't match what the code actually does, you've found a bug. Write tests before accepting generated code. Run it with edge-case inputs before committing."},
			{Kind: types.Code, Content: "  Verification workflow:\n  1. Read every line of the generated code\n  2. Check that imported methods actually exist\n  3. Run tests with edge-case inputs\n  4. Check error paths and nil handling\n  5. Only then: commit the code"},
			{Kind: types.Heading, Content: "Example: Spot the Bug"},
			{Kind: types.Code, Content: "  func isValidAge(age int) bool {\n      if age < 0 || age > 150 {\n          return true   // ← Bug! Returns true for INVALID ages\n      }\n      return false\n  }"},
			{Kind: types.Paragraph, Content: "This function returns true when age IS out of range — the exact opposite of what \"isValid\" implies. The code reads naturally and compiles without error, but the logic is inverted. AI frequently produces this kind of subtle inversion."},
			{Kind: types.Heading, Content: "Building Review Habits"},
			{Kind: types.Bullet, Content: "Treat AI output like a junior developer's pull request — review with the same rigor\nNever commit AI code you haven't read line by line\nRun the test suite after every AI-generated change\nIf you can't explain what a line does, don't keep it\nUse the AI itself as a second reviewer: \"Find bugs in this code\""},
			{Kind: types.Callout, Content: "The most dangerous AI bugs aren't the ones that crash — they're the ones that silently produce wrong results. A missing bounds check, an inverted condition, a race condition: all compile cleanly and may only surface in production."},
		},
		Questions: []types.QuizQuestion{
			{
				Kind:        types.MultipleChoice,
				Prompt:      "What makes AI-generated bugs harder to spot than typical bugs?",
				Choices:     []string{"AI bugs always cause compiler errors", "AI code is poorly formatted", "AI code looks polished and authoritative, masking subtle errors", "AI bugs only appear in production"},
				CorrectIdx:  2,
				Explanation: "AI produces clean, well-formatted code that passes the \"looks right\" test, making it easy to skim past subtle logical errors.",
			},
			{
				Kind:        types.Ordering,
				Prompt:      "Put these code review steps in the most effective order:",
				Choices:     []string{"Read every line of the generated code", "Check that imported methods actually exist", "Run tests with edge-case inputs", "Commit the code"},
				CorrectIdx:  0,
				Explanation: "Start with a thorough read, verify APIs exist, test edge cases, then commit only after verification passes.",
			},
			{
				Kind:        types.MultipleChoice,
				Prompt:      "Which AI code mistake is MOST dangerous because it may only surface in production?",
				Choices:     []string{"Syntax error", "Missing import statement", "Race condition in concurrent code", "Wrong variable name"},
				CorrectIdx:  2,
				Explanation: "Race conditions compile cleanly, pass most tests, and only manifest under concurrent load — making them extremely dangerous in production.",
			},
			{
				Kind:        types.FillBlank,
				Prompt:      "You should treat AI-generated code like a _______ developer's pull request.",
				Answer:      "junior",
				Explanation: "Reviewing AI output with the same rigor as a junior developer's PR ensures you catch subtle errors instead of trusting the polished appearance.",
			},
		},
	})
}
