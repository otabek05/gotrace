APP_NAME := gotrace
SRC := backend/cmd/server/main.go
BIN_DIR := bin
OUT := $(BIN_DIR)/$(APP_NAME)

.PHONY: all server run clean fmt

all: server

server:
	@echo ">> Building $(APP_NAME)..."
	@mkdir -p $(BIN_DIR)
	@go build -o $(OUT) $(SRC)
	@echo ">> Build complete: $(OUT)"

run: server
	@echo ">> Running $(APP_NAME)..."
	@./$(OUT)

clean:
	@echo ">> Cleaning build artifacts..."
	@rm -rf $(BIN_DIR)

fmt:
	@echo ">> Formatting Go code..."
	@go fmt ./...
