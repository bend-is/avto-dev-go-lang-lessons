 #!/usr/bin/make

help: ## Show this help
	@printf "\033[33m%s:\033[0m\n" 'Available commands'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[32m%-11s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

run: ### Run application
	go run main.go

build: ### Build application
	go build .

lint: ### Lint application code
	golangci-lint run

test: ### Run tests
	go test ./...

bench: ### Run benchmarks
	go test -bench=. -benchmem ./...
