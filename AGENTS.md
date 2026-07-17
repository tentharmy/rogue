# Repository Guidelines

## Project Structure & Module Organization

This is a small Go terminal game module, `github.com/tentharmy/rogue`.

- `main.go` is the executable entry point and wires the app service into the TUI.
- `inner/domain/` contains core game entities and rules, such as `Game` and `Actor`.
- `inner/app/` contains the application service layer and snapshots exposed to the UI.
- `inner/geom/` contains shared geometry types like `Vec2`.
- `inner/tui/` contains Bubble Tea and Lip Gloss rendering and input handling.
- `.github/workflows/ci.yml` defines the required verification pipeline.

Keep domain behavior independent from terminal UI concerns. UI code should depend on app snapshots and commands rather than directly mutating domain state.

## Build, Test, and Development Commands

- `go run .` starts the terminal UI locally.
- `go test ./...` runs all Go tests in the module.
- `go vet ./...` runs standard Go static analysis.
- `gofmt -w .` formats all Go files before committing.
- `golangci-lint run ./...` runs the same lint family used by CI, if installed locally.
- `go build -o build/gorogue .` builds the binary produced by CI.

CI runs formatting checks, `go vet`, `golangci-lint`, tests, and a Linux binary build on pushes and pull requests to `main`.

## Coding Style & Naming Conventions

Use standard Go formatting and idioms. Package names are short lowercase nouns (`app`, `domain`, `geom`, `tui`). Export names only when they are part of another package's API, such as `GameService`, `GameSnapshot`, and `Vec2`.

Keep functions small and place behavior near the type it modifies. Prefer explicit domain methods, for example `Actor.Move`, over having UI code edit fields directly.

## Testing Guidelines

There are no tests yet, but new behavior should include focused Go tests using the standard `testing` package. Place tests beside the code under test with `_test.go` suffixes, for example `inner/domain/actor_test.go`. Name tests by behavior, such as `TestActorMoveUpdatesPosition`.

Run `go test ./...` before opening a pull request. For UI work, prefer testing service/domain behavior directly and keep Bubble Tea rendering tests narrow.

## Commit & Pull Request Guidelines

Recent history uses short imperative commit subjects, sometimes with a scope prefix, such as `add lipgloss and rendering` or `ci: run tests`. Keep subjects concise and describe the change, not the process.

Pull requests should include a brief summary, commands run locally, linked issues when applicable, and screenshots or terminal recordings for visible TUI changes. Ensure CI passes before requesting review.
