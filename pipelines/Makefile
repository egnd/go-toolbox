#!make

-include ../.ci/package.Makefile

benchmarks: ## Run benchmarks
	go test -benchmem -bench . benchmarks_test.go

profiles: ## Generate svg-profiles
	@mkdir -p .profiles
	go test -cpuprofile .profiles/cpu.out -memprofile .profiles/mem.out -bench . benchmarks_test.go
	go tool pprof -svg pipeline.test .profiles/cpu.out > .profiles/cpu.svg
	go tool pprof -svg -alloc_space pipeline.test .profiles/mem.out > .profiles/mem.svg
	go tool pprof -svg -alloc_objects pipeline.test .profiles/mem.out > .profiles/obj.svg
