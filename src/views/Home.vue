
<template>
  <div class="home">
    <div class="container">
      <!-- 搜索区域 -->
      <section class="search-section">
        <div class="search-container">
          <h1 class="search-title">网盘资源搜索</h1>
          <p class="search-subtitle">快速找到您需要的网盘资源</p>
          <div class="search-box">
            <div class="search-input-wrapper">
              <input
                v-model="searchKeyword"
                type="text"
                placeholder="输入关键词搜索资源..."
                class="search-input"
                @keyup.enter="handleSearch"
              />
              <button class="search-btn" @click="handleSearch">
                <span class="search-icon">🔍</span>
                搜索
              </button>
            </div>
          </div>
          <div class="category-tags">
            <span 
              v-for="category in categories" 
              :key="category.value"
              :class="['category-tag', { active: selectedCategory === category.value }]"
              @click="selectCategory(category.value)"
            >
              {{ category.label }}
            </span>
          </div>
        </div>
      </section>

      <!-- 热门推荐 -->
      <section class="recommend-section">
        <div class="section-header">
          <h2 class="section-title">热门资源推荐</h2>
          <p class="section-subtitle">大家都在搜索的热门资源</p>
        </div>
        <div class="recommend-grid">
          <div 
            v-for="item in hotResources" 
            :key="item.id"
            class="recommend-card"
            @click="handleResourceClick(item)"
          >
            <div class="card-header">
              <div class="rank-badge" :class="getRankClass(item.rank)">
                <span class="rank-number">{{ item.rank }}</span>
              </div>
              <div class="resource-type">{{ item.type }}</div>
            </div>
            <h3 class="card-title">{{ item.title }}</h3>
            <p class="card-description">{{ item.description }}</p>
            <div class="card-footer">
              <span class="file-size">{{ item.size }}</span>
              <span class="search-count">{{ formatSearchCount(item.searchCount) }}次搜索</span>
            </div>
          </div>
        </div>
      </section>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Home',
  data() {
    return {
      searchKeyword: '',
      selectedCategory: 'all',
      categories: [
        { label: '全部', value: 'all' },
        { label: '电影', value: 'movie' },
        { label: '电视剧', value: 'tv' },
        { label: '音乐', value: 'music' },
        { label: '软件', value: 'software' },
        { label: '文档', value: 'document' },
        { label: '其他', value: 'other' }
      ],
      hotResources: [
        {
          id: 1,
          rank: 1,
          title: '2024最新电影合集',
          description: '包含2024年最新上映的国内外热门电影',
          size: '15.2GB',
          type: '电影',
          searchCount: 12500
        },
        {
          id: 2,
          rank: 2,
          title: '经典电视剧全集',
          description: '经典国产电视剧高清全集资源',
          size: '89.5GB',
          type: '电视剧',
          searchCount: 9800
        },
        {
          id: 3,
          rank: 3,
          title: '无损音乐合集',
          description: '高品质无损音乐资源包',
          size: '32.1GB',
          type: '音乐',
          searchCount: 7600
        },
        {
          id: 4,
          rank: 4,
          title: '办公软件套装',
          description: '常用办公软件及工具包',
          size: '8.7GB',
          type: '软件',
          searchCount: 6500
        },
        {
          id: 5,
          rank: 5,
          title: '学习资料大全',
          description: '各类学习教程和电子书籍',
          size: '45.3GB',
          type: '文档',
          searchCount: 5400
        },
        {
          id: 6,
          rank: 6,
          title: '游戏资源合集',
          description: '热门游戏安装包和补丁',
          size: '120.8GB',
          type: '其他',
          searchCount: 4300
        },
        {
          id: 7,
          title: '编程教程合集',
          description: '各类编程语言学习资料',
          size: '28.4GB',
          type: '文档',
          searchCount: 3200
        },
        {
          id: 9,
          rank: 9,
          title: '动漫资源全集',
          description: '热门动漫高清资源',
          size: '156.2GB',
          type: '其他',
          searchCount: 2900
        },
        {
          id: 10,
          rank: 10,
          title: '摄影教程合集',
          description: '专业摄影技巧和后期教程',
          size: '22.7GB',
          type: '文档',
          searchCount: 2500
        }
      ]
    }
  },
  methods: {
    handleSearch() {
      if (this.searchKeyword.trim()) {
        this.$router.push({
          path: '/search',
          query: {
            q: this.searchKeyword,
            category: this.selectedCategory
          }
        })
      }
    },
    selectCategory(category) {
      this.selectedCategory = category
    },
    handleResourceClick(item) {
      this.searchKeyword = item.title
      this.handleSearch()
    },
    getRankClass(rank) {
      if (rank <= 3) return 'rank-top'
      if (rank <= 6) return 'rank-middle'
      return 'rank-bottom'
    }
  }
}
</script>

<style scoped>
.home {
  padding: 40px 0;
}

.search-section {
  background: #f8f9fa;
  color: #333;
  padding: 60px 0;
  border-radius: 12px;
  margin-bottom: 40px;
  text-align: center;
  border: 1px solid #eee;
}

.search-container {
  max-width: 600px;
  margin: 0 auto;
}

.search-title {
  font-size: 36px;
  font-weight: 700;
  margin-bottom: 10px;
}

.search-subtitle {
  font-size: 18px;
  opacity: 0.9;
  margin-bottom: 30px;
}

.search-box {
  display: flex;
  gap: 10px;
  margin-bottom: 20px;
}

.search-input {
  flex: 1;
  padding: 15px 20px;
  border: none;
  border-radius: 8px;
  font-size: 16px;
  outline: none;
}

.search-btn {
  padding: 15px 30px;
  background: #007bff;
  color: white;
  border: none;
  border-radius: 8px;
  font-size: 16px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
}

.search-btn:hover {
  background: #0056b3;
  transform: translateY(-2px);
}

.category-tags {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
  justify-content: center;
}

.category-tag {
  padding: 8px 16px;
  background: #ffffff;
  border: 1px solid #ddd;
  border-radius: 20px;
  cursor: pointer;
  transition: all 0.3s ease;
  font-size: 14px;
  color: #666;
}

.category-tag:hover {
  background: #f8f9fa;
  border-color: #ccc;
}

.category-tag.active {
  background: #007bff;
  color: white;
  border-color: #007bff;
  font-weight: 600;
}

.recommend-section {
  margin-top: 40px;
}

.section-title {
  font-size: 28px;
  font-weight: 600;
  margin-bottom: 30px;
  text-align: center;
  color: #2c3e50;
}

.recommend-list {
  display: grid;
  gap: 16px;
}

.recommend-item {
  display: flex;
  align-items: center;
  gap: 16px;
  padding: 20px;
  background: white;
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  cursor: pointer;
  transition: all 0.3s ease;
}

.recommend-item:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 16px rgba(0, 0, 0, 0.15);
}

.rank {
  width: 40px;
  height: 40px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 8px;
  font-weight: 700;
  font-size: 16px;
  color: white;
}

.rank-top {
  background: #ff6b6b;
}

.rank-middle {
  background: #4ecdc4;
}

.rank-bottom {
  background: #45b7d1;
}

.content {
  flex: 1;
}

.title {
  font-size: 18px;
  font-weight: 600;
  margin-bottom: 8px;
  color: #2c3e50;
}

.description {
  color: #7f8c8d;
  margin-bottom: 8px;
  line-height: 1.4;
}

.meta {
  display: flex;
  gap: 16px;
  font-size: 14px;
  color: #95a5a6;
}

.meta span {
  padding: 4px 8px;
  background: #f8f9fa;
  border-radius: 4px;
}

/* 移动端适配 */
@media (max-width: 768px) {
  .search-title {
    font-size: 28px;
  }
  
  .search-subtitle {
    font-size: 16px;
  }
  
  .search-box {
    flex-direction: column;
  }
  
  .category-tags {
    gap: 8px;
  }
  
  .category-tag {
    font-size: 12px;
    padding: 6px 12px;
  }
  
  .recommend-item {
    flex-direction: column;
    text-align: center;
  }
  
  .content {
    text-align: center;
  }
  
  .meta {
    justify-content: center;
  }
}
</style>