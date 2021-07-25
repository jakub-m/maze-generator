default: test fmt build run
build:
	go build -o maze cli/main.go
test:
	go test ./...
fmt:
	go fmt ./...
run:
	./maze -v -h 6 -w 6 -o tmp.svg

generate-examples:
	mkdir -p examples
	./maze -r 0 -h 10 -w 10 -o examples/1010.svg
	./maze -r 0 -h 20 -w 20 -s 25 -o examples/2020.svg
	./maze -r 0 -h 40 -w 80 -s 10 -o examples/4080.svg

.phony: build generate-examples
