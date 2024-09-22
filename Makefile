build:
	@go build -o bin/monkey cmd/monkey-interpreter/main.go

run: build
	@./bin/monkey

test:
	@go test -v ./internal/...