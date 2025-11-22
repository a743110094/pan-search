import axios from 'axios'

// 创建axios实例
const apiClient = axios.create({
  baseURL: process.env.VITE_API_BASE_URL || 'https://api.pansearch.com/v1',
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json'
  }
})

// 请求拦截器
apiClient.interceptors.request.use(
  (config) => {
    // 添加认证token
    const token = localStorage.getItem('auth_token')
    if (token) {
      config.headers.Authorization = `Bearer ${token}`
    }
    return config
  },
  (error) => {
    return Promise.reject(error)
  }
)

// 响应拦截器
apiClient.interceptors.response.use(
  (response) => {
    return response.data
  },
  (error) => {
    console.error('API请求错误:', error)
    
    // 统一错误处理
    if (error.response) {
      const { status, data } = error.response
      switch (status) {
        case 401:
          // 未授权，跳转到登录页
          localStorage.removeItem('auth_token')
          window.location.href = '/login'
          break
        case 403:
          console.error('权限不足:', data.message)
          break
        case 404:
          console.error('资源不存在:', data.message)
          break
        case 500:
          console.error('服务器错误:', data.message)
          break
        default:
          console.error('请求错误:', data.message)
      }
    } else if (error.request) {
      console.error('网络错误: 请检查网络连接')
    } else {
      console.error('请求配置错误:', error.message)
    }
    
    return Promise.reject(error)
  }
)

// API服务类
class ApiService {
  // 搜索资源
  async searchResources(params) {
    return apiClient.get('/resources/search', { params })
  }

  // 获取热门推荐
  async getHotResources(limit = 10) {
    return apiClient.get('/resources/hot', { params: { limit } })
  }

  // 提交资源求助
  async submitHelpRequest(data) {
    return apiClient.post('/requests', data)
  }

  // 获取分类列表
  async getCategories() {
    return apiClient.get('/categories')
  }

  // 记录下载行为
  async recordDownload(resourceId, data = {}) {
    return apiClient.post(`/resources/${resourceId}/download`, data)
  }

  // 获取搜索建议
  async getSearchSuggestions(keyword) {
    return apiClient.get('/search/suggestions', { params: { q: keyword } })
  }

  // 获取求助请求状态
  async getRequestStatus(requestId) {
    return apiClient.get(`/requests/${requestId}`)
  }

  // 用户登录
  async login(credentials) {
    return apiClient.post('/auth/login', credentials)
  }

  // 用户注册
  async register(userData) {
    return apiClient.post('/auth/register', userData)
  }

  // 获取用户信息
  async getUserProfile() {
    return apiClient.get('/user/profile')
  }

  // 更新用户信息
  async updateUserProfile(userData) {
    return apiClient.put('/user/profile', userData)
  }
}

// 创建服务实例
const apiService = new ApiService()

export default apiService

// 工具函数
export const formatFileSize = (bytes) => {
  if (bytes === 0) return '0 B'
  const k = 1024
  const sizes = ['B', 'KB', 'MB', 'GB', 'TB']
  const i = Math.floor(Math.log(bytes) / Math.log(k))
  return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i]
}

export const formatTime = (timestamp) => {
  const now = new Date()
  const time = new Date(timestamp)
  const diff = now - time
  
  const minute = 60 * 1000
  const hour = minute * 60
  const day = hour * 24
  const month = day * 30
  const year = day * 365
  
  if (diff < minute) {
    return '刚刚'
  } else if (diff < hour) {
    return Math.floor(diff / minute) + '分钟前'
  } else if (diff < day) {
    return Math.floor(diff / hour) + '小时前'
  } else if (diff < month) {
    return Math.floor(diff / day) + '天前'
  } else if (diff < year) {
    return Math.floor(diff / month) + '个月前'
  } else {
    return Math.floor(diff / year) + '年前'
  }
}

export const validateEmail = (email) => {
  const regex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
  return regex.test(email)
}

export const validatePhone = (phone) => {
  const regex = /^1[3-9]\d{9}$/
  return regex.test(phone)
}