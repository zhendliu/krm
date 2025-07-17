# KRM - Kubernetes Resource Manager Makefile
# 作者: kunyu
# 描述: 用于构建、运行和部署KRM项目的Makefile

# 变量定义
PROJECT_NAME := krm
BACKEND_DIR := backend
FRONTEND_DIR := frontend
BACKEND_BINARY := krm-backend
FRONTEND_DIST := dist

# Go相关变量
GO := go
GOOS := $(shell go env GOOS)
GOARCH := $(shell go env GOARCH)
GO_VERSION := 1.20
# Docker相关变量
DOCKER_REGISTRY := localhost:5000ACKEND_IMAGE := $(PROJECT_NAME)-backend
FRONTEND_IMAGE := $(PROJECT_NAME)-frontend
IMAGE_TAG := latest

# 端口配置
BACKEND_PORT := 8080
FRONTEND_PORT := 5174 默认目标
.PHONY: help
help: ## 显示帮助信息
	@echoKRM - Kubernetes Resource Manager"
	@echo "==================================
	@echo 可用的命令:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$$(MAKEFILE_LIST) | sort | awkBEGIN[object Object]FS = ":.*?## };[object Object]printf \03336m%-20s\330m %s\n", $$1, $$2}

# 开发环境命令
.PHONY: dev
dev: ## 启动开发环境（后端+前端）
	@echo 启动开发环境..."
	@make -j2 dev-backend dev-frontend

.PHONY: dev-backend
dev-backend: ## 启动后端开发服务器
	@echo启动后端开发服务器..."
	@cd $(BACKEND_DIR) && $(GO) run main.go

.PHONY: dev-frontend
dev-frontend: ## 启动前端开发服务器
	@echo启动前端开发服务器..."
	@cd $(FRONTEND_DIR) && npm run dev

# 构建命令
.PHONY: build
build: ## 构建整个项目
	@echo 构建整个项目..."
	@make build-backend
	@make build-frontend

.PHONY: build-backend
build-backend: ## 构建后端二进制文件
	@echo构建后端二进制文件..."
	@cd $(BACKEND_DIR) && CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build -ldflags="-s -w" -o $(BACKEND_BINARY) main.go
	@echo 后端构建完成: $(BACKEND_DIR)/$(BACKEND_BINARY)"

.PHONY: build-frontend
build-frontend: ## 构建前端生产版本
	@echo 构建前端生产版本..."
	@cd $(FRONTEND_DIR) && npm run build
	@echo "前端构建完成: $(FRONTEND_DIR)/$(FRONTEND_DIST)"

# 清理命令
.PHONY: clean
clean: ## 清理构建产物
	@echo 清理构建产物..."
	@rm -f $(BACKEND_DIR)/$(BACKEND_BINARY)
	@rm -rf $(FRONTEND_DIR)/$(FRONTEND_DIST)
	@rm -rf $(FRONTEND_DIR)/node_modules
	@echo "清理完成"

.PHONY: clean-backend
clean-backend: ## 清理后端构建产物
	@echo 清理后端构建产物..."
	@rm -f $(BACKEND_DIR)/$(BACKEND_BINARY)
	@echo后端清理完成"

.PHONY: clean-frontend
clean-frontend: ## 清理前端构建产物
	@echo 清理前端构建产物...@rm -rf $(FRONTEND_DIR)/$(FRONTEND_DIST)
	@echo前端清理完成"

# 依赖管理
.PHONY: deps
deps: ## 安装所有依赖
	@echo 安装所有依赖...@make deps-backend
	@make deps-frontend

.PHONY: deps-backend
deps-backend: ## 安装后端依赖
	@echo "安装后端依赖..."
	@cd $(BACKEND_DIR) && $(GO) mod download
	@echo 后端依赖安装完成"

.PHONY: deps-frontend
deps-frontend: ## 安装前端依赖
	@echo "安装前端依赖..."
	@cd $(FRONTEND_DIR) && npm install
	@echo 前端依赖安装完成"

# 测试命令
.PHONY: test
test: ## 运行所有测试
	@echo 运行所有测试...@make test-backend
	@make test-frontend

.PHONY: test-backend
test-backend: ## 运行后端测试
	@echo "运行后端测试..."
	@cd $(BACKEND_DIR) && $(GO) test -v ./...

.PHONY: test-frontend
test-frontend: ## 运行前端测试
	@echo "运行前端测试..."
	@cd $(FRONTEND_DIR) && npm test

# 代码质量检查
.PHONY: lint
lint: ## 运行代码质量检查
	@echo "运行代码质量检查...@make lint-backend
	@make lint-frontend

.PHONY: lint-backend
lint-backend: ## 运行后端代码质量检查
	@echo 运行后端代码质量检查..."
	@cd $(BACKEND_DIR) && golangci-lint run

.PHONY: lint-frontend
lint-frontend: ## 运行前端代码质量检查
	@echo 运行前端代码质量检查..."
	@cd $(FRONTEND_DIR) && npm run lint

# Docker相关命令
.PHONY: docker-build
docker-build: ## 构建Docker镜像
	@echo 构建Docker镜像..."
	@make docker-build-backend
	@make docker-build-frontend

.PHONY: docker-build-backend
docker-build-backend: ## 构建后端Docker镜像
	@echo 构建后端Docker镜像...
	@docker build -t $(BACKEND_IMAGE):$(IMAGE_TAG) $(BACKEND_DIR)

.PHONY: docker-build-frontend
docker-build-frontend: ## 构建前端Docker镜像
	@echo 构建前端Docker镜像...
	@docker build -t $(FRONTEND_IMAGE):$(IMAGE_TAG) $(FRONTEND_DIR)

.PHONY: docker-run
docker-run: ## 运行Docker容器
	@echo 运行Docker容器...@docker-compose up -d

.PHONY: docker-stop
docker-stop: ## 停止Docker容器
	@echo 停止Docker容器...
	@docker-compose down

# 部署命令
.PHONY: deploy
deploy: ## 部署到Kubernetes
	@echo "部署到Kubernetes..."
	@kubectl apply -f k8s/

.PHONY: deploy-backend
deploy-backend: ## 部署后端到Kubernetes
	@echo "部署后端到Kubernetes..."
	@kubectl apply -f k8s/backend/

.PHONY: deploy-frontend
deploy-frontend: ## 部署前端到Kubernetes
	@echo "部署前端到Kubernetes..."
	@kubectl apply -f k8/frontend/

# 工具命令
.PHONY: fmt
fmt: ## 格式化代码
	@echo 格式化代码..."
	@cd $(BACKEND_DIR) && $(GO) fmt ./...
	@cd $(FRONTEND_DIR) && npm run format

.PHONY: generate
generate: ## 生成代码
	@echo生成代码..."
	@cd $(BACKEND_DIR) && $(GO) generate ./...

.PHONY: swagger
swagger: ## 生成API文档
	@echo生成API文档..."
	@cd $(BACKEND_DIR) && swag init

# 监控和日志
.PHONY: logs
logs: ## 查看应用日志
	@echo "查看应用日志...
	@docker-compose logs -f

.PHONY: logs-backend
logs-backend: ## 查看后端日志
	@echo "查看后端日志...
	@docker-compose logs -f backend

.PHONY: logs-frontend
logs-frontend: ## 查看前端日志
	@echo "查看前端日志...
	@docker-compose logs -f frontend

# 数据库相关
.PHONY: db-migrate
db-migrate: ## 运行数据库迁移
	@echo运行数据库迁移..."
	@cd $(BACKEND_DIR) && $(GO) run cmd/migrate/main.go

.PHONY: db-seed
db-seed: ## 填充数据库种子数据
	@echo填充数据库种子数据..."
	@cd $(BACKEND_DIR) && $(GO) run cmd/seed/main.go

# 性能测试
.PHONY: bench
bench: ## 运行性能测试
	@echo "运行性能测试..."
	@cd $(BACKEND_DIR) && $(GO) test -bench=. ./...

# 安全检查
.PHONY: security
security: ## 运行安全检查
	@echo "运行安全检查..."
	@cd $(BACKEND_DIR) && gosec ./...
	@cd $(FRONTEND_DIR) && npm audit

# 版本信息
.PHONY: version
version: ## 显示版本信息
	@echoKRM - Kubernetes Resource Manager"
	@echo "版本: $(shell git describe --tags --always --dirty)
	@echoGo版本:$(shell go version)"
	@echo "Node版本: $(shell node --version)"
	@echo "NPM版本: $(shell npm --version)"

# 安装工具
.PHONY: install-tools
install-tools: ## 安装开发工具
	@echo "安装开发工具..."
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@go install github.com/swaggo/swag/cmd/swag@latest
	@go install github.com/securecodewarrior/gosec/v2/cmd/gosec@latest
	@echo开发工具安装完成

#初始化项目
.PHONY: init
init: ## 初始化项目
	@echo初始化项目...
	@make install-tools
	@make deps
	@make build
	@echo项目初始化完成

# 生产环境构建
.PHONY: prod-build
prod-build: ## 生产环境构建
	@echo 生产环境构建..."
	@make clean
	@make deps
	@make build
	@make docker-build
	@echo 生产环境构建完成

# 快速启动（用于演示）
.PHONY: demo
demo: ## 快速启动演示环境
	@echo 启动演示环境..."
	@make build
	@make docker-run
	@echo演示环境启动完成，访问 http://localhost:$(FRONTEND_PORT)

# 默认目标
.DEFAULT_GOAL := help 