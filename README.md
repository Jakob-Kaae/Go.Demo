# Go.Demo

This repository is a small Go demo bootstrapped for development inside VS Code on WSL.

Quick start (WSL + VS Code):

1. Open the repository in VS Code using the Remote - WSL extension.
2. Make sure Go is installed in your WSL distribution (e.g. `sudo apt install golang` or follow golang.org/install).
3. Initialize the module and run tests:

```bash
# from the repository root
go mod init github.com/Jakob-Kaae/Go.Demo
go test ./...
go build -o bin/app ./cmd/app
./bin/app
```

What I added:

- `cmd/app/main.go` — small CLI that prints a greeting using the `greet` package.
- `pkg/greet/greet.go` and `pkg/greet/greet_test.go` — a tiny package with a unit test.
- `.vscode/settings.json` and `.vscode/extensions.json` — recommended VS Code settings for Go in WSL.
- `.gitignore` — common ignores for Go projects.

Next steps (optional): add CI (GitHub Actions), add golangci-lint, and enable `gopls` features in settings.