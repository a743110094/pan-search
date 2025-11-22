# 网盘资源搜索平台 - 前端应用

基于 Vue 3 + Vite 开发的现代化网盘资源搜索平台前端应用，提供流畅的搜索体验和响应式界面。

## ✨ 功能特性

- 🔍 **智能搜索**: 实时搜索、搜索建议、历史记录
- 📊 **热门推荐**: 热门资源展示，一键查看
- 📂 **分类浏览**: 按分类筛选资源
- 📝 **求助中心**: 提交资源求助请求
- 🎨 **现代化UI**: 响应式设计，支持移动端
- ⚡ **高性能**: Vite 构建，快速热更新
- 🔄 **路由懒加载**: 按需加载，提升性能
- 💾 **本地存储**: 搜索历史本地缓存

## 🛠️ 技术栈

- **框架**: [Vue 3](https://vuejs.org/)
- **构建工具**: [Vite](https://vitejs.dev/)
- **路由**: [Vue Router](https://router.vuejs.org/)
- **HTTP 客户端**: Fetch API
- **样式**: 原生 CSS

## 📁 项目结构

```
frontend/
├── index.html              # 主页面入口
├── package.json            # 项目依赖和脚本
├── vite.config.js          # Vite 配置
├── yarn.lock               # 依赖版本锁定
└── src/                    # 源代码目录
    ├── main.js             # 应用入口
    ├── App.vue             # 根组件
    ├── style.css           # 全局样式
    ├── components/         # 公共组件
    │   ├── Header.vue      # 头部导航
    │   └── Footer.vue      # 页脚
    ├── views/              # 页面视图
    │   ├── Home.vue        # 首页
    │   ├── Search.vue      # 搜索页
    │   └── Request.vue     # 求助页
    ├── router/             # 路由配置
    │   └── index.js        # 路由定义
    ├── services/           # API 服务
    │   └── api.js          # API 封装
    └── utils/              # 工具函数
        └── debounce.js     # 防抖函数
```

## 🚀 快速开始

### 前置要求

- Node.js 16+
- Yarn 或 npm

### 安装依赖

```bash
yarn install
# 或
npm install
```

### 启动开发服务器

```bash
yarn dev
# 或
npm run dev
```

开发服务器将在 http://localhost:3000 启动

### 构建生产版本

```bash
yarn build
# 或
npm run build
```

构建文件将输出到 `dist/` 目录

### 预览生产版本

```bash
yarn preview
# 或
npm run preview
```

## 🎯 使用指南

### 页面说明

#### 首页 (Home)
- 展示平台介绍
- 热门资源推荐
- 快速搜索入口

#### 搜索页 (Search)
- 资源搜索功能
- 分类筛选
- 排序选项
- 分页浏览

#### 求助页 (Request)
- 提交资源求助
- 填写求助信息
- 查看求助列表

### API 服务

API 调用封装在 `src/services/api.js` 中：

```javascript
// 搜索资源
api.searchResources(query, category, sort, page)

// 获取热门推荐
api.getHotResources()

// 获取分类列表
api.getCategories()

// 提交求助请求
api.submitRequest(requestData)

// 获取搜索建议
api.getSuggestions(query)
```

### 路由配置

路由定义在 `src/router/index.js`：

```javascript
const routes = [
  {
    path: '/',
    name: 'Home',
    component: () => import('../views/Home.vue')
  },
  {
    path: '/search',
    name: 'Search',
    component: () => import('../views/Search.vue')
  },
  {
    path: '/request',
    name: 'Request',
    component: () => import('../views/Request.vue')
  }
]
```

## 🎨 自定义样式

### 全局样式

在 `src/style.css` 中定义全局样式：

```css
:root {
  --primary-color: #409eff;
  --success-color: #67c23a;
  --warning-color: #e6a23c;
  --danger-color: #f56c6c;
  --text-color: #303133;
  --bg-color: #ffffff;
}
```

### 组件样式

每个 `.vue` 文件的 `<style>` 部分：

```vue
<style scoped>
/* 组件样式，scoped 表示仅影响当前组件 */
.search-box {
  /* ... */
}
</style>
```

## 🛠️ 开发指南

### 创建新页面

1. 在 `src/views/` 中创建新的 `.vue` 文件
2. 在 `src/router/index.js` 中添加路由
3. 在需要的地方添加导航链接

示例：
```vue
<template>
  <div class="new-page">
    <h1>新页面</h1>
  </div>
</template>

<script>
export default {
  name: 'NewPage'
}
</script>

<style scoped>
.new-page {
  /* 样式 */
}
</style>
```

### 创建新组件

1. 在 `src/components/` 中创建组件文件
2. 在需要的页面中引入并使用

示例：
```vue
<template>
  <div class="custom-component">
    {{ message }}
  </div>
</template>

<script>
export default {
  name: 'CustomComponent',
  props: {
    message: String
  }
}
</script>
```

### API 调用

1. 在 `src/services/api.js` 中添加新的 API 方法
2. 在组件中引入并使用

```javascript
// services/api.js
export const newApi = async (params) => {
  const response = await fetch('/api/new', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(params)
  })
  return response.json()
}
```

```vue
<!-- 在组件中使用 -->
<script>
import { newApi } from '../services/api.js'

export default {
  async mounted() {
    const data = await newApi(params)
    // 处理数据
  }
}
</script>
```

## ⚡ 性能优化

### 代码分割

使用动态导入实现路由懒加载：

```javascript
{
  path: '/search',
  name: 'Search',
  component: () => import('../views/Search.vue')
}
```

### 资源优化

- Vite 自动进行代码分割
- 图片懒加载
- CSS 压缩和优化

## 🔧 Vite 配置

`vite.config.js` 包含以下配置：

```javascript
import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

export default defineConfig({
  plugins: [vue()],
  server: {
    port: 3000,
    open: true
  }
})
```

## 📱 响应式设计

项目采用响应式设计，支持多种屏幕尺寸：

- 桌面端 (>1024px)
- 平板端 (768px - 1024px)
- 移动端 (<768px)

使用 CSS Media Queries 实现：

```css
@media (max-width: 768px) {
  .mobile-hidden {
    display: none;
  }
}
```

## 🌐 浏览器兼容性

支持以下现代浏览器：

- Chrome >= 87
- Firefox >= 78
- Safari >= 14
- Edge >= 88

## 📦 依赖说明

### 生产依赖

- `vue`: Vue 3 框架
- `vue-router`: Vue 路由

### 开发依赖

- `vite`: 构建工具
- `@vitejs/plugin-vue`: Vite 的 Vue 插件

## 🐛 调试

### 查看详细日志

```bash
yarn dev --debug
```

### 网络请求调试

打开浏览器开发者工具，在 Network 面板查看 API 请求

### Vue DevTools

安装 Vue DevTools 浏览器扩展，查看组件状态和事件

## 📄 许可证

MIT License

## 👥 贡献

欢迎提交 Issue 和 Pull Request！

## 🙏 致谢

- [Vue.js](https://vuejs.org/) - 渐进式 JavaScript 框架
- [Vite](https://vitejs.dev/) - 下一代前端构建工具
- [Vue Router](https://router.vuejs.org/) - Vue.js 官方路由
