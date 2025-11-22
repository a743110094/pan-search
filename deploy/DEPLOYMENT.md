# 网盘资源搜索平台部署文档

本文档描述了如何部署网盘资源搜索平台的前端项目。

## 环境要求

### 开发环境
- Node.js >= 16.0.0
- npm >= 8.0.0 或 yarn >= 1.22.0
- 现代浏览器（Chrome >= 88, Firefox >= 78, Safari >= 14, Edge >= 88）

### 生产环境
- Node.js >= 16.0.0
- Web服务器（Nginx, Apache等）
- CDN（可选，用于静态资源加速）

## 项目结构

```
pan-search-web/
├── public/                 # 静态资源
├── src/                   # 源代码
│   ├── components/        # Vue组件
│   ├── views/            # 页面组件
│   ├── router/           # 路由配置
│   ├── services/         # API服务
│   ├── utils/            # 工具函数
│   └── assets/           # 资源文件
├── dist/                 # 构建输出目录
├── .env.example          # 环境变量示例
├── vite.config.js        # Vite配置
└── package.json          # 项目配置
```

## 部署步骤

### 1. 开发环境部署

#### 1.1 克隆项目
```bash
git clone <repository-url>
cd pan-search-web
```

#### 1.2 安装依赖
```bash
npm install
# 或使用 yarn
yarn install
```

#### 1.3 配置环境变量
```bash
cp .env.example .env
# 编辑 .env 文件，配置相应的环境变量
```

#### 1.4 启动开发服务器
```bash
npm run dev
# 或使用 yarn
yarn dev
```

开发服务器将在 http://localhost:3000 启动，支持热重载。

### 2. 生产环境部署

#### 2.1 构建项目
```bash
npm run build
# 或使用 yarn
yarn build
```

构建完成后，所有静态文件将生成在 `dist/` 目录中。

#### 2.2 环境变量配置
创建生产环境配置文件 `.env.production`：
```bash
VITE_API_BASE_URL=https://api.pansearch.com/v1
VITE_APP_TITLE=网盘资源搜索
VITE_APP_DESCRIPTION=专业的网盘资源搜索平台
VITE_DEBUG=false
VITE_USE_MOCK=false
```

#### 2.3 构建优化版本
```bash
npm run build
```

### 3. 静态文件部署

#### 3.1 Nginx配置
```nginx
server {
    listen 80;
    server_name pansearch.com;
    
    root /var/www/pan-search-web/dist;
    index index.html;
    
    # Gzip压缩
    gzip on;
    gzip_types text/plain text/css application/json application/javascript text/xml application/xml application/xml+rss text/javascript;
    
    # 缓存设置
    location ~* \.(js|css|png|jpg|jpeg|gif|ico|svg)$ {
        expires 1y;
        add_header Cache-Control "public, immutable";
    }
    
    # SPA路由支持
    location / {
        try_files $uri $uri/ /index.html;
    }
    
    # API代理（如果需要）
    location /api/ {
        proxy_pass https://api.pansearch.com/v1/;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
    }
}
```

#### 3.2 Apache配置
```apache
<VirtualHost *:80>
    ServerName pansearch.com
    DocumentRoot /var/www/pan-search-web/dist
    
    # 启用重写模块
    RewriteEngine On
    
    # SPA路由支持
    RewriteCond %{REQUEST_FILENAME} !-f
    RewriteCond %{REQUEST_FILENAME} !-d
    RewriteRule . /index.html [L]
    
    # 缓存设置
    <FilesMatch "\.(js|css|png|jpg|jpeg|gif|ico|svg)$">
        ExpiresActive On
        ExpiresDefault "access plus 1 year"
        Header append Cache-Control "public, immutable"
    </FilesMatch>
</VirtualHost>
```

### 4. Docker部署

#### 4.1 创建Dockerfile
```dockerfile
# 使用官方Node.js镜像
FROM node:16-alpine as builder

# 设置工作目录
WORKDIR /app

# 复制package文件
COPY package*.json ./

# 安装依赖
RUN npm ci --only=production

# 复制源代码
COPY . .

# 构建应用
RUN npm run build

# 使用Nginx作为生产服务器
FROM nginx:alpine

# 复制构建结果到Nginx目录
COPY --from=builder /app/dist /usr/share/nginx/html

# 复制Nginx配置
COPY nginx.conf /etc/nginx/conf.d/default.conf

# 暴露端口
EXPOSE 80

# 启动Nginx
CMD ["nginx", "-g", "daemon off;"]
```

#### 4.2 创建docker-compose.yml
```yaml
version: '3.8'

services:
  pan-search-web:
    build: .
    ports:
      - "80:80"
    environment:
      - VITE_API_BASE_URL=https://api.pansearch.com/v1
    restart: unless-stopped
```

#### 4.3 构建和运行
```bash
# 构建镜像
docker-compose build

# 启动服务
docker-compose up -d

# 查看日志
docker-compose logs -f
```

### 5. CDN部署

#### 5.1 上传到CDN
将 `dist/` 目录中的文件上传到CDN提供商：
- 阿里云OSS
- 腾讯云COS
- AWS S3
- 其他CDN服务

#### 5.2 配置CDN缓存策略
- HTML文件：不缓存或短时间缓存
- JS/CSS文件：长期缓存（使用文件哈希）
- 图片资源：长期缓存

### 6. 持续集成/持续部署（CI/CD）

#### 6.1 GitHub Actions配置
创建 `.github/workflows/deploy.yml`：
```yaml
name: Deploy to Production

on:
  push:
    branches: [ main ]

jobs:
  deploy:
    runs-on: ubuntu-latest
    
    steps:
    - uses: actions/checkout@v2
    
    - name: Setup Node.js
      uses: actions/setup-node@v2
      with:
        node-version: '16'
        cache: 'npm'
    
    - name: Install dependencies
      run: npm ci
    
    - name: Build project
      run: npm run build
      env:
        VITE_API_BASE_URL: ${{ secrets.API_BASE_URL }}
    
    - name: Deploy to server
      uses: appleboy/scp-action@master
      with:
        host: ${{ secrets.SERVER_HOST }}
        username: ${{ secrets.SERVER_USERNAME }}
        key: ${{ secrets.SERVER_SSH_KEY }}
        source: "dist/*"
        target: "/var/www/pan-search-web"
    
    - name: Restart Nginx
      uses: appleboy/ssh-action@master
      with:
        host: ${{ secrets.SERVER_HOST }}
        username: ${{ secrets.SERVER_USERNAME }}
        key: ${{ secrets.SERVER_SSH_KEY }}
        script: |
          sudo systemctl reload nginx
```

## 性能优化建议

### 1. 代码分割
- 使用Vue Router的懒加载
- 按页面拆分代码块

### 2. 资源优化
- 压缩图片资源
- 使用WebP格式图片
- 启用Gzip/Brotli压缩

### 3. 缓存策略
- 静态资源长期缓存
- API响应合理缓存
- Service Worker缓存（PWA）

### 4. 监控和分析
- 添加Google Analytics
- 使用Sentry错误监控
- 性能指标监控（LCP, FID, CLS）

## 安全考虑

### 1. 环境安全
- 保护环境变量
- 使用HTTPS
- 设置CSP头

### 2. 应用安全
- 输入验证和清理
- XSS防护
- CSRF防护

### 3. 服务器安全
- 定期更新依赖
- 使用防火墙
- 监控异常访问

## 故障排除

### 常见问题

1. **构建失败**
   - 检查Node.js版本
   - 清理node_modules重新安装
   - 检查环境变量配置

2. **路由404**
   - 检查服务器SPA配置
   - 验证Nginx/Apache重写规则

3. **API请求失败**
   - 检查API基础URL配置
   - 验证CORS设置
   - 检查网络连接

4. **性能问题**
   - 启用Gzip压缩
   - 优化图片资源
   - 使用CDN加速

## 维护和更新

### 定期任务
- 更新依赖包
- 备份配置文件
- 监控性能指标
- 检查安全漏洞

### 版本管理
- 使用语义化版本
- 维护CHANGELOG
- 测试新版本功能

## 支持联系方式

- 技术问题：dev@pansearch.com
- 部署问题：ops@pansearch.com
- 紧急支持：+86-400-123-4567

---

**注意**: 部署前请确保已获得相应的环境访问权限和配置信息。