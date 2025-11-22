<template>
  <header class="header">
    <div class="container">
      <div class="header-content">
        <div class="logo">
          <h1>网盘搜索</h1>
          <span class="tagline">专业的网盘资源搜索平台</span>
        </div>
        <nav class="nav">
          <router-link to="/" class="nav-link">首页</router-link>
          <router-link to="/search" class="nav-link">搜索</router-link>
          <router-link to="/request" class="nav-link">资源求助</router-link>
        </nav>
      </div>
    </div>
  </header>
</template>

<script>
export default {
  name: 'Header'
}
</script>

<style scoped>
.header {
  background: white;
  border-bottom: 1px solid #e1e5e9;
  padding: 16px 0;
  position: sticky;
  top: 0;
  z-index: 100;
  backdrop-filter: blur(8px);
  background: rgba(255, 255, 255, 0.95);
}

.header-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
  flex-wrap: wrap;
  gap: 16px;
}

.logo h1 {
  font-size: 20px;
  font-weight: 700;
  margin-bottom: 2px;
  color: #1a1a1a;
  letter-spacing: -0.5px;
}

.tagline {
  font-size: 12px;
  color: #666;
  font-weight: 400;
}

.nav {
  display: flex;
  gap: 24px;
}

.nav-link {
  color: #666;
  text-decoration: none;
  font-weight: 500;
  padding: 8px 12px;
  border-radius: 6px;
  transition: all 0.2s ease;
  position: relative;
  font-size: 14px;
}

.nav-link:hover {
  background: #f6f8fa;
  color: #1a1a1a;
}

.nav-link.router-link-active {
  background: #0969da;
  color: white;
}

.nav-link.router-link-active::after {
  content: '';
  position: absolute;
  bottom: -16px;
  left: 50%;
  transform: translateX(-50%);
  width: 4px;
  height: 4px;
  background: #0969da;
  border-radius: 50%;
}

/* 移动端适配 */
@media (max-width: 768px) {
  .header {
    padding: 12px 0;
  }
  
  .header-content {
    flex-direction: column;
    text-align: center;
    gap: 12px;
  }
  
  .logo h1 {
    font-size: 18px;
  }
  
  .tagline {
    font-size: 11px;
  }
  
  .nav {
    gap: 16px;
  }
  
  .nav-link {
    padding: 6px 10px;
    font-size: 13px;
  }
  
  .nav-link.router-link-active::after {
    bottom: -12px;
  }
}

@media (max-width: 480px) {
  .nav {
    gap: 12px;
  }
  
  .nav-link {
    padding: 4px 8px;
    font-size: 12px;
  }
}
</style>