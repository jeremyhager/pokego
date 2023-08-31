GO_CMD=go
GO_BUILD=$(GO_CMD) build
BIN_PATH=bin
RELEASE_PATH=dist
BINARY_NAME=pokego
GOARCH_x64=amd64
TAR_CMD=tar
TAR_GZ= $(TAR_CMD) -zcf

default: test build

release: build
	mkdir -p $(RELEASE_PATH)
	$(TAR_GZ) $(RELEASE_PATH)/linux.gz -C $(BIN_PATH)/ $(BINARY_NAME)
	$(TAR_GZ) $(RELEASE_PATH)/mac.gz -C $(BIN_PATH)/mac $(BINARY_NAME)
	$(TAR_GZ) $(RELEASE_PATH)/windows.gz -C $(BIN_PATH)/windows $(BINARY_NAME).exe

build: build-linux build-mac build-windows

build-linux: clean
	GOARCH=$(GOARCH_x64) GOOS=linux $(GO_BUILD) -o $(BIN_PATH)/$(BINARY_NAME)

build-mac: clean
	GOARCH=$(GOARCH_x64) GOOS=darwin $(GO_BUILD) -o $(BIN_PATH)/mac/$(BINARY_NAME)

build-windows: clean
	GOARCH=$(GOARCH_x64) GOOS=windows $(GO_BUILD) -o $(BIN_PATH)/windows/$(BINARY_NAME).exe

clean:
	rm -rf $(BIN_PATH)
	rm -rf $(RELEASE_PATH)

test:
	$(GO_CMD) test ./...

fmt:
	$(GO_CMD) fmt ./...