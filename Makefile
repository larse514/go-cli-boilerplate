GOFILES = $(shell find . -name '*.go' -not -path './vendor/*')
BINARY_NAME = hello-cli
default: clean dependencies test build

build: easyasync-cli

easyasync-cli: $(GOFILES)
	go build -o easyasync-cli -o $(BINARY_NAME)  . 

dependencies: 
	@go get gopkg.in/urfave/cli.v1
test: test-all

clean: 
	rm -f $(BINARY_NAME)

test-all:
	@go test -v ./...

test-min:
	@go test ./...

release:
	echo "not implemented"