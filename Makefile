GO_CMD=go
GO_BUILD=$(GO_CMD) build
BIN_PATH=bin
BINARY_NAME=pokego
GOARCH_x64=amd64

default: test build

build: build-linux build-mac build-windows

build-linux: clean
	GOARCH=$(GOARCH_x64) GOOS=linux $(GO_BUILD) -o $(BIN_PATH)/$(BINARY_NAME)

build-mac: clean
	GOARCH=$(GOARCH_x64) GOOS=darwin $(GO_BUILD) -o $(BIN_PATH)/mac/$(BINARY_NAME)

build-windows: clean
	GOARCH=$(GOARCH_x64) GOOS=windows $(GO_BUILD) -o $(BIN_PATH)/windows/$(BINARY_NAME).exe

clean:
	rm -rf $(BIN_PATH)

test:
	$(GO_CMD) test ./...

fmt:
	$(GO_CMD) fmt ./...