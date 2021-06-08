default: test fmt build
build:
	go build -o maze cli/main.go
test:
	go test ./...
fmt:
	go fmt ./...
.phony: build
