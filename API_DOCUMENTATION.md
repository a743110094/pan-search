# 网盘资源搜索平台 API 文档

本文档描述了网盘资源搜索平台的前后端接口规范，基于当前前端Mock实现。

## 基础信息

- **基础URL**: `https://api.pansearch.com/v1`
- **认证方式**: Bearer Token (可选，用于用户相关操作)
- **数据格式**: JSON
- **字符编码**: UTF-8

## 通用响应格式

### 成功响应
```json
{
  "code": 200,
  "message": "success",
  "data": {},
  "timestamp": 1630000000000
}
```

### 错误响应
```json
{
  "code": 400,
  "message": "错误描述",
  "data": null,
  "timestamp": 1630000000000
}
```

### 分页响应
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [],
    "pagination": {
      "page": 1,
      "pageSize": 10,
      "total": 100,
      "totalPages": 10
    }
  },
  "timestamp": 1630000000000
}
```

## 状态码说明

| 状态码 | 说明 |
|--------|------|
| 200 | 成功 |
| 400 | 请求参数错误 |
| 401 | 未授权 |
| 403 | 禁止访问 |
| 404 | 资源不存在 |
| 500 | 服务器内部错误 |

## 接口列表

### 1. 搜索资源

**接口**: `GET /resources/search`

**描述**: 根据关键词和分类搜索网盘资源

**请求参数**:
| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| q | string | 是 | 搜索关键词 |
| category | string | 否 | 资源分类 (all/movie/tv/music/software/document/other) |
| sort | string | 否 | 排序方式 (relevance/time/size) |
| page | integer | 否 | 页码，默认1 |
| pageSize | integer | 否 | 每页数量，默认10 |

**请求示例**:
```bash
GET /resources/search?q=电影&category=movie&sort=relevance&page=1&pageSize=10
```

**响应数据**:
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": "resource_001",
        "title": "2024最新电影合集",
        "description": "包含2024年最新上映的国内外热门电影",
        "size": "15.2GB",
        "type": "movie",
        "category": "电影",
        "source": "百度网盘",
        "downloadUrl": "https://pan.baidu.com/s/xxxxx",
        "extractCode": "abcd",
        "uploadTime": "2024-01-15T10:30:00Z",
        "fileCount": 25,
        "viewCount": 12500,
        "downloadCount": 8900,
        "tags": ["2024", "高清", "合集"],
        "valid": true,
        "expireTime": "2024-12-31T23:59:59Z"
      }
    ],
    "pagination": {
      "page": 1,
      "pageSize": 10,
      "total": 156,
      "totalPages": 16
    }
  },
  "timestamp": 1630000000000
}
```

### 2. 获取热门推荐

**接口**: `GET /resources/hot`

**描述**: 获取热门资源推荐列表

**请求参数**:
| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| limit | integer | 否 | 返回数量，默认10 |

**请求示例**:
```bash
GET /resources/hot?limit=10
```

**响应数据**:
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "id": "resource_001",
        "rank": 1,
        "title": "2024最新电影合集",
        "description": "包含2024年最新上映的国内外热门电影",
        "size": "15.2GB",
        "type": "movie",
        "category": "电影",
        "searchCount": 12500,
        "trend": "up", // up/down/stable
        "trendValue": 15 // 百分比变化
      }
    ]
  },
  "timestamp": 1630000000000
}
```

### 3. 提交资源求助

**接口**: `POST /requests`

**描述**: 提交资源求助请求

**请求头**:
```
Content-Type: application/json
```

**请求体**:
```json
{
  "resourceName": "需要查找的资源名称",
  "resourceType": "movie",
  "description": "详细描述资源需求",
  "contact": "user@example.com",
  "contactType": "email" // email/phone
}
```

**响应数据**:
```json
{
  "code": 200,
  "message": "求助提交成功",
  "data": {
    "requestId": "req_001",
    "status": "pending",
    "submitTime": "2024-01-15T10:30:00Z",
    "estimatedProcessTime": "1-3个工作日"
  },
  "timestamp": 1630000000000
}
```

### 4. 获取分类列表

**接口**: `GET /categories`

**描述**: 获取所有资源分类

**请求示例**:
```bash
GET /categories
```

**响应数据**:
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "list": [
      {
        "value": "all",
        "label": "全部",
        "count": 15600,
        "icon": "📁"
      },
      {
        "value": "movie",
        "label": "电影",
        "count": 5200,
        "icon": "🎬"
      },
      {
        "value": "tv",
        "label": "电视剧",
        "count": 3800,
        "icon": "📺"
      },
      {
        "value": "music",
        "label": "音乐",
        "count": 2100,
        "icon": "🎵"
      },
      {
        "value": "software",
        "label": "软件",
        "count": 1800,
        "icon": "💻"
      },
      {
        "value": "document",
        "label": "文档",
        "count": 1500,
        "icon": "📄"
      },
      {
        "value": "other",
        "label": "其他",
        "count": 1200,
        "icon": "📦"
      }
    ]
  },
  "timestamp": 1630000000000
}
```

### 5. 资源下载统计

**接口**: `POST /resources/{id}/download`

**描述**: 记录资源下载行为

**请求头**:
```
Content-Type: application/json
Authorization: Bearer {token}
```

**请求体**:
```json
{
  "userId": "user_001", // 可选，如果用户已登录
  "userAgent": "Mozilla/5.0...",
  "ip": "192.168.1.1"
}
```

**响应数据**:
```json
{
  "code": 200,
  "message": "下载记录成功",
  "data": {
    "downloadId": "dl_001",
    "resourceId": "resource_001",
    "downloadUrl": "https://pan.baidu.com/s/xxxxx",
    "expireTime": "2024-01-15T11:30:00Z"
  },
  "timestamp": 1630000000000
}
```

### 6. 获取搜索建议

**接口**: `GET /search/suggestions`

**描述**: 获取搜索关键词建议

**请求参数**:
| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| q | string | 是 | 搜索关键词 |

**请求示例**:
```bash
GET /search/suggestions?q=电影
```

**响应数据**:
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "suggestions": [
      "电影合集",
      "电影资源",
      "电影下载",
      "最新电影",
      "高清电影"
    ]
  },
  "timestamp": 1630000000000
}
```

## 数据模型

### Resource (资源)
```typescript
interface Resource {
  id: string;                    // 资源ID
  title: string;                 // 资源标题
  description: string;           // 资源描述
  size: string;                  // 文件大小
  type: string;                  // 资源类型
  category: string;              // 分类名称
  source: string;                // 来源平台
  downloadUrl: string;           // 下载链接
  extractCode?: string;          // 提取码
  uploadTime: string;            // 上传时间
  fileCount: number;             // 文件数量
  viewCount: number;             // 查看次数
  downloadCount: number;         // 下载次数
  tags: string[];                // 标签
  valid: boolean;                // 是否有效
  expireTime?: string;           // 过期时间
}
```

### SearchRequest (搜索请求)
```typescript
interface SearchRequest {
  q: string;                     // 搜索关键词
  category?: string;             // 分类
  sort?: string;                 // 排序方式
  page?: number;                 // 页码
  pageSize?: number;             // 每页数量
}
```

### HelpRequest (求助请求)
```typescript
interface HelpRequest {
  resourceName: string;          // 资源名称
  resourceType?: string;         // 资源类型
  description?: string;          // 详细描述
  contact: string;               // 联系方式
  contactType: string;           // 联系类型
}
```

### Category (分类)
```typescript
interface Category {
  value: string;                 // 分类值
  label: string;                 // 分类标签
  count: number;                 // 资源数量
  icon: string;                  // 图标
}
```

## 前端Mock数据映射

当前前端Mock数据与后端API的对应关系：

### 首页热门推荐
- **Mock数据**: `src/views/Home.vue` 中的 `hotResources`
- **对应API**: `GET /resources/hot`

### 搜索功能
- **Mock数据**: `src/views/Search.vue` 中的 `generateMockResults()`
- **对应API**: `GET /resources/search`

### 资源求助
- **Mock数据**: `src/views/Request.vue` 中的表单提交
- **对应API**: `POST /requests`

### 分类数据
- **Mock数据**: 各页面中的 `categories` 数组
- **对应API**: `GET /categories`

## 前端接口调用示例

### 搜索资源
```javascript
// 使用 axios 调用搜索接口
const searchResources = async (params) => {
  try {
    const response = await axios.get('/resources/search', { params });
    return response.data;
  } catch (error) {
    console.error('搜索失败:', error);
    throw error;
  }
};

// 调用示例
const results = await searchResources({
  q: '电影',
  category: 'movie',
  sort: 'relevance',
  page: 1,
  pageSize: 10
});
```

### 提交求助
```javascript
const submitHelpRequest = async (data) => {
  try {
    const response = await axios.post('/requests', data);
    return response.data;
  } catch (error) {
    console.error('提交失败:', error);
    throw error;
  }
};

// 调用示例
const result = await submitHelpRequest({
  resourceName: '需要查找的资源',
  resourceType: 'movie',
  description: '详细描述',
  contact: 'user@example.com',
  contactType: 'email'
});
```

## 注意事项

1. **防抖处理**: 搜索接口建议在前端实现防抖，减少不必要的请求
2. **分页处理**: 所有列表接口都支持分页，前端需要处理分页逻辑
3. **错误处理**: 前端需要处理各种HTTP状态码和错误信息
4. **数据验证**: 提交数据前进行前端验证，提高用户体验
5. **加载状态**: 异步请求时显示加载状态，避免用户困惑

## 版本历史

| 版本 | 日期 | 说明 |
|------|------|------|
| v1.0 | 2024-01-15 | 初始版本，基于前端Mock实现 |