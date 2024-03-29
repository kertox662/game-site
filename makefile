BINARY_NAME=game-site

build: generate-proto
	@echo "Building..."
	@go build -o bin/$(BINARY_NAME) -v

run: build
	@echo "Running..."
	@./bin/$(BINARY_NAME) api game

test: generate-proto
	@echo "Testing..."
	@go test -v ./...

bin:
	@mkdir -p bin

.PHONY: generate-proto
generate-proto:
	buf generate
