
<template>
  <div class="search-page">
    <div class="container">
      <!-- 搜索框 -->
      <div class="search-header">
        <div class="search-box">
          <div class="search-input-wrapper">
            <input
              v-model="searchKeyword"
              type="text"
              placeholder="输入关键词搜索资源..."
              class="search-input"
              @input="handleSearchInput"
              @keyup.enter="performSearch"
            />
            <button class="search-btn" @click="performSearch" :disabled="loading">
              <span class="search-icon">🔍</span>
              {{ loading ? '搜索中...' : '搜索' }}
            </button>
          </div>
        </div>
        <div class="filter-options">
          <div class="filter-group">
            <label class="filter-label">分类</label>
            <select v-model="selectedCategory" class="filter-select" @change="performSearch">
              <option value="all">全部分类</option>
              <option value="movie">电影</option>
              <option value="tv">电视剧</option>
              <option value="music">音乐</option>
              <option value="software">软件</option>
              <option value="document">文档</option>
              <option value="other">其他</option>
            </select>
          </div>
          <div class="filter-group">
            <label class="filter-label">排序</label>
            <select v-model="sortBy" class="filter-select" @change="performSearch">
              <option value="relevance">相关度</option>
              <option value="time">最新</option>
              <option value="size">文件大小</option>
            </select>
          </div>
        </div>
      </div>

      <!-- 搜索结果 -->
      <div class="search-results">
        <!-- 骨架屏 -->
        <div v-if="loading" class="skeleton-container">
          <div v-for="n in 6" :key="n" class="skeleton-card">
            <div class="skeleton-image"></div>
            <div class="skeleton-content">
              <div class="skeleton-title"></div>
              <div class="skeleton-description"></div>
              <div class="skeleton-meta"></div>
            </div>
          </div>
        </div>

        <!-- 搜索结果列表 -->
        <div v-else-if="searchResults.length > 0" class="results-grid">
          <div 
            v-for="result in searchResults" 
            :key="result.id"
            class="result-card"
          >
            <div class="result-image">
              <div class="image-placeholder" :class="getTypeClass(result.type)">
                {{ getTypeIcon(result.type) }}
              </div>
            </div>
            <div class="result-content">
              <h3 class="result-title">{{ result.title }}</h3>
              <p class="result-description">{{ result.description }}</p>
              <div class="result-meta">
                <span class="file-size">{{ result.size }}</span>
                <span class="file-type">{{ result.type }}</span>
                <span class="upload-time">{{ result.time }}</span>
                <span class="source">{{ result.source }}</span>
              </div>
              <div class="result-actions">
                <button class="action-btn primary" @click="handleDownload(result)">
                  下载
                </button>
                <button class="action-btn secondary" @click="handleCopyLink(result)">
                  复制链接
                </button>
              </div>
            </div>
          </div>
        </div>

        <!-- 空状态 -->
        <div v-else-if="hasSearched" class="empty-state">
          <div class="empty-icon">🔍</div>
          <h3>未找到相关资源</h3>
          <p>尝试使用其他关键词或分类进行搜索</p>
          <button class="btn" @click="$router.push('/request')">资源求助</button>
        </div>

        <!-- 初始状态 -->
        <div v-else class="initial-state">
          <div class="initial-icon">📁</div>
          <h3>开始搜索网盘资源</h3>
          <p>输入关键词搜索您需要的资源</p>
        </div>

        <!-- 分页 -->
        <div v-if="searchResults.length > 0" class="pagination">
          <button 
            class="pagination-btn" 
            :disabled="currentPage === 1"
            @click="goToPage(currentPage - 1)"
          >
            上一页
          </button>
          <span class="page-info">第 {{ currentPage }} 页</span>
          <button 
            class="pagination-btn" 
            :disabled="!hasNextPage"
            @click="goToPage(currentPage + 1)"
          >
            下一页
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { debounce } from '../utils/debounce'

export default {
  name: 'Search',
  data() {
    return {
      searchKeyword: '',
      selectedCategory: 'all',
      sortBy: 'relevance',
      searchResults: [],
      loading: false,
      hasSearched: false,
      currentPage: 1,
      hasNextPage: false,
      debouncedSearch: null
    }
  },
  created() {
    // 从URL参数初始化搜索条件
    const { q, category } = this.$route.query
    if (q) {
      this.searchKeyword = q
      this.selectedCategory = category || 'all'
      this.performSearch()
    }
    
    // 创建防抖搜索函数
    this.debouncedSearch = debounce(this.performSearch, 500)
  },
  methods: {
    handleSearchInput() {
      if (this.searchKeyword.trim()) {
        this.debouncedSearch()
      }
    },
    async performSearch() {
      if (!this.searchKeyword.trim()) {
        this.searchResults = []
        this.hasSearched = false
        return
      }

      this.loading = true
      this.hasSearched = true
      
      try {
        // 模拟API调用
        await new Promise(resolve => setTimeout(resolve, 1000))
        
        // 模拟搜索结果
        this.searchResults = this.generateMockResults()
        this.hasNextPage = this.searchResults.length === 10
      } catch (error) {
        console.error('搜索失败:', error)
      } finally {
        this.loading = false
      }
    },
    generateMockResults() {
      const types = ['电影', '电视剧', '音乐', '软件', '文档', '其他']
      const sources = ['百度网盘', '阿里云盘', '腾讯微云', '天翼云盘']
      const results = []
      
      for (let i = 0; i < 10; i++) {
        const type = types[Math.floor(Math.random() * types.length)]
        const source = sources[Math.floor(Math.random() * sources.length)]
        
        results.push({
          id: Date.now() + i,
          title: `${this.searchKeyword} ${type}资源 ${i + 1}`,
          description: `这是一个关于${this.searchKeyword}的${type}资源，包含完整内容和高清版本。`,
          size: `${(Math.random() * 100 + 1).toFixed(1)}GB`,
          type: type,
          time: `${Math.floor(Math.random() * 30) + 1}天前`,
          source: source,
          downloadUrl: `https://example.com/download/${Date.now() + i}`
        })
      }
      
      return results
    },
    getTypeClass(type) {
      const typeMap = {
        '电影': 'type-movie',
        '电视剧': 'type-tv',
        '音乐': 'type-music',
        '软件': 'type-software',
        '文档': 'type-document',
        '其他': 'type-other'
      }
      return typeMap[type] || 'type-other'
    },
    getTypeIcon(type) {
      const iconMap = {
        '电影': '🎬',
        '电视剧': '📺',
        '音乐': '🎵',
        '软件': '💻',
        '文档': '📄',
        '其他': '📦'
      }
      return iconMap[type] || '📦'
    },
    handleDownload(result) {
      alert(`开始下载: ${result.title}\n来源: ${result.source}`)
      // 实际项目中这里会处理下载逻辑
    },
    handleCopyLink(result) {
      navigator.clipboard.writeText(result.downloadUrl)
        .then(() => alert('链接已复制到剪贴板'))
        .catch(() => alert('复制失败，请手动复制链接'))
    },
    goToPage(page) {
      this.currentPage = page
      this.performSearch()
    }
  }
}
</script>

<style scoped>
.search-page {
  padding: 24px 0 40px;
}

.search-header {
  background: white;
  padding: 24px;
  border-radius: 8px;
  border: 1px solid #e1e5e9;
  margin-bottom: 24px;
}

.search-box {
  margin-bottom: 20px;
}

.search-input-wrapper {
  display: flex;
  gap: 12px;
  max-width: 600px;
}

.search-input {
  flex: 1;
  padding: 12px 16px;
  border: 1px solid #d0d7de;
  border-radius: 8px;
  font-size: 16px;
  transition: all 0.2s ease;
  background: white;
}

.search-input:focus {
  outline: none;
  border-color: #0969da;
  box-shadow: 0 0 0 3px rgba(9, 105, 218, 0.1);
}

.search-btn {
  padding: 12px 24px;
  background: #656d76;
  color: white;
  border: none;
  border-radius: 8px;
  font-size: 14px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.2s ease;
  display: flex;
  align-items: center;
  gap: 8px;
  min-width: 100px;
}

.search-btn:hover:not(:disabled) {
  background: #57606a;
  transform: translateY(-1px);
}

.search-btn:disabled {
  background: #8c959f;
  cursor: not-allowed;
  transform: none;
}

.search-icon {
  font-size: 16px;
}

.filter-options {
  display: flex;
  gap: 20px;
  flex-wrap: wrap;
}

.filter-group {
  display: flex;
  align-items: center;
  gap: 8px;
}

.filter-label {
  font-size: 13px;
  font-weight: 600;
  color: #656d76;
  white-space: nowrap;
}

.filter-select {
  padding: 8px 12px;
  border: 1px solid #d0d7de;
  border-radius: 6px;
  background: white;
  font-size: 13px;
  cursor: pointer;
  min-width: 120px;
}

.filter-select:focus {
  outline: none;
  border-color: #0969da;
}

/* 骨架屏样式 */
.skeleton-container {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 20px;
}

.skeleton-card {
  background: white;
  border: 1px solid #e1e5e9;
  border-radius: 8px;
  padding: 20px;
  display: flex;
  gap: 16px;
}

.skeleton-image {
  width: 60px;
  height: 60px;
  border-radius: 6px;
  background: linear-gradient(90deg, #f6f8fa 25%, #eaeef2 50%, #f6f8fa 75%);
  background-size: 200% 100%;
  animation: loading 1.5s infinite;
  flex-shrink: 0;
}

.skeleton-content {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.skeleton-title {
  height: 16px;
  width: 70%;
  border-radius: 4px;
  background: linear-gradient(90deg, #f6f8fa 25%, #eaeef2 50%, #f6f8fa 75%);
  background-size: 200% 100%;
  animation: loading 1.5s infinite;
}

.skeleton-description {
  height: 14px;
  width: 90%;
  border-radius: 4px;
  background: linear-gradient(90deg, #f6f8fa 25%, #eaeef2 50%, #f6f8fa 75%);
  background-size: 200% 100%;
  animation: loading 1.5s infinite;
}

.skeleton-meta {
  height: 12px;
  width: 60%;
  border-radius: 4px;
  background: linear-gradient(90deg, #f6f8fa 25%, #eaeef2 50%, #f6f8fa 75%);
  background-size: 200% 100%;
  animation: loading 1.5s infinite;
}

/* 搜索结果样式 */
.results-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 20px;
}

.result-card {
  background: white;
  border: 1px solid #e1e5e9;
  border-radius: 8px;
  padding: 20px;
  display: flex;
  gap: 16px;
  transition: all 0.2s ease;
  cursor: pointer;
}

.result-card:hover {
  transform: translateY(-2px);
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.08);
  border-color: #d0d7de;
}

.result-image {
  flex-shrink: 0;
}

.image-placeholder {
  width: 60px;
  height: 60px;
  border-radius: 6px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 20px;
  color: white;
}

.type-movie { background: #e74c3c; }
.type-tv { background: #3498db; }
.type-music { background: #9b59b6; }
.type-software { background: #2ecc71; }
.type-document { background: #f39c12; }
.type-other { background: #95a5a6; }

.result-content {
  flex: 1;
  min-width: 0;
}

.result-title {
  font-size: 16px;
  font-weight: 600;
  margin-bottom: 8px;
  color: #1a1a1a;
  line-height: 1.4;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.result-description {
  color: #666;
  margin-bottom: 12px;
  line-height: 1.5;
  font-size: 13px;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.result-meta {
  display: flex;
  gap: 12px;
  margin-bottom: 16px;
  flex-wrap: wrap;
}

.result-meta span {
  padding: 4px 8px;
  background: #f6f8fa;
  border-radius: 4px;
  font-size: 11px;
  color: #656d76;
  font-weight: 500;
}

.result-actions {
  display: flex;
  gap: 8px;
}

.action-btn {
  padding: 6px 12px;
  border: none;
  border-radius: 6px;
  font-size: 12px;
  cursor: pointer;
  transition: all 0.2s ease;
  font-weight: 500;
}

.action-btn.primary {
  background: #0969da;
  color: white;
}

.action-btn.primary:hover {
  background: #0550ae;
}

.action-btn.secondary {
  background: #656d76;
  color: white;
}

.action-btn.secondary:hover {
  background: #57606a;
}

/* 空状态和初始状态 */
.empty-state,
.initial-state {
  text-align: center;
  padding: 48px 20px;
  background: white;
  border-radius: 8px;
  border: 1px solid #e1e5e9;
}

.empty-icon,
.initial-icon {
  font-size: 48px;
  margin-bottom: 16px;
}

.empty-state h3,
.initial-state h3 {
  font-size: 18px;
  margin-bottom: 8px;
  color: #1a1a1a;
}

.empty-state p,
.initial-state p {
  color: #666;
  margin-bottom: 20px;
  font-size: 14px;
}

/* 分页样式 */
.pagination {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 16px;
  margin-top: 32px;
  padding: 20px;
}

.pagination-btn {
  padding: 8px 16px;
  border: 1px solid #d0d7de;
  background: white;
  border-radius: 6px;
  cursor: pointer;
  transition: all 0.2s ease;
  font-size: 13px;
  font-weight: 500;
}

.pagination-btn:hover:not(:disabled) {
  background: #0969da;
  color: white;
  border-color: #0969da;
}

.pagination-btn:disabled {
  background: #f6f8fa;
  color: #8c959f;
  cursor: not-allowed;
}

.page-info {
  color: #656d76;
  font-size: 13px;
}

/* 移动端适配 */
@media (max-width: 768px) {
  .search-page {
    padding: 16px 0 32px;
  }
  
 
  .search-header {
    padding: 20px;
    margin-bottom: 20px;
  }
  
  .search-input-wrapper {
    flex-direction: column;
    gap: 12px;
  }
  
  .search-btn {
    justify-content: center;
  }
  
  .filter-options {
    flex-direction: column;
    gap: 12px;
  }
  
  .filter-group {
    justify-content: space-between;
  }
  
  .filter-select {
    min-width: 140px;
  }
  
  .skeleton-container,
  .results-grid {
    grid-template-columns: 1fr;
    gap: 16px;
  }
  
  .skeleton-card,
  .result-card {
    padding: 16px;
  }
  
  .pagination {
    flex-direction: column;
    gap: 12px;
  }
}

@media (max-width: 480px) {
  .search-header {
    padding: 16px;
  }
  
  .search-input {
    font-size: 14px;
  }
  
  .result-title {
    font-size: 15px;
  }
  
  .result-description {
    font-size: 12px;
  }
  
  .result-meta {
    gap: 8px;
  }
  
  .result-meta span {
    font-size: 10px;
  }
}

@keyframes loading {
  0% {
    background-position: 200% 0;
  }
  100% {
    background-position: -200% 0;
  }
}
</style>