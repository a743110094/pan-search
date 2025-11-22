# 网盘资源搜索平台 - 后端 API

基于 Golang + Gin 开发的网盘资源搜索平台后端 API 服务，提供资源搜索、分类管理、热门推荐、求助请求等功能。

## 🚀 功能特性

- 🔍 **资源搜索**: 支持关键词搜索、分类筛选、多种排序方式
- 📊 **热门推荐**: 基于浏览量和下载量的热门资源推荐
- 📂 **分类管理**: 完整的资源分类体系
- 📝 **求助请求**: 用户可提交资源求助请求
- 🔐 **认证授权**: JWT Token 认证，支持管理员权限
- 📚 **API文档**: 完整的 Swagger API 文档
- 🛡️ **安全防护**: CORS 支持，输入验证

## 🛠️ 技术栈

- **Web 框架**: [Gin](https://github.com/gin-gonic/gin)
- **数据库**: MySQL + [GORM](https://gorm.io/)
- **认证**: JWT
- **API 文档**: [Swagger](https://swagger.io/)
- **配置管理**: YAML + 环境变量

## 📁 项目结构

```
backend/
├── main.go                 # 程序入口
├── go.mod / go.sum         # Go 依赖管理
├── .env.example            # 环境变量示例
├── config/                 # 配置管理
│   ├── config.go
│   └── config.yaml
├── database/               # 数据库连接
│   ├── database.go
│   └── pan_search_minimal.sql
├── models/                 # 数据模型
│   └── models.go
├── handlers/               # 请求处理器
│   ├── resource.go
│   ├── category.go
│   ├── request.go
│   └── search.go
├── routes/                 # 路由定义
│   └── routes.go
├── middleware/             # 中间件
│   ├── auth.go
│   └── cors.go
├── common/                 # 通用组件
│   └── response.go
└── deploy/                 # 部署配置
    ├── Dockerfile
    ├── docker-compose.yml
    ├── Makefile
    ├── DEPLOYMENT.md
    └── DEPLOYMENT_GUIDE.md
```

## 🔧 快速开始

### 前置要求

- Go 1.21+
- MySQL 5.7+

### 安装依赖

```bash
go mod tidy
```

### 配置

1. 复制环境变量文件：
```bash
cp .env.example .env
```

2. 修改配置文件 `config/config.yaml` 或环境变量：
```yaml
database:
  host: "localhost"
  port: 3306
  username: "root"
  password: ""
  database: "pan_search"
```

### 运行服务

```bash
# 开发模式
go run main.go

# 生产模式
go build -o pan-search-api main.go
./pan-search-api
```

服务启动后访问：
- API服务: http://localhost:8080
- Swagger文档: http://localhost:8080/swagger/index.html

## 🐳 使用 Docker

### 单独运行后端

```dockerfile
FROM golang:1.21-alpine

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o pan-search-api main.go

EXPOSE 8080
CMD ["./pan-search-api"]
```

### 使用 docker-compose

```bash
cd deploy
docker-compose up -d
```

## 📊 数据库设计

数据库设计详见 `database/pan_search_minimal.sql`，包含以下表：

- `categories` - 分类表
- `resources` - 资源表
- `resource_tags` - 资源标签表
- `help_requests` - 求助请求表
- `download_records` - 下载记录表
- `search_records` - 搜索记录表
- `users` - 用户表（预留）
- `system_configs` - 系统配置表

## 📡 API 接口

### 无需认证的接口

| 接口 | 方法 | 描述 |
|------|------|------|
| `/api/v1/resources/search` | GET | 搜索资源 |
| `/api/v1/resources/hot` | GET | 热门推荐 |
| `/api/v1/categories` | GET | 分类列表 |
| `/api/v1/requests` | POST | 提交求助 |
| `/api/v1/search/suggestions` | GET | 搜索建议 |

### 需要认证的接口

| 接口 | 方法 | 描述 |
|------|------|------|
| `/api/v1/resources/{id}/download` | POST | 记录下载 |

### 管理员接口

| 接口 | 方法 | 描述 |
|------|------|------|
| `/api/v1/admin/*` | 所有 | 管理功能 |

## 🔨 开发指南

### 添加新接口

1. 在 `models/models.go` 中定义请求/响应模型
2. 在 `handlers/` 中创建处理器
3. 在 `routes/routes.go` 中注册路由
4. 添加 Swagger 注释

### 认证中间件

- `middleware.AuthMiddleware()` - 需要认证
- `middleware.AdminAuthMiddleware()` - 需要管理员权限
- `middleware.OptionalAuthMiddleware()` - 可选认证

### 统一响应格式

使用 `common` 包中的响应函数：

```go
common.Success(c, data)
common.SuccessWithPagination(c, list, page, pageSize, total)
common.Error(c, code, message)
common.BadRequest(c, message)
// ...
```

## 🛠️ Makefile 命令

项目提供了便捷的 Makefile 命令：

```bash
cd deploy

make build          # 构建应用
make run            # 运行应用
make test           # 运行测试
make clean          # 清理构建文件
make swagger        # 生成 Swagger 文档
make docker-build   # 构建 Docker 镜像
make docker-run     # 运行 Docker 容器
make deps           # 安装开发依赖
make dev            # 开发模式（自动重载）
make migrate        # 数据库迁移
make fmt            # 代码格式化
make lint           # 代码检查
make tools          # 安装所有开发工具
make help           # 查看所有可用命令
```

## 🌍 环境变量

所有配置都支持环境变量覆盖，详见 `.env.example`：

```bash
# 应用配置
APP_NAME=pan-search-api
APP_VERSION=1.0.0
APP_PORT=8080
APP_MODE=debug

# 数据库配置
DB_HOST=localhost
DB_PORT=3306
DB_USERNAME=root
DB_PASSWORD=
DB_DATABASE=pan_search

# JWT 配置
JWT_SECRET=your-secret-key
JWT_EXPIRE=24h
```

## 📝 日志

应用使用结构化日志，支持多种输出格式：

- `stdout` - 输出到控制台
- `file` - 输出到文件
- `json` - JSON 格式日志

日志级别：`debug`, `info`, `warn`, `error`

## 🔐 安全

- JWT Token 认证
- CORS 跨域支持
- 输入验证和过滤
- SQL 注入防护（GORM）
- XSS 防护

## 🧪 测试

```bash
# 运行所有测试
go test ./...

# 运行测试并查看覆盖率
go test -cover

# 运行特定包的测试
go test ./handlers
```

## 📦 部署

详见 `deploy/DEPLOYMENT.md` 和 `deploy/DEPLOYMENT_GUIDE.md`

## 📄 许可证

MIT License

## 👥 贡献

欢迎提交 Issue 和 Pull Request！

## 🙏 致谢

- [Gin](https://github.com/gin-gonic/gin) - 高性能 Go Web 框架
- [GORM](https://gorm.io/) - 强大的 Go ORM 库
- [Swagger](https://swagger.io/) - API 文档生成工具
