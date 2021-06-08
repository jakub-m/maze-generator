default: test fmt build run
build:
	go build -o maze cli/main.go
test:
	go test ./...
fmt:
	go fmt ./...
run:
	./maze -v -h 6 -w 6 -o tmp.svg
.phony: build
