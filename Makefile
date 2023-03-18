#!make

-include .ci/base.Makefile

CALL_PARAM=$(filter-out $@,$(MAKECMDGOALS))
COVER_REPORT=html

_packages:
	@for package_dir in $$(go list -m -f '{{.Dir}}' | xargs); do \
		echo "Package $$(basename $$package_dir):"; \
		cd $$package_dir && make $$cmd; \
		echo ""; \
	done

################################################################################################ 

check-conflicts: ## Check for git conflicts
	@if grep -rn '^<<<\<<<< ' .; then exit 1; fi
	@if grep -rn '^===\====$$' .; then exit 1; fi
	@if grep -rn '^>>>\>>>> ' .; then exit 1; fi
	@echo "All is OK"

check-todos: ## Check for TODO's
	@if grep -rn '@TO\DO:' .; then exit 1; fi
	@echo "All is OK"

check-master: ## Check for latest master in current branch
	@git remote update
	@if ! git log --pretty=format:'%H' | grep $$(git log --pretty=format:'%H' -n 1 origin/master) > /dev/null; then exit 1; fi
	@echo "All is OK"

mocks: ## Generate mocks
	@$(MAKE) _packages cmd=mocks

test: ## Run tests
	@$(MAKE) _packages cmd=test

coverage: ## Show code coverage
	@mkdir -p .profiles && echo "mode: atomic" > .profiles/cover.out
	@$(MAKE) _packages cmd=coverage COVER_REPORT=false
	go tool cover -func=.profiles/cover.out
ifeq ($(COVER_REPORT),html)
	go tool cover -html=.profiles/cover.out -o .profiles/report.html
endif

lint: ## Lint all packages
	go list -m -f '{{.Dir}}' | xargs golangci-lint run --color=always --config=.golangci.yml

vendor: ## Install packages required modules
	@$(MAKE) _packages cmd=vendor
