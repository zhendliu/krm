# KRM - Kubernetes资源管理系统

## 项目简介

KRM (Kubernetes Resource Manager) 是一个基于Web的Kubernetes资源管理平台，提供直观的图形化界面来管理多个Kubernetes集群中的各种资源。该系统采用前后端分离架构，支持多集群管理、资源监控和操作。

## 技术栈

### 后端 (Backend)
- **语言**: Go1.20
- **Web框架**: Gin
- **Kubernetes客户端**: client-go v028.0*认证**: JWT (golang-jwt/jwt/v5)
- **配置管理**: Viper
- **日志**: Logrus
- **Kubernetes工具库**: kubeutils

### 前端 (Frontend)
- **框架**: Vue30.34
- **构建工具**: Vite 4.4**UI组件库**: Element Plus2.30.8
- **状态管理**: Pinia 20.1**路由**: Vue Router42.4
- **HTTP客户端**: Axios 1.4
- **代码编辑器**: CodeMirror 6.0.1
- **样式**: Less 40.13

## 功能特性

### 🔐 认证与授权
- JWT Token认证
- 用户登录/登出
- 路由守卫保护

### 🏢 多集群管理
- 支持添加多个Kubernetes集群
- 集群状态监控
- 集群配置管理
- 集群切换功能

### 📦 集群资源管理
- **命名空间 (Namespace)**: 创建、删除、查看、编辑
- **节点 (Node)**: 查看节点列表、节点详情、节点编辑
- **Pod**: 创建、删除、查看、更新
- **Deployment**: 创建、删除、查看、更新
- **StatefulSet**: 创建、删除、查看、更新
- **DaemonSet**: 创建、删除、查看、更新
- **CronJob**: 创建、删除、查看、更新
- **ReplicaSet**: 查看列表

### 🌐 网络资源管理
- **Service**: 创建、删除、查看、更新
- **Ingress**: 创建、删除、查看、更新

### 💾 存储资源管理
- **PersistentVolume (PV)**: 创建、删除、查看、更新
- **PersistentVolumeClaim (PVC)**: 创建、删除、查看、更新
- **StorageClass**: 查看列表

### 🔧 配置资源管理
- **ConfigMap**: 创建、删除、查看、更新
- **Secret**: 创建、删除、查看、更新

### 🎨 用户界面特性
- 响应式设计
- 现代化UI界面
- 代码编辑器支持YAML编辑
- 集群和命名空间选择器
- 资源列表和详情视图
- 实时状态更新

## 项目结构

```
krm/
├── backend/                 # 后端服务
│   ├── config/             # 配置文件
│   ├── controllers/        # 控制器层
│   │   ├── auth/          # 认证控制器
│   │   ├── cluster/       # 集群管理
│   │   ├── deployment/    # Deployment管理
│   │   ├── namespace/     # 命名空间管理
│   │   ├── node/          # 节点管理
│   │   ├── pod/           # Pod管理
│   │   ├── service/       # Service管理
│   │   └── ...           # 其他资源控制器
│   ├── middlerwares/      # 中间件
│   ├── models/           # 数据模型
│   ├── routers/          # 路由定义
│   ├── utils/            # 工具函数
│   └── main.go           # 主入口文件
├── frontend/              # 前端应用
│   ├── src/
│   │   ├── api/          # API接口
│   │   ├── components/   # 公共组件
│   │   ├── config/       # 配置文件
│   │   ├── router/       # 路由配置
│   │   ├── store/        # 状态管理
│   │   ├── utils/        # 工具函数
│   │   └── view/         # 页面组件
│   └── package.json      # 依赖配置
└── README.md             # 项目文档
```

## 快速开始

### 环境要求
- Go 120s 16+
- Kubernetes集群访问权限

### 后端启动

1 进入后端目录
```bash
cd backend
```

2装依赖
```bash
go mod download
```

3群信息
   - 在 `config/meta.kubeconfig` 中配置管理集群的kubeconfig
   - 或通过环境变量设置

4动服务
```bash
go run main.go
```

### 前端启动

1 进入前端目录
```bash
cd frontend
```2. 安装依赖
```bash
npm install
```

3. 启动开发服务器
```bash
npm run dev
```

4构建生产版本
```bash
npm run build
```

## API接口

### 认证接口
- `POST /api/auth/login` - 用户登录
- `POST /api/auth/logout` - 用户登出

### 集群管理接口
- `GET /api/cluster/list` - 获取集群列表
- `POST /api/cluster/add` - 添加集群
- `PUT /api/cluster/update` - 更新集群
- `DELETE /api/cluster/delete` - 删除集群

### 资源管理接口
每个Kubernetes资源都提供标准的CRUD操作：
- `GET /api/[object Object]resource}/list` - 获取资源列表
- `GET /api/{resource}/get` - 获取单个资源
- `POST /api/{resource}/create` - 创建资源
- `PUT /api/{resource}/update` - 更新资源
- `DELETE /api/{resource}/delete` - 删除资源

## 部署

### Docker部署

1 构建后端镜像
```bash
cd backend
docker build -t krm-backend:v1 .
```

2 构建前端镜像
```bash
cd frontend
docker build -t krm-frontend:v1 .
```3. 使用Docker Compose启动
```yaml
version:3.8ervices:
  backend:
    image: krm-backend:v1    ports:
      - "8080:8080
    environment:
      - GIN_MODE=release
      
  frontend:
    image: krm-frontend:v1    ports:
      - "8080 depends_on:
      - backend
```

### Kubernetes部署

参考 `backend/README.md` 中的详细部署说明。

## 配置说明

### 后端配置
- 端口配置: 通过环境变量或配置文件设置
- 日志级别: 支持debug、info、warn、error级别
- 集群配置: 支持多集群kubeconfig配置

### 前端配置
- API地址配置: 在 `src/config/index.js` 中设置
- 菜单配置: 在 `src/config/menu.js` 中自定义菜单结构

## 开发指南

### 添加新的Kubernetes资源支持
1. 在后端添加控制器
2添加路由配置
3. 在前端添加对应的页面和API4. 更新菜单配置

### 代码规范
- 后端遵循Go语言规范
- 前端使用ESLint和Prettier
- 提交信息使用约定式提交

## 贡献指南

1 Fork 项目
2. 创建功能分支
3. 提交更改
45. 创建Pull Request

## 许可证

本项目采用 MIT 许可证。

## 联系方式

如有问题或建议，请通过以下方式联系：
- 提交Issue
- 发送邮件
- 参与讨论

---

**注意**: 这是一个用于学习和演示的Kubernetes资源管理平台，在生产环境中使用前请确保充分测试和配置安全措施。 