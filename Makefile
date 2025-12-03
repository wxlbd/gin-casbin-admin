# Makefile for Gin-Casbin-Admin
# 基于 DDD 架构的企业级后台管理系统

# 变量定义
APP_NAME := gin-casbin-admin
MAIN_FILE := cmd/server/main.go
BUILD_DIR := build
BINARY_NAME := $(APP_NAME)
CONFIG_FILE := configs/config.yaml
GO := go
AIR := air
SWAG := swag

# Go 相关变量
GOPATH := $(shell go env GOPATH)
GOBIN := $(GOPATH)/bin
GOOS := $(shell go env GOOS)
GOARCH := $(shell go env GOARCH)

# 版本信息
VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
BUILD_TIME := $(shell date -u '+%Y-%m-%d_%H:%M:%S')
GIT_COMMIT := $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")

# 构建参数
LDFLAGS := -X 'main.Version=$(VERSION)' \
           -X 'main.BuildTime=$(BUILD_TIME)' \
           -X 'main.GitCommit=$(GIT_COMMIT)'

# 默认目标
.PHONY: all
all: clean deps build

# 清理构建产物
.PHONY: clean
clean:
	@echo "🧹 清理构建产物..."
	@rm -rf $(BUILD_DIR)
	@rm -f $(BINARY_NAME)
	@echo "✅ 清理完成"

# 下载依赖
.PHONY: deps
deps:
	@echo "📦 下载项目依赖..."
	@$(GO) mod download
	@$(GO) mod tidy
	@echo "✅ 依赖下载完成"

# 更新依赖
.PHONY: update-deps
update-deps:
	@echo "🔄 更新项目依赖..."
	@$(GO) get -u ./...
	@$(GO) mod tidy
	@echo "✅ 依赖更新完成"

# 构建项目
.PHONY: build
build: deps
	@echo "🔨 构建项目..."
	@mkdir -p $(BUILD_DIR)
	@CGO_ENABLED=0 $(GO) build -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_FILE)
	@echo "✅ 构建完成: $(BUILD_DIR)/$(BINARY_NAME)"

# 跨平台构建
.PHONY: build-all
build-all: deps
	@echo "🔨 跨平台构建..."
	@mkdir -p $(BUILD_DIR)
	# Linux AMD64
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO) build -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 $(MAIN_FILE)
	# Linux ARM64
	@CGO_ENABLED=0 GOOS=linux GOARCH=arm64 $(GO) build -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/$(BINARY_NAME)-linux-arm64 $(MAIN_FILE)
	# Darwin AMD64
	@CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GO) build -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 $(MAIN_FILE)
	# Darwin ARM64
	@CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 $(GO) build -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 $(MAIN_FILE)
	# Windows AMD64
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GO) build -ldflags "$(LDFLAGS)" -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe $(MAIN_FILE)
	@echo "✅ 跨平台构建完成"

# 开发模式（热重载）
.PHONY: dev
dev:
	@echo "🚀 启动开发模式..."
	@if command -v $(AIR) >/dev/null 2>&1; then \
		$(AIR); \
	else \
		echo "❌ air 未安装，正在安装..."; \
		go install github.com/cosmtrek/air@latest; \
		$(AIR); \
	fi

# 生成 Swagger 文档
.PHONY: swagger
swagger:
	@echo "📚 生成 Swagger 文档..."
	@if command -v $(SWAG) >/dev/null 2>&1; then \
		$(SWAG) init -g $(MAIN_FILE) -o docs/swagger; \
	else \
		echo "❌ swag 未安装，正在安装..."; \
		go install github.com/swaggo/swag/cmd/swag@latest; \
		$(SWAG) init -g $(MAIN_FILE) -o docs/swagger; \
	fi
	@echo "✅ Swagger 文档生成完成"

# 运行测试
.PHONY: test
test:
	@echo "🧪 运行测试..."
	@$(GO) test -v -race -coverprofile=coverage.out ./...
	@echo "✅ 测试完成"

# 查看测试覆盖率
.PHONY: coverage
coverage: test
	@echo "📊 查看测试覆盖率..."
	@$(GO) tool cover -html=coverage.out -o coverage.html
	@echo "✅ 覆盖率报告生成: coverage.html"

# 运行基准测试
.PHONY: bench
bench:
	@echo "⚡ 运行基准测试..."
	@$(GO) test -bench=. -benchmem ./...
	@echo "✅ 基准测试完成"

# 代码格式化
.PHONY: fmt
fmt:
	@echo "🎨 代码格式化..."
	@$(GO) fmt ./...
	@echo "✅ 代码格式化完成"

# 代码静态检查
.PHONY: lint
lint:
	@echo "🔍 代码静态检查..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "❌ golangci-lint 未安装，请先安装"; \
		echo "安装命令: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
	fi

# 代码审查
.PHONY: review
review: fmt lint
	@echo "🔍 代码审查完成"

# 安全扫描
.PHONY: security
security:
	@echo "🔒 安全扫描..."
	@if command -v gosec >/dev/null 2>&1; then \
		gosec ./...; \
	else \
		echo "❌ gosec 未安装，正在安装..."; \
		go install github.com/securego/gosec/v2/cmd/gosec@latest; \
		gosec ./...; \
	fi

# 依赖漏洞扫描
.PHONY: vuln
vuln:
	@echo "🛡️ 依赖漏洞扫描..."
	@if command -v govulncheck >/dev/null 2>&1; then \
		govulncheck ./...; \
	else \
		echo "❌ govulncheck 未安装，正在安装..."; \
		go install golang.org/x/vuln/cmd/govulncheck@latest; \
		govulncheck ./...; \
	fi

# 数据库迁移
.PHONY: migrate-up
migrate-up: build
	@echo "🗄️ 执行数据库迁移..."
	@$(BUILD_DIR)/$(BINARY_NAME) migrate up

# 数据库回滚
.PHONY: migrate-down
migrate-down: build
	@echo "🗄️ 回滚数据库..."
	@$(BUILD_DIR)/$(BINARY_NAME) migrate down

# 创建新的迁移文件
.PHONY: migrate-create
migrate-create:
	@if [ -z "$(name)" ]; then \
		echo "❌ 请指定迁移名称: make migrate-create name=add_user_table"; \
	else \
		echo "📝 创建迁移文件: $(name)"; \
		migrate create -ext sql -dir cmd/server/migrations -seq $(name); \
	fi

# 启动服务
.PHONY: run
run: build
	@echo "🚀 启动服务..."
	@$(BUILD_DIR)/$(BINARY_NAME) start

# 后台启动服务
.PHONY: run-daemon
run-daemon: build
	@echo "👻 后台启动服务..."
	@nohup $(BUILD_DIR)/$(BINARY_NAME) start >$(APP_NAME).log 2>&1 &
	@echo "✅ 服务已后台启动，日志文件: $(APP_NAME).log"

# 停止后台服务
.PHONY: stop
stop:
	@echo "⏹️ 停止后台服务..."
	@pkill -f $(BINARY_NAME) || true
	@echo "✅ 后台服务已停止"

# 查看服务状态
.PHONY: status
status:
	@echo "📊 服务状态..."
	@if pgrep -f $(BINARY_NAME) >/dev/null; then \
		echo "✅ 服务正在运行"; \
	else \
		echo "❌ 服务未运行"; \
	fi

# 查看日志
.PHONY: logs
logs:
	@echo "📋 查看日志..."
	@tail -f $(APP_NAME).log

# Docker 构建
.PHONY: docker-build
docker-build:
	@echo "🐳 Docker 构建..."
	@docker build -t $(APP_NAME):$(VERSION) .
	@echo "✅ Docker 镜像构建完成: $(APP_NAME):$(VERSION)"

# Docker 运行
.PHONY: docker-run
docker-run:
	@echo "🐳 Docker 运行..."
	@docker run -d --name $(APP_NAME) -p 8080:8080 $(APP_NAME):$(VERSION)
	@echo "✅ Docker 容器启动完成"

# Docker Compose 启动
.PHONY: compose-up
compose-up:
	@echo "🐳 Docker Compose 启动..."
	@docker-compose up -d
	@echo "✅ Docker Compose 启动完成"

# Docker Compose 停止
.PHONY: compose-down
compose-down:
	@echo "🐳 Docker Compose 停止..."
	@docker-compose down
	@echo "✅ Docker Compose 停止完成"

# 清理 Docker
.PHONY: docker-clean
docker-clean:
	@echo "🧹 清理 Docker..."
	@docker stop $(APP_NAME) 2>/dev/null || true
	@docker rm $(APP_NAME) 2>/dev/null || true
	@docker rmi $(APP_NAME):$(VERSION) 2>/dev/null || true
	@echo "✅ Docker 清理完成"

# Wire 依赖注入生成
.PHONY: wire
wire:
	@echo "🔗 生成 Wire 依赖注入代码..."
	@if command -v wire >/dev/null 2>&1; then \
		cd cmd/server/wire && wire; \
	else \
		echo "❌ wire 未安装，正在安装..."; \
		go install github.com/google/wire/cmd/wire@latest; \
		cd cmd/server/wire && wire; \
	fi
	@echo "✅ Wire 代码生成完成"

# 生成所有代码
.PHONY: generate
generate: wire swagger
	@echo "🎯 生成所有代码..."
	@$(GO) generate ./...
	@echo "✅ 代码生成完成"

# 清理模块缓存
.PHONY: clean-mod
clean-mod:
	@echo "🧹 清理模块缓存..."
	@$(GO) clean -modcache
	@echo "✅ 模块缓存清理完成"

# 查看模块依赖
.PHONY: deps-graph
deps-graph:
	@echo "📊 模块依赖图..."
	@$(GO) mod graph

# 模块验证
.PHONY: verify
verify:
	@echo "✅ 验证模块依赖..."
	@$(GO) mod verify

# 更新所有工具
.PHONY: update-tools
update-tools:
	@echo "🔧 更新开发工具..."
	@go install github.com/cosmtrek/air@latest
	@go install github.com/swaggo/swag/cmd/swag@latest
	@go install github.com/google/wire/cmd/wire@latest
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@go install github.com/securego/gosec/v2/cmd/gosec@latest
	@go install golang.org/x/vuln/cmd/govulncheck@latest
	@echo "✅ 开发工具更新完成"

# 一键初始化开发环境
.PHONY: init-dev
init-dev: update-tools deps generate
	@echo "🚀 开发环境初始化完成"
	@echo "📋 可用命令:"
	@echo "  make dev          - 启动热重载开发模式"
	@echo "  make build        - 构建项目"
	@echo "  make test         - 运行测试"
	@echo "  make run          - 运行服务"

# 一键部署
.PHONY: deploy
deploy: clean build migrate-up
	@echo "🚀 准备部署..."
	@echo "✅ 构建和迁移完成，可以启动服务了"

# 完整 CI 流程
.PHONY: ci
ci: deps generate fmt lint test security
	@echo "✅ CI 流程完成"

# 发布流程
.PHONY: release
release: clean verify ci build-all
	@echo "🎉 发布版本: $(VERSION)"
	@echo "✅ 发布构建完成"

# 帮助信息
.PHONY: help
help:
	@echo "🎯 Gin-Casbin-Admin Makefile"
	@echo ""
	@echo "📦 依赖管理:"
	@echo "  make deps         - 下载项目依赖"
	@echo "  make update-deps  - 更新项目依赖"
	@echo "  make clean-mod    - 清理模块缓存"
	@echo "  make verify       - 验证模块依赖"
	@echo ""
	@echo "🔨 构建编译:"
	@echo "  make build        - 构建项目"
	@echo "  make build-all    - 跨平台构建"
	@echo "  make clean        - 清理构建产物"
	@echo ""
	@echo "🚀 开发运行:"
	@echo "  make dev          - 启动热重载开发模式"
	@echo "  make run          - 运行服务"
	@echo "  make run-daemon   - 后台运行服务"
	@echo "  make stop         - 停止后台服务"
	@echo "  make status       - 查看服务状态"
	@echo "  make logs         - 查看日志"
	@echo ""
	@echo "🧪 测试质量:"
	@echo "  make test         - 运行测试"
	@echo "  make coverage     - 查看测试覆盖率"
	@echo "  make bench        - 运行基准测试"
	@echo "  make fmt          - 代码格式化"
	@echo "  make lint         - 代码静态检查"
	@echo "  make review       - 代码审查"
	@echo "  make security     - 安全扫描"
	@echo "  make vuln         - 依赖漏洞扫描"
	@echo ""
	@echo "🗄️ 数据库:"
	@echo "  make migrate-up   - 执行数据库迁移"
	@echo "  make migrate-down - 回滚数据库"
	@echo "  make migrate-create name=xxx - 创建迁移文件"
	@echo ""
	@echo "📚 文档代码:"
	@echo "  make swagger      - 生成 Swagger 文档"
	@echo "  make wire         - 生成 Wire 依赖注入代码"
	@echo "  make generate     - 生成所有代码"
	@echo ""
	@echo "🐳 Docker:"
	@echo "  make docker-build - Docker 构建"
	@echo "  make docker-run   - Docker 运行"
	@echo "  make compose-up   - Docker Compose 启动"
	@echo "  make compose-down - Docker Compose 停止"
	@echo "  make docker-clean - 清理 Docker"
	@echo ""
	@echo "🎯 高级命令:"
	@echo "  make init-dev     - 初始化开发环境"
	@echo "  make deploy       - 一键部署"
	@echo "  make ci           - 完整 CI 流程"
	@echo "  make release      - 发布流程"
	@echo "  make update-tools - 更新开发工具"
	@echo ""
	@echo "ℹ️ 其他:"
	@echo "  make help         - 显示帮助信息"
	@echo "  make deps-graph   - 查看模块依赖图"

# 默认显示帮助
.DEFAULT_GOAL := help