.PHONY: build clean test install-deps release

# 变量
PROJECT_NAME := system-monitor
VERSION ?= dev
BUILD_DIR := build

# 默认目标
all: build

# 构建
build:
	@echo "Building $(PROJECT_NAME)..."
	go build -o $(BUILD_DIR)/$(PROJECT_NAME) ./cmd/system-monitor/

# 清理
clean:
	@echo "Cleaning build files..."
	rm -rf $(BUILD_DIR)
	rm -rf release

# 测试
test:
	@echo "Running tests..."
	go test -v ./...

# 安装依赖
install-deps:
	@echo "Installing dependencies..."
	./scripts/install-deps.sh

# 发布构建
release:
	@echo "Creating release build..."
	./scripts/release.sh $(VERSION)

# 运行
run:
	go run ./cmd/system-monitor/

# 安装到系统
install: build
	@echo "Installing to /usr/local/bin..."
	sudo cp $(BUILD_DIR)/$(PROJECT_NAME) /usr/local/bin/