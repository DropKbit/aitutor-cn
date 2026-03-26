# AITutor-ZH Design

**Goal:** Turn the existing AITutor project into a complete Chinese-localized version named `AITutor-ZH`, preserving the full 17-lesson curriculum, interactive flow, and upstream attribution.

## Scope

- Translate all lesson metadata and content in `internal/content/`.
- Translate all user-visible terminal UI strings in `internal/app/`, `internal/lesson/`, `internal/quiz/`, `internal/ui/`, and `pkg/types/`.
- Rewrite `README.md` into Chinese using the `AITutor-ZH` name.
- Preserve the upstream MIT license and add explicit attribution to the original author in the Chinese README.

## Non-Goals

- Do not change the Go module path or package import paths.
- Do not redesign visualizations or alter the lesson/state machine architecture.
- Do not change release automation, npm package names, or Homebrew formula names beyond documentation wording.

## Design Decisions

### Branding

- The localized product name shown to users becomes `AITutor-ZH`.
- Internal import paths remain `github.com/DropKbit/aitutor-cn` to avoid unnecessary refactors.

### Localization Strategy

- Translate content in place rather than maintaining parallel English and Chinese lesson trees.
- Keep technical terms precise; when an English term is standard in the ecosystem, retain it where helpful, such as `MCP`, `AGENTS.md`, `git worktree`, and `ToolSearch`.
- Preserve code samples and commands when translation would reduce accuracy.

### UI Coverage

- Translate difficulty tiers, lesson phases, help text, welcome screen, completion screen, quiz prompts, and sidebar/footer hints.
- Keep keyboard keys themselves unchanged while translating their descriptions.

### Content Coverage

- Every lesson retains its original structure: title, summary, theory blocks, quiz prompts, choices, and explanations.
- ASCII diagrams are translated where they are explanatory content rather than literal code/config.

### Attribution

- Leave `LICENSE` intact.
- Add a dedicated README section stating that `AITutor-ZH` is a Chinese-localized derivative of the original AITutor by Naor Peled, distributed under the original MIT license.

## Risks

- Chinese text is visually wider in terminal layouts, which may increase wrapping and sidebar truncation.
- Some ASCII diagrams may need careful wording to stay aligned.
- README branding may describe `AITutor-ZH` while install commands still target the upstream package unless the release pipeline is separately renamed.

## Validation

- Add regression tests for translated tier/phase labels and lesson registration/localization coverage.
- Run `go test ./...`, `go build ./...`, and `go vet ./...`.
