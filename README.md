# 网盘资源搜索平台

[![Go Version](https://img.shields.io/badge/Go-1.21-blue.svg)](https://golang.org/dl/)
[![Vue.js](https://img.shields.io/badge/Vue.js-3.3-green.svg)](https://vuejs.org/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

一个基于 Go + Vue.js 开发的全栈网盘资源搜索平台，支持资源搜索、分类管理、热门推荐等功能。

## ✨ 功能特性

- 🔍 **智能搜索**: 支持关键词搜索、分类筛选、多种排序方式
- 📊 **热门推荐**: 基于浏览量和下载量的热门资源推荐
- 📂 **分类管理**: 完整的资源分类体系
- 📝 **求助请求**: 用户可提交资源求助请求
- 🎨 **现代化UI**: 基于 Vue 3 + Vite 的响应式界面
- 🔐 **认证授权**: JWT Token 认证，支持管理员权限
- 📚 **API文档**: 完整的 Swagger API 文档
- 🐳 **容器化部署**: 支持 Docker 一键部署

## 🏗️ 项目结构

```
pan-search/
├── 📄 README.md                    # 项目总体说明（本文件）
├── 📄 API_DOCUMENTATION.md         # API 文档
├── 📄 PROJECT_SUMMARY.md          # 项目总结
├── 📁 backend/                    # 后端服务 (Go)
│   ├── 📄 README.md               # 后端说明文档
│   ├── 📄 main.go                 # Go 入口文件
│   ├── 📄 go.mod / go.sum         # Go 依赖管理
│   ├── 📄 .env.example            # 环境变量示例
│   ├── 📁 config/                 # 配置文件
│   ├── 📁 database/               # 数据库文件
│   ├── 📁 handlers/               # 路由处理
│   ├── 📁 middleware/             # 中间件
│   ├── 📁 models/                 # 数据模型
│   ├── 📁 routes/                 # 路由定义
│   ├── 📁 common/                 # 公共模块
│   └── 📁 deploy/                 # 部署配置
│       ├── 📄 Dockerfile
│       ├── 📄 docker-compose.yml
│       ├── 📄 Makefile
│       ├── 📄 DEPLOYMENT.md
│       └── 📄 DEPLOYMENT_GUIDE.md
└── 📁 frontend/                   # 前端应用 (Vue)
    ├── 📄 README.md               # 前端说明文档
    ├── 📄 index.html              # 主页面
    ├── 📄 package.json            # 前端依赖
    ├── 📄 vite.config.js          # Vite 配置
    ├── 📄 yarn.lock               # 锁定版本
    └── 📁 src/                    # 源代码
        ├── 📄 App.vue
        ├── 📄 main.js
        ├── 📁 components/         # 公共组件
        ├── 📁 views/              # 页面视图
        ├── 📁 router/             # 路由配置
        ├── 📁 services/           # API 服务
        ├── 📁 utils/              # 工具函数
        └── 📄 style.css           # 样式文件
```

## 🚀 快速开始

### 系统要求

- Go 1.21+
- Node.js 16+ / Yarn
- MySQL 5.7+
- Git

### 克隆项目

```bash
git clone https://github.com/a743110094/pan-search.git
cd pan-search
```

### 后端启动

```bash
cd backend

# 安装依赖
go mod tidy

# 配置环境
cp .env.example .env
# 编辑 .env 文件，配置数据库信息

# 运行服务
go run main.go
# 或使用 Make
cd deploy && make run
```

后端服务将启动在 http://localhost:8080

### 前端启动

```bash
cd frontend

# 安装依赖
yarn install
# 或
npm install

# 启动开发服务器
yarn dev
# 或
npm run dev
```

前端应用将启动在 http://localhost:3000

### 使用 Docker 启动（推荐）

一键启动前后端和数据库：

```bash
cd backend/deploy

# 启动所有服务
docker-compose up -d

# 查看服务状态
docker-compose ps

# 查看日志
docker-compose logs -f api
```

服务访问地址：
- 前端: http://localhost:3000
- 后端 API: http://localhost:8080
- API 文档: http://localhost:8080/swagger/index.html

## 🛠️ 技术栈

### 后端 (Backend)

- **Web 框架**: [Gin](https://github.com/gin-gonic/gin)
- **数据库**: MySQL + [GORM](https://gorm.io/)
- **认证**: JWT
- **API 文档**: [Swagger](https://swagger.io/)
- **配置管理**: YAML + 环境变量

### 前端 (Frontend)

- **框架**: [Vue 3](https://vuejs.org/)
- **构建工具**: [Vite](https://vitejs.dev/)
- **路由**: [Vue Router](https://router.vuejs.org/)
- **HTTP 客户端**: Fetch API
- **样式**: 原生 CSS

## 📖 文档说明

- **[README.md](./README.md)** - 项目总体介绍（本文件）
- **[backend/README.md](./backend/README.md)** - 后端开发指南
- **[frontend/README.md](./frontend/README.md)** - 前端开发指南
- **[API_DOCUMENTATION.md](./API_DOCUMENTATION.md)** - 完整 API 文档
- **[PROJECT_SUMMARY.md](./PROJECT_SUMMARY.md)** - 项目详细总结
- **[deploy/DEPLOYMENT.md](./backend/deploy/DEPLOYMENT.md)** - 部署指南

## 🔧 开发指南

### 开发后端

详见 [backend/README.md](./backend/README.md)

### 开发前端

详见 [frontend/README.md](./frontend/README.md)

### API 开发

1. 在 `backend/models/models.go` 中定义请求/响应模型
2. 在 `backend/handlers/` 中创建处理器
3. 在 `backend/routes/routes.go` 中注册路由
4. 添加 Swagger 注释生成文档

### 前端开发

1. 在 `frontend/src/views/` 中创建新页面
2. 在 `frontend/src/components/` 中创建可复用组件
3. 在 `frontend/src/services/api.js` 中封装 API 调用

## 🐳 部署

支持多种部署方式：

### 1. Docker 部署（推荐）

```bash
cd backend/deploy
docker-compose up -d
```

### 2. 生产环境构建

```bash
# 构建后端
cd backend
go build -o pan-search-api main.go

# 构建前端
cd frontend
yarn build
```

### 3. 使用 Makefile

```bash
cd backend/deploy
make build    # 构建
make run      # 运行
make test     # 测试
make clean    # 清理
```

详细部署说明请参考：[deploy/DEPLOYMENT_GUIDE.md](./backend/deploy/DEPLOYMENT_GUIDE.md)

## 📝 API 文档

启动后端服务后，访问 http://localhost:8080/swagger/index.html 查看完整 API 文档。

### 核心接口

| 接口 | 方法 | 描述 |
|------|------|------|
| `/api/v1/resources/search` | GET | 搜索资源 |
| `/api/v1/resources/hot` | GET | 热门推荐 |
| `/api/v1/categories` | GET | 分类列表 |
| `/api/v1/requests` | POST | 提交求助 |

## 🤝 贡献指南

1. Fork 本仓库
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 创建 Pull Request

## 📄 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情

## 👥 作者

- **a743110094** - *初始开发* - [GitHub](https://github.com/a743110094)

## 🙏 致谢

- [Gin](https://github.com/gin-gonic/gin) - 高性能 Go Web 框架
- [Vue.js](https://vuejs.org/) - 渐进式 JavaScript 框架
- [Vite](https://vitejs.dev/) - 下一代前端构建工具
- [MySQL](https://www.mysql.com/) - 关系型数据库
