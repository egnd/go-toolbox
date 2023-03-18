-include ../.ci/base.Makefile

COVER_REPORT=html

mocks: ## Generate package mocks
	rm -rf mocks && mockery

test: ## Test package
	@mkdir -p .profiles
	go test -race -cover -covermode=atomic -coverprofile=.profiles/cover.out.tmp ./...
	@cat .profiles/cover.out.tmp | grep -v "mock_" > .profiles/cover.out

coverage: test ## Check package code coverage
	@mkdir -p ../.profiles && cat .profiles/cover.out  | tail -n +2 >> ../.profiles/cover.out
ifeq ($(COVER_REPORT),html)
	go tool cover -html=.profiles/cover.out -o .profiles/report.html
else ifeq ($(COVER_REPORT),cli)
	go tool cover -func=.profiles/cover.out
endif

lint: ## Lint package
	golangci-lint run --color=always --config=../.golangci.yml ./...

vendor: ## Install required modules
	go mod tidy
