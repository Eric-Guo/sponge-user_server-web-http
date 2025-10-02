# Repository Guidelines

## Project Structure & Module Organization
- `cmd/user_server/main.go` boots the HTTP server; traffic flows through `internal/server`, `internal/routers`, and `internal/handler`.
- Domain logic and persistence sit in `internal/model` and `internal/dao`; shared DTOs live in `internal/types`.
- Swagger specs reside in `docs/`, deployment assets in `deployments/`, and configuration templates in `configs/` (local default: `configs/user_server.yml`).
- Tests live alongside source files as `*_test.go`; keep fixtures near their consuming packages.

## Build, Test & Development Commands
- `make run` rebuilds and starts the service; pass `Config=configs/user_server.yml` to exercise alternate configs.
- `make build` emits the deployable binary at `cmd/user_server/user_server` for packaging.
- `make test` runs `go test -count=1 -short ./...` to cover all non-generated packages.
- `make cover` produces `cover.out` and an HTML report for database-heavy updates.
- `make ci-lint` enforces `gofmt -s`, `goimports`, and `golangci-lint` per `.golangci.yml`.
- `make docs` regenerates Swagger definitions via `scripts/swag-docs.sh`.

## Coding Style & Naming Conventions
- Target Go 1.25.1; keep source formatted with `gofmt` (tabs) and `goimports` using `test-user-server` as the local prefix.
- Exported handlers use PascalCase (`UserRegister`), while locals and private helpers stay lowerCamelCase.
- Break work into focused functions to keep `gocyclo` under 20 and prefer the standard `errors` package over third-party wrappers.

## Testing Guidelines
- Name unit tests `Test<Subject>`; colocate mocks and fixtures with the code under test.
- Use `testify/require` for assertions and `go-sqlmock` when exercising DAO paths.
- Run `make test` before committing; when touching persistence or critical flows, add `make cover` artifacts to the PR.

## Commit & Pull Request Guidelines
- Write short, imperative commit subjects (e.g., `Add audit logging`, `Fix user lookup`).
- Keep commits single-purpose; avoid mixing refactors with feature work.
- Every PR should include a brief summary, linked issues, validation notes (`make test`, `make ci-lint`), and refreshed Swagger screenshots or URLs for API changes.
- Tag owners of the affected packages and call out deployment implications if files in `deployments/` move.

## Security & Configuration Tips
- Do not add `.env` files; inject secrets via runtime environment variables instead of editing `configs/user_server.yml`.
- Ensure the local override for `github.com/go-dev-frame/sponge` points to `/Users/guochunzhong/git/oss/sponge` before running generators.
