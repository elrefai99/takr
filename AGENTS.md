# Repository Guidelines

## Project Structure & Module Organization
The repository has two active codebases:

- Root Go CLI: [main.go](/Users/mohamedmostafa/Documents/elrefai/takr/main.go), [cmd/](/Users/mohamedmostafa/Documents/elrefai/takr/cmd), and [pkg/data/](/Users/mohamedmostafa/Documents/elrefai/takr/pkg/data). `cmd/` holds CLI actions, and `pkg/data/` handles filesystem bootstrapping.
- TypeScript API: [api/src/](/Users/mohamedmostafa/Documents/elrefai/takr/api/src) with `core/` for environment and MongoDB setup, and `utils/` for shared helpers such as logging, pagination, and request limits.

Generated runtime data lives under `json/` and is gitignored. API logs are written to `api/logs/`.

## Build, Test, and Development Commands
Run Go commands from the repository root:

- `go run .` starts the CLI.
- `go build ./...` verifies all Go packages compile.
- `go test ./...` runs Go tests when present.

Run API commands from `api/`:

- `pnpm dev` starts the Express app and message-queue worker in development.
- `pnpm build` compiles TypeScript to `api/dist/`.
- `pnpm test` runs Jest tests.
- `pnpm test:e2e` runs Vitest end-to-end checks.
- `pnpm lint` runs ESLint.

## Coding Style & Naming Conventions
For Go, keep code `gofmt`-clean, use tabs, and prefer short package names. Exported identifiers use PascalCase; internal helpers use camelCase.

For TypeScript, follow the existing `api/src/` style: single quotes, semicolon-light formatting, and clear utility names such as `limit-request.ts` and `hashText.ts`. Keep path aliases consistent with `@/` and `@src/` from [api/tsconfig.json](/Users/mohamedmostafa/Documents/elrefai/takr/api/tsconfig.json).

## Testing Guidelines
No committed tests are present yet, but new work should add them alongside the changed module. Use Go’s standard `*_test.go` pattern and TypeScript names like `*.test.ts` or `*.spec.ts`. Cover CLI flows, JSON file handling, and API startup or utility behavior before merging.

## Commit & Pull Request Guidelines
Git history follows Conventional Commits, for example `feat: initialize Express app with middleware and logging` and `chore: add TypeScript configuration file for API module`. Keep commits focused and use prefixes such as `feat:`, `fix:`, `chore:`, or `test:`.

Pull requests should include a short summary, affected areas (`cmd/`, `pkg/`, `api/src/`), linked issues if any, and the commands you ran to verify the change. Include screenshots only when output or logs materially change.

## Permissions

Global rule:

- Ask the user first before making any code change.
- Show the intended change for review when possible.
- Wait for the user to accept or reject the change before editing project code.

### Allow

- Read tracked project files needed for the task.
- Read source code under `src/`, configuration under `config/`, and docs such as `README.md`, `CLAUDE.md`, and this file.
- Create new source or documentation files when they are required for the requested change.
- Edit application code, route files, models, middleware, utilities, tests, and markdown documentation.
- Update `package.json` when the task explicitly requires script or dependency changes.
- Run safe repo-local commands such as `rg`, `ls`, `sed`, `git status`, `pnpm lint`, and `pnpm test`.

### Deny

- Do not read, print, or copy secrets from `.env`, `.env.dev`, or any credential file.
- Do not modify `.env`, `.env.dev`, or other secret-bearing files unless the user explicitly asks.
- Do not modify `node_modules/`, generated caches, or log files.
- Do not change `pnpm-lock.yaml` unless dependency work is part of the task.
- Do not delete files, rename major directories, or rewrite large parts of the codebase without explicit approval.
- Do not run destructive git or shell commands such as `git reset --hard`, `git checkout --`, or broad `rm` operations.
- Do not alter deployment/infrastructure files (`Dockerfile`, `docker-compose.yml`, `ecosystem.config.cjs`) unless the task explicitly requires it.
