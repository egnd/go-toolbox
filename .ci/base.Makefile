MAKEFLAGS += --always-make --no-print-directory

.PHONY: help

help:
	@grep -hE '[a-zA-Z_-]+:[a-zA-Z _-]+## ' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

%:
	@:

owner: ## Reset folder owner
	sudo chown --changes -R $$(whoami) ./
	@echo "Success"
