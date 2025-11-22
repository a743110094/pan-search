# 网盘资源搜索平台 - 部署指南

## 项目概述

基于Vite + Vue3开发的网盘资源搜索平台，具备资源搜索、资源求助、资源推荐等功能。

## 技术栈

- **前端**: Vue 3 + Vite + Vue Router
- **样式**: 纯CSS自定义UI组件
- **数据库**: MySQL 8.0+
- **部署**: 支持Docker部署

## 快速开始

### 1. 环境要求

- Node.js 16.0+
- MySQL 8.0+
- 现代浏览器（Chrome 80+, Firefox 75+, Safari 13+）

### 2. 前端部署

#### 开发环境

```bash
# 安装依赖
npm install

# 启动开发服务器
npm run dev

# 访问 http://localhost:5173
```

#### 生产构建

```bash
# 构建生产版本
npm run build

# 预览生产版本
npm run preview
```

### 3. 数据库部署

#### 自动部署（推荐）

**Windows系统**:
```bash
cd database
init_database.bat
```

**Linux/Mac系统**:
```bash
cd database
chmod +x init_database.sh
./init_database.sh
```

#### 手动部署

1. 创建数据库：
```sql
CREATE DATABASE pan_search CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

2. 执行SQL脚本：
```bash
mysql -u root -p pan_search < database/pan_search.sql
```

### 4. 数据库连接信息

- **主机**: localhost
- **端口**: 3306
- **数据库**: pan_search
- **用户名**: root
- **密码**: (空)

## 项目结构

```
pan-search-web/
├── src/                    # 源代码目录
│   ├── components/         # 公共组件
│   ├── views/             # 页面组件
│   ├── router/            # 路由配置
│   ├── services/          # API服务
│   ├── utils/             # 工具函数
│   └── style.css          # 全局样式
├── database/              # 数据库相关
│   ├── pan_search.sql     # 数据库建表脚本
│   ├── init_database.sh   # Linux初始化脚本
│   ├── init_database.bat  # Windows初始化脚本
│   └── README.md          # 数据库文档
├── public/                # 静态资源
├── dist/                  # 构建输出目录
└── 配置文件
    ├── package.json
    ├── vite.config.js
    ├── index.html
    └── README.md
```

## 核心功能

### 1. 资源搜索
- 支持按分类搜索
- 防抖搜索（500ms）
- 搜索结果分页
- 骨架屏加载动画

### 2. 资源求助
- 工单式提交
- 联系方式验证
- 资源上架通知
- 节流提交（1000ms）

### 3. 资源推荐
- 热搜资源展示
- 分类推荐
- 热门标签

## API接口

### 搜索相关
- `GET /api/search` - 搜索资源
- `GET /api/categories` - 获取分类列表
- `GET /api/hot-resources` - 获取热门资源

### 求助相关
- `POST /api/help-requests` - 提交资源求助
- `GET /api/help-requests/:id` - 获取求助详情

### 统计相关
- `GET /api/stats` - 获取平台统计信息

详细API文档请参考：[API_DOCUMENTATION.md](./API_DOCUMENTATION.md)

## 数据库设计

### 核心表结构

1. **categories** - 资源分类表
2. **resources** - 资源信息表
3. **resource_tags** - 资源标签表
4. **help_requests** - 资源求助表
5. **search_records** - 搜索记录表
6. **download_records** - 下载记录表
7. **users** - 用户表
8. **system_configs** - 系统配置表

详细表结构请参考：[database/README.md](./database/README.md)

## 响应式设计

- **移动端**: < 768px
- **平板端**: 768px - 1024px
- **桌面端**: > 1024px

## 性能优化

### 前端优化
- 组件懒加载
- 图片懒加载
- 防抖搜索
- 节流提交
- 骨架屏加载

### 数据库优化
- 全文搜索索引
- 查询缓存
- 分页查询
- 存储过程优化

## 安全考虑

### 前端安全
- XSS防护
- CSRF防护
- 输入验证
- 输出编码

### 数据库安全
- SQL注入防护
- 参数化查询
- 权限控制
- 数据备份

## 监控和日志

### 前端监控
- 错误监控
- 性能监控
- 用户行为分析

### 数据库监控
- 慢查询日志
- 连接数监控
- 性能指标监控

## 备份策略

### 数据库备份
```bash
# 每日备份
mysqldump -u root pan_search > backup_$(date +%Y%m%d).sql

# 恢复备份
mysql -u root pan_search < backup_20240115.sql
```

### 代码备份
- Git版本控制
- 定期代码归档
- 部署包备份

## 故障排除

### 常见问题

1. **数据库连接失败**
   - 检查MySQL服务是否运行
   - 验证连接信息是否正确
   - 检查防火墙设置

2. **前端构建失败**
   - 清除node_modules重新安装
   - 检查Node.js版本
   - 查看构建日志

3. **API调用失败**
   - 检查网络连接
   - 验证API地址
   - 查看浏览器控制台

### 日志查看

```bash
# 查看MySQL错误日志
tail -f /var/log/mysql/error.log

# 查看应用日志
tail -f /var/log/pan-search/app.log
```

## 扩展开发

### 添加新功能
1. 在`src/views/`创建新页面组件
2. 在`src/router/index.js`添加路由
3. 在`src/services/api.js`添加API调用
4. 更新数据库表结构（如需要）

### 自定义样式
- 修改`src/style.css`全局样式
- 组件内使用scoped样式
- 响应式断点：768px, 1024px

## 联系方式

如有问题请联系：
- 邮箱：admin@pan-search.com
- GitHub：https://github.com/pan-search/web

## 许可证

MIT License