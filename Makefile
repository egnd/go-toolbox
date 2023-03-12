#!make

MAKEFLAGS += --always-make --no-print-directory

.PHONY: help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

%:
	@:

########################################################################################################################

owner: ## Reset folder owner
	sudo chown --changes -R $$(whoami) ./
	@echo "Success"

check-conflicts: ## Find git conflicts
	@if grep -rn '^<<<\<<<< ' .; then exit 1; fi
	@if grep -rn '^===\====$$' .; then exit 1; fi
	@if grep -rn '^>>>\>>>> ' .; then exit 1; fi
	@echo "All is OK"

check-todos: ## Find TODO's
	@if grep -rn '@TO\DO:' .; then exit 1; fi
	@echo "All is OK"

check-master: ## Check for latest master in current branch
	@git remote update
	@if ! git log --pretty=format:'%H' | grep $$(git log --pretty=format:'%H' -n 1 origin/master) > /dev/null; then exit 1; fi
	@echo "All is OK"

mocks: ## Generate mocks
# go generate ./...
	rm -rf pipelines/mocks && mockery --config=.mockery.yaml --name=. --dir=pipelines --output=pipelines/mocks
	rm -rf xmlparse/mocks && mockery --config=.mockery.yaml --name=. --dir=xmlparse --output=xmlparse/mocks
	rm -rf tg/tgchain/mocks && mockery --config=.mockery.yaml --name=. --dir=tg/tgchain --output=tg/tgchain/mocks
	rm -rf tg/telebotmdw/mocks && mockery --config=.mockery.yaml --name=. --dir=tg/telebotmdw --output=tg/telebotmdw/mocks
	rm -rf metrics/mocks && mockery --config=.mockery.yaml --name=. --dir=metrics --output=metrics/mocks

tests: ## Run unit tests
	@mkdir -p .profiles
	go test -race -cover -covermode=atomic -coverprofile=.profiles/cover.out.tmp ./...

coverage: tests ## Check code coveragem
	@cat .profiles/cover.out.tmp | grep -v "mock_" > .profiles/cover.out
	go tool cover -func=.profiles/cover.out
ifneq ($(DISABLE_HTML),true)
	go tool cover -html=.profiles/cover.out -o .profiles/report.html
endif

lint: ## Lint source code
	@clear
	golangci-lint run --color=always --config=.golangci.yml ./...
