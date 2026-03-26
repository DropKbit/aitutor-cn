# AITutor-ZH Implementation Plan

> **For agentic workers:** REQUIRED: Use superpowers:subagent-driven-development (if subagents available) or superpowers:executing-plans to implement this plan. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Fully localize AITutor into Chinese as `AITutor-ZH` while preserving all lessons, quizzes, terminal flows, and upstream attribution.

**Architecture:** Keep the existing lesson registry, visualization implementations, and module paths unchanged. Replace user-facing English strings in place, add regression tests for localization-sensitive behavior, and rewrite the README with explicit attribution.

**Tech Stack:** Go, Bubble Tea, Lip Gloss, Bubbles, Markdown docs

---

## Chunk 1: Regression Coverage

### Task 1: Add localization regression tests

**Files:**
- Create: `internal/lesson/localization_test.go`
- Test: `internal/lesson/localization_test.go`

- [ ] **Step 1: Write the failing test**

Create tests that assert:
- `types.Beginner.String()`, `types.Intermediate.String()`, and `types.Advanced.String()` return Chinese labels.
- `lesson.PhaseTheory.String()` and related phase labels return Chinese labels.
- All 17 lessons still register after blank-importing all content packages.
- Every lesson title and summary contain Chinese text.

- [ ] **Step 2: Run test to verify it fails**

Run: `go test ./internal/lesson`
Expected: FAIL because current labels and lesson metadata are still English.

- [ ] **Step 3: Keep tests as regression guard**

Do not weaken the assertions once production changes begin.

## Chunk 2: Core UI Localization

### Task 2: Translate shared labels and lesson lifecycle text

**Files:**
- Modify: `pkg/types/types.go`
- Modify: `internal/lesson/model.go`
- Modify: `internal/quiz/model.go`
- Modify: `internal/quiz/feedback.go`
- Modify: `internal/quiz/fill_blank.go`
- Modify: `internal/quiz/ordering.go`
- Modify: `internal/ui/footer.go`
- Modify: `internal/app/keys.go`
- Modify: `internal/app/app.go`

- [ ] **Step 1: Translate tier names and phase names**
- [ ] **Step 2: Translate quiz status text, feedback, placeholders, and instructions**
- [ ] **Step 3: Translate welcome/help/completion screens and footer bindings**
- [ ] **Step 4: Run `go test ./internal/lesson`**
- [ ] **Step 5: Run `gofmt -w` on modified Go files**

## Chunk 3: Curriculum Localization

### Task 3: Translate beginner lessons

**Files:**
- Modify: `internal/content/beginner/01_what_is_ai.go`
- Modify: `internal/content/beginner/02_context_window.go`
- Modify: `internal/content/beginner/03_tools.go`
- Modify: `internal/content/beginner/04_prompt_engineering.go`

- [ ] **Step 1: Translate titles, summaries, theory, quizzes, and explanations**
- [ ] **Step 2: Preserve code/config examples that should remain literal**
- [ ] **Step 3: Run `go test ./internal/lesson`**

### Task 4: Translate intermediate lessons

**Files:**
- Modify: `internal/content/intermediate/05_agents_md.go`
- Modify: `internal/content/intermediate/06_execution_modes.go`
- Modify: `internal/content/intermediate/07_hooks.go`
- Modify: `internal/content/intermediate/08_memory.go`
- Modify: `internal/content/intermediate/09_agentic_loop.go`
- Modify: `internal/content/intermediate/10_codegen_prompts.go`
- Modify: `internal/content/intermediate/11_code_review.go`

- [ ] **Step 1: Translate all metadata and lesson bodies**
- [ ] **Step 2: Keep research references and command literals accurate**
- [ ] **Step 3: Run `go test ./internal/lesson`**

### Task 5: Translate advanced lessons

**Files:**
- Modify: `internal/content/advanced/12_mcp.go`
- Modify: `internal/content/advanced/13_skills.go`
- Modify: `internal/content/advanced/14_subagents.go`
- Modify: `internal/content/advanced/15_git_worktrees.go`
- Modify: `internal/content/advanced/16_tool_search.go`
- Modify: `internal/content/advanced/17_batch_tools.go`

- [ ] **Step 1: Translate all metadata and lesson bodies**
- [ ] **Step 2: Preserve protocol names, commands, and tool identifiers**
- [ ] **Step 3: Run `go test ./internal/lesson`**

## Chunk 4: Documentation Localization

### Task 6: Rewrite README and attribution

**Files:**
- Modify: `README.md`

- [ ] **Step 1: Rewrite the README in Chinese using the `AITutor-ZH` name**
- [ ] **Step 2: Add curriculum overview, usage, and project structure in Chinese**
- [ ] **Step 3: Add explicit attribution and copyright notice for Naor Peled**

## Chunk 5: Final Verification

### Task 7: Full project validation

**Files:**
- Modify: `internal/lesson/localization_test.go` if test fixtures need tightening after implementation

- [ ] **Step 1: Run `gofmt -w` on all modified Go files**
- [ ] **Step 2: Run `go test ./...`**
- [ ] **Step 3: Run `go build ./...`**
- [ ] **Step 4: Run `go vet ./...`**
- [ ] **Step 5: Review `git diff --stat` for unintended changes**
