
CLI_SOURCE_FILES?=./cmd/plugin
CLI_BINARY_NAME=binary
CLI_DESTINATION=./bin/$(CLI_BINARY_NAME)

TEST_CMD?=go test
UNIT_TAGS?=unit
COVERAGE?=coverage.out

.PHONY: build
build: ## Generate the binary in ./bin
	@echo "==> Building $(CLI_BINARY_NAME) binary"
	go build -o $(CLI_DESTINATION) $(CLI_SOURCE_FILES)

.PHONY: unit-test
unit-test: ## Run unit-tests
	@echo "==> Running unit tests..."
	$(TEST_CMD) --tags="$(UNIT_TAGS)" -race -cover -coverprofile $(COVERAGE) -count=1 ./...

.PHONY: fuzz-normalizer-test
fuzz-normalizer-test: ## Run fuzz test
	@echo "==> Running fuzz test..."
	$(TEST_CMD) -fuzz=Fuzz -fuzztime 50s --tags="$(UNIT_TAGS)" -race ./internal/kubernetes/operator/resources

.PHONY: lint
lint: ## Run linter
	golangci-lint run

.PHONY: help
.DEFAULT_GOAL := help
help:
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'