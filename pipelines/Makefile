#!make

MAKEFLAGS += --always-make --no-print-directory

.PHONY: help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

%:
	@:

########################################################################################################################

benchmarks: ## Run benchmarks
	go test -benchmem -bench . benchmarks_test.go

ocker-benchmarks:
	docker run --rm -it -v $$(pwd):/src -w /src --entrypoint make golang:1.19 benchmarks

profiles: ## Generate svg-profiles
	@mkdir -p .profiles
	go test -cpuprofile .profiles/cpu.out -memprofile .profiles/mem.out -bench . benchmarks_test.go
	go tool pprof -svg pipeline.test .profiles/cpu.out > .profiles/cpu.svg
	go tool pprof -svg -alloc_space pipeline.test .profiles/mem.out > .profiles/mem.svg
	go tool pprof -svg -alloc_objects pipeline.test .profiles/mem.out > .profiles/obj.svg

docker-profiles:
	docker run --rm -it -v $$(pwd):/src -w /src --entrypoint bash golang:1.19 -c "apt-get update -qq && apt-get install -y graphviz > /dev/null && make profiles $(CALL_PARAM)"
	@echo "Read reports at:"
	@echo "- file://$$(pwd)/.profiles/cpu.svg"
	@echo "- file://$$(pwd)/.profiles/mem.svg"
	@echo "- file://$$(pwd)/.profiles/obj.svg"
