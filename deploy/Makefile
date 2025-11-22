.PHONY: build run test clean swagger docker-build docker-run

# 应用名称
APP_NAME=pan-search-api

# 构建应用
build:
	@echo "Building $(APP_NAME)..."
	go build -o $(APP_NAME) main.go

# 运行应用
run:
	@echo "Running $(APP_NAME)..."
	go run main.go

# 生成Swagger文档
swagger:
	@echo "Generating Swagger documentation..."
	swag init -g main.go -o docs

# 清理构建文件
clean:
	@echo "Cleaning up..."
	rm -f $(APP_NAME)
	rm -rf docs

# 运行测试
test:
	@echo "Running tests..."
	go test ./...

# Docker构建
docker-build:
	@echo "Building Docker image..."
	docker build -t $(APP_NAME) .

# Docker运行
docker-run:
	@echo "Running Docker container..."
	docker run -p 8080:8080 $(APP_NAME)

# 安装开发依赖
deps:
	@echo "Installing development dependencies..."
	go install github.com/swaggo/swag/cmd/swag@latest

# 开发模式（自动重载）
dev:
	@echo "Starting development server with auto-reload..."
	air

# 数据库迁移
migrate:
	@echo "Running database migrations..."
	# 这里可以添加数据库迁移命令

# 代码格式化
fmt:
	@echo "Formatting code..."
	go fmt ./...

# 代码检查
lint:
	@echo "Linting code..."
	golangci-lint run

# 安装所有工具
tools:
	@echo "Installing development tools..."
	go install github.com/cosmtrek/air@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/swaggo/swag/cmd/swag@latest

help:
	@echo "Available commands:"
	@echo "  build        - Build the application"
	@echo "  run          - Run the application"
	@echo "  swagger      - Generate Swagger documentation"
	@echo "  test         - Run tests"
	@echo "  clean        - Clean build files"
	@echo "  docker-build - Build Docker image"
	@echo "  docker-run   - Run Docker container"
	@echo "  deps         - Install development dependencies"
	@echo "  dev          - Start development server with auto-reload"
	@echo "  migrate      - Run database migrations"
	@echo "  fmt          - Format code"
	@echo "  lint         - Lint code"
	@echo "  tools        - Install all development tools"
	@echo "  help         - Show this help message"