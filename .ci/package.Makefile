-include ../.ci/base.Makefile

mocks: ## Generate package mocks
# go generate ./...
	rm -rf mocks && mockery --config=../.mockery.yaml --name=. --dir=. --output=mocks

tests: ## Test package
	@mkdir -p .profiles
	go test -race -cover -covermode=atomic -coverprofile=.profiles/cover.out.tmp ./...

coverage: tests ## Check package code coverage
	@cat .profiles/cover.out.tmp | grep -v "mock_" > .profiles/cover.out
	go tool cover -func=.profiles/cover.out
ifneq ($(DISABLE_HTML),true)
	go tool cover -html=.profiles/cover.out -o .profiles/report.html
endif

lint: ## Lint package
	golangci-lint run --color=always --config=../.golangci.yml ./...
