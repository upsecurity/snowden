build:
	@go build -o bin/snowden

run: build
	@./bin/snowden

test:
	@go test -v ./...