# AGENTS.md

## Project Overview

notify-mail is a Go CLI tool that sends email notifications via Gmail's SMTP server. It supports plain message bodies and HTML templates with variable substitution.

## Build, Lint, and Test Commands

### Build
```bash
make                # Build the notify-mail binary (sets GOOS=linux GOARCH based on host)
go build            # Alternative: build without Makefile
```

### Test
```bash
make test           # Run all tests with race detector and coverage
go test ./... -race -cover              # Equivalent to make test
go test ./mail -run TestTemplateName -race -cover   # Run a single test in the mail package
go test . -run TestFunctionName -race -cover         # Run a single test in the root package
```

### Lint / Format
```bash
go vet ./...        # Run Go vet (static analysis)
gofmt -d .          # Check formatting (dry run)
gofmt -w .          # Format all files in place
```

### Other
```bash
make tidy           # Run go mod tidy
make vendor         # Tidy and vendor dependencies
make clean          # Remove the binary
make install        # Install binary to DESTDIR/usr/bin/
```

## Project Structure

```
main.go              # Entry point: CLI argument parsing, notification dispatch
mail/
  notify.go          # Mail struct, NewNotification constructor, Send, SendTemplate, internal send
  template.go        # Template struct, NewTemplate, ReplaceContent
debian/              # Debian packaging files for dpkg-buildpackage
Makefile             # Build, test, clean, install targets
```

## Code Style Guidelines

### Imports

Group imports with blank lines separating stdlib, then third-party, then local packages. Within each group, sort alphabetically:

```go
import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/adelolmo/notify-mail/mail"
)
```

Use explicit imports — no dot imports or blank imports unless required.

### Formatting

- Use `gofmt` formatting: tabs for indentation, standard Go spacing.
- Use `gofmt -w .` before committing changes.
- No trailing whitespace; files end with a newline.

### Types and Structs

- Define struct types for domain concepts (e.g., `Mail`, `Template`).
- Exported fields use PascalCase; unexported fields use camelCase.
- Use concrete types rather than interfaces unless polymorphism is needed.
- Factory functions follow the `New*` naming pattern and return a pointer.

```go
type Mail struct {
	Sender         string
	Authentication smtp.Auth
}

func NewNotification() (*Mail, error) { ... }
```

### Naming Conventions

- Packages: lowercase, single-word names (e.g., `mail`).
- Exported functions/types: PascalCase (`NewNotification`, `SendTemplate`).
- Unexported functions: camelCase (`send`).
- Acronyms remain uppercase: `SMTP`, `URL`, `ID`.
- Constructor pattern: `New` prefix returning `(*Type, error)`.

### Error Handling

- Return `error` as the last return value from functions.
- Wrap errors with context using `fmt.Errorf`:

```go
return fmt.Errorf("cannot replace variables for placeholders in template file %s. Error: %s",
	templateFilename, err)
```

- Use `log.Fatal(err)` only in `main()` or top-level CLI handling — never in library code.
- Do not panic in library packages; prefer error returns.

### Environment Variables

- Configuration is via environment variables: `NOTIFY_MAIL_ACCOUNT`, `NOTIFY_MAIL_PASSWORD`.
- Validate required env vars in constructors; return errors if missing.

```go
v := os.Getenv(key)
if v == "" {
	return nil, fmt.Errorf("environment variable %q is required", key)
}
```

### Testing Conventions

- Test files go alongside source files: `notify_test.go` next to `notify.go`.
- Use the standard `testing` package.
- Test function naming: `TestFunctionName` for unit tests.
- Run tests with race detector: `go test ./... -race -cover`.
- Mock external dependencies (SMTP) in tests rather than hitting real servers.

### Package Layout

- The `mail` package contains all library logic; `main.go` is a thin CLI wrapper.
- Keep business logic out of `main.go`; it should only parse args and call `mail` package functions.
- Unexported helper functions (e.g., `send`) are package-private implementation details.

### General Go Conventions

- Follow Effective Go guidelines: https://go.dev/doc/effective_go
- Prefer composition over inheritance.
- Use `strings.Split`, `strings.Replace`, `fmt.Sprintf` from the stdlib — avoid unnecessary third-party dependencies.
- Keep the module path consistent: `github.com/adelolmo/notify-mail`.
- Target Go 1.24 as specified in go.mod.