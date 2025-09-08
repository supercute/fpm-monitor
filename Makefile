.PHONY: build clean test

# Binary name
BINARY_NAME=fpm-monitor
# Binary directory
BINARY_DIR=bin
CMD_DIR=./cmd/fpm-monitor

# Binary releases
BINARY_LINUX=fpm-monitor_linux_amd64
BINARY_MACOS=fpm-monitor_macos_arm64

build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BINARY_DIR)
	@go build -o $(BINARY_DIR)/$(BINARY_NAME) $(CMD_DIR)
release:
	@echo "Build release..."
	@GOOS=linux GOARCH=amd64 go build -o $(BINARY_DIR)/$(BINARY_LINUX) $(CMD_DIR)
	@GOOS=darwin GOARCH=arm64 go build -o $(BINARY_DIR)/$(BINARY_MACOS) $(CMD_DIR)
clean:
	@echo "Cleaning..."
	@rm -rf $(BINARY_DIR)

test:
	@go test -v ./...