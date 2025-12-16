APP := gotrace
BACKEND_DIR := backend
BUILD_DIR := build

.PHONY: linux darwin windows frontend

frontend:
	cd frontend && npm run build

linux: frontend
	cd $(BACKEND_DIR) && \
	GOOS=linux GOARCH=amd64 go build -o ../$(BUILD_DIR)/linux/$(APP) ./cmd

darwin: frontend
	cd $(BACKEND_DIR) && \
	GOOS=darwin GOARCH=amd64 go build -o ../$(BUILD_DIR)/darwin/$(APP) ./cmd

windows: frontend
	cd $(BACKEND_DIR) && \
	GOOS=windows GOARCH=amd64 go build -o ../$(BUILD_DIR)/windows/$(APP).exe ./cmd
