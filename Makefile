#!make

-include .ci/base.Makefile

CALL_PARAM=$(filter-out $@,$(MAKECMDGOALS))

_packages:
	@for package_dir in $$(go list -f '{{.Dir}}' -m | xargs); do \
		echo "\nPackage $$(basename $$package_dir):"; \
		cd $$package_dir && make $$cmd; \
	done

############################################################################################################################## 

check-conflicts: ## Check for git conflicts
	@if grep -rn '^<<<\<<<< ' .; then exit 1; fi
	@if grep -rn '^===\====$$' .; then exit 1; fi
	@if grep -rn '^>>>\>>>> ' .; then exit 1; fi
	@echo "All is OK"

check-todos: ## CHeck for TODO's
	@if grep -rn '@TO\DO:' .; then exit 1; fi
	@echo "All is OK"

check-master: ## Check for latest master in current branch
	@git remote update
	@if ! git log --pretty=format:'%H' | grep $$(git log --pretty=format:'%H' -n 1 origin/master) > /dev/null; then exit 1; fi
	@echo "All is OK"

mocks:
	@$(MAKE) _packages cmd=mocks

tests:
	@$(MAKE) _packages cmd=tests

coverage:
	@$(MAKE) _packages cmd=coverage DISABLE_HTML=$(DISABLE_HTML)

lint:
	@$(MAKE) _packages cmd=lint

vendor:
	@$(MAKE) _packages cmd=vendor


asdfg:
	@mkdir -p .profiles
	go test -race -cover -covermode=atomic -coverprofile=.profiles/cover.out.tmp metrics