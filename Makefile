gofiles = $(shell find . -type f -name \*.go)
bin = maze

default: test fmt run

build: $(bin)

$(bin): $(gofiles)
	go build -o $(bin) cli/main.go

test: $(gofiles)
	go test ./...

fmt: $(gofiles)
	go fmt ./...

run: $(bin)
	./$(bin) -r 0 -v -h 8 -w 8 -s 20 -o tmp.svg

generate-examples: $(bin)
	mkdir -p examples
	./$(bin) -r 0 -h 10 -w 10 -o examples/1010.svg
	./$(bin) -r 0 -h 20 -w 20 -s 25 -o examples/2020.svg
	./$(bin) -r 0 -h 40 -w 80 -s 10 -o examples/4080.svg

.phony:
