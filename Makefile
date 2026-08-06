# Run from Git Bash on Windows — recipes are POSIX sh.
BINARY := arowana

.DEFAULT_GOAL := help

.PHONY: help
help: ## Show available targets
	@grep -E '^[a-zA-Z_-]+:.*?## ' $(MAKEFILE_LIST) | \
		awk 'BEGIN{FS=":.*?## "}{printf "  %-12s %s\n", $$1, $$2}'

.PHONY: dev-api
dev-api: ## Run the Go API server on :8080
	go run ./cmd/arowana explore .

.PHONY: dev-web
dev-web: ## Run Vite on :5173 (proxies /api -> :8080)
	cd web && pnpm dev

.PHONY: types
types: ## Regenerate TypeScript types from Go
	go tool tygo generate

.PHONY: test
test: ## Run Go + frontend tests
	go test ./...
	cd web && pnpm vitest run

.PHONY: lint
lint: ## Lint Go + frontend
	go vet ./...
	cd web && pnpm lint

.PHONY: fmt
fmt: ## Format everything
	go fmt ./...
	cd web && pnpm exec prettier --write src

.PHONY: build
build: ## Build frontend, then the binary
	cd web && pnpm build
	go build -o $(BINARY) ./cmd/arowana

.PHONY: clean
clean: ## Remove build output
	rm -rf $(BINARY) $(BINARY).exe web/dist