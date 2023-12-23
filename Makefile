GO := go
BUILD_DIR := build
APP_NAME := app
SRC := cmd/main.go

GO_FLAGS := -o $(BUILD_DIR)/$(APP_NAME) -ldflags "-s -w"

.PHONY: all clean example

all: build

build: $(SRC)
	@mkdir -p $(BUILD_DIR)
	$(GO) build $(GO_FLAGS) $(SRC)

clean:
	@rm -rf $(BUILD_DIR)

example: 
	@go run cmd/main.go  -db "root:examplepassword@tcp(localhost:3306)/exampledb" -ipfs "http://localhost:5001" -table users
