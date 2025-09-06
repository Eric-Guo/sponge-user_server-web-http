# Repository Guidelines

## Project Structure & Module Organization
This Go monolith exposes HTTP endpoints from `cmd/user_server/main.go` and routes traffic through `internal/server`, `internal/routers`, and `internal/handler`. Domain models and persistence live under `internal/model` and `internal/dao`, while reusable DTOs sit in `internal/types`. Shared assets such as Swagger specs are tracked in `docs/`, and deployment assets (Docker, K8s, binary packaging) stay under `deployments/`. Configuration templates live in `configs/`, with `configs/user_server.yml` as the local default.

## Build, Test & Development Commands
- `make run` – rebuilds the binary and launches the server; pass `Config=configs/user_server.yml` when testing alternate settings.
- `make build` – produces `cmd/user_server/user_server` for use in CI packaging.
- `make test` – runs `go test -count=1 -short` on all non-generated packages.
- `make cover` – emits `cover.out` and opens an HTML report.
- `make ci-lint` – applies `gofmt -s` and executes `golangci-lint` per `.golangci.yml`.
- `make docs` – regenerates Swagger definitions via `scripts/swag-docs.sh`.

## Coding Style & Naming Conventions
Go 1.25.1 is the target runtime; always run `make ci-lint` before review. Source files stay `gofmt`-ed (tabs for indentation) and `goimports` enforces import ordering with `test-user-server` as the local prefix. Use PascalCase for exported handlers (e.g., `UserRegister`) and lowerCamelCase for locals. Keep functions focused; `gocyclo` warns once complexity exceeds 20 and long lines are tolerated up to 200 characters. Avoid `github.com/pkg/errors`; prefer the standard library `errors` package.

## Testing Guidelines
Place unit tests alongside the implementation as `*_test.go` files and name cases `Test<Thing>` to keep `golangci-lint` happy. Leverage `testify/require` for assertions and `go-sqlmock` when persisting through `internal/dao`. Run `make test` before committing and `make cover` when touching database-critical logic; keep coverage stable and upload `cover.out` artifacts when PRs modify persistence paths.

## Commit & Pull Request Guidelines
Commit history favors short, imperative subjects (e.g., `Upgrade to sponge v1.15.0`, `Add Rails Auth`). Group related changes and avoid mixed-purpose commits. For every PR, include: a concise summary, links to tracking issues, validation notes (`make test`, `make ci-lint`), and updated Swagger screenshots or URLs when APIs change. Tag reviewers who own the touched package and mention deployment impacts if scripts under `deployments/` shift.

## Configuration & Ops Notes
`.env` files are not used; pass secrets via runtime environment variables instead of editing `configs/user_server.yml`. The overridden `github.com/go-dev-frame/sponge` module resolves to `/Users/guochunzhong/git/oss/sponge`; verify that path is reachable in your environment before running generators or builds.
