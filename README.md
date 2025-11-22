# 网盘资源搜索平台后端API

基于Golang开发的网盘资源搜索平台后端API服务，提供资源搜索、热门推荐、分类管理、求助请求等功能。

## 功能特性

- 🔍 **资源搜索**: 支持关键词搜索、分类筛选、多种排序方式
- 📊 **热门推荐**: 基于浏览量和下载量的热门资源推荐
- 📂 **分类管理**: 完整的资源分类体系
- 📝 **求助请求**: 用户可提交资源求助请求
- 🔐 **认证授权**: JWT Token认证，支持管理员权限
- 📚 **API文档**: 完整的Swagger API文档
- 🛡️ **安全防护**: CORS支持，输入验证

## 技术栈

- **框架**: Gin
- **数据库**: MySQL + GORM
- **认证**: JWT
- **文档**: Swagger
- **配置**: YAML + 环境变量

## 项目结构

```
pan-search-api/
├── main.go                 # 程序入口
├── config/                 # 配置管理
│   ├── config.go
│   └── config.yaml
├── database/               # 数据库连接
│   └── database.go
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
└── docs/                   # Swagger文档
```

## 快速开始

### 前置要求

- Go 1.21+
- MySQL 5.7+
- 已创建数据库 `pan_search`（可使用 `database/pan_search_minimal.sql` 初始化）

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

## API接口

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

## 数据库设计

数据库设计详见 `database/pan_search_minimal.sql`，包含以下表：

- `categories` - 分类表
- `resources` - 资源表
- `resource_tags` - 资源标签表
- `help_requests` - 求助请求表
- `download_records` - 下载记录表
- `search_records` - 搜索记录表
- `users` - 用户表（预留）
- `system_configs` - 系统配置表

## 开发指南

### 添加新接口

1. 在 `models/models.go` 中定义请求/响应模型
2. 在 `handlers/` 中创建处理器
3. 在 `routes/routes.go` 中注册路由
4. 添加Swagger注释

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

## 部署

### Docker部署

```dockerfile
FROM golang:1.21-alpine

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o pan-search-api main.go

EXPOSE 8080
CMD ["./pan-search-api"]
```

### 环境变量

所有配置都支持环境变量覆盖，详见 `.env.example`。

## 许可证

MIT License