<template>
  <div class="request-page">
    <div class="container">
      <div class="request-container">
        <div class="request-header">
          <h2 class="page-title">资源求助</h2>
          <p class="page-subtitle">找不到您需要的资源？告诉我们，我们会尽快为您寻找并通知您</p>
        </div>

        <div class="request-form-container">
          <form @submit.prevent="submitRequest" class="request-form">
            <div class="form-group">
              <label for="resourceName" class="form-label">资源名称 *</label>
              <input
                id="resourceName"
                v-model="form.resourceName"
                type="text"
                class="form-input"
                placeholder="请输入您需要查找的资源名称"
                required
              />
            </div>

            <div class="form-group">
              <label for="resourceType" class="form-label">资源类型</label>
              <select
                id="resourceType"
                v-model="form.resourceType"
                class="form-select"
              >
                <option value="">请选择资源类型</option>
                <option value="movie">电影</option>
                <option value="tv">电视剧</option>
                <option value="music">音乐</option>
                <option value="software">软件</option>
                <option value="document">文档</option>
                <option value="other">其他</option>
              </select>
            </div>

            <div class="form-group">
              <label for="description" class="form-label">详细描述</label>
              <textarea
                id="description"
                v-model="form.description"
                class="form-textarea"
                rows="4"
                placeholder="请详细描述您需要的资源，包括版本、格式、大小等信息"
              ></textarea>
            </div>

            <div class="form-group">
              <label for="contact" class="form-label">联系方式 *</label>
              <input
                id="contact"
                v-model="form.contact"
                type="text"
                class="form-input"
                placeholder="请输入您的邮箱或手机号"
                required
              />
              <p class="form-hint">当资源找到时，我们会通过此联系方式通知您</p>
            </div>

            <div class="form-actions">
              <button 
                type="submit" 
                class="submit-btn" 
                :disabled="submitting"
              >
                <span v-if="submitting" class="loading-spinner"></span>
                {{ submitting ? '提交中...' : '提交求助' }}
              </button>
              <button 
                type="button" 
                class="reset-btn" 
                @click="resetForm"
                :disabled="submitting"
              >
                重置
              </button>
            </div>
          </form>

          <!-- 提交成功提示 -->
          <div v-if="submitSuccess" class="success-message">
            <div class="success-icon">✅</div>
            <h3>提交成功！</h3>
            <p>我们已收到您的求助请求，找到相关资源后会第一时间通知您。</p>
            <button class="btn" @click="submitSuccess = false">继续提交</button>
          </div>
        </div>

        <!-- 求助说明 -->
        <div class="request-info">
          <h3>求助说明</h3>
          <div class="info-list">
            <div class="info-item">
              <div class="info-icon">📋</div>
              <div class="info-content">
                <h4>填写准确信息</h4>
                <p>请尽可能详细地描述您需要的资源，包括名称、版本、格式等信息</p>
              </div>
            </div>
            <div class="info-item">
              <div class="info-icon">⏱️</div>
              <div class="info-content">
                <h4>处理时间</h4>
                <p>我们会在1-3个工作日内处理您的求助请求</p>
              </div>
            </div>
            <div class="info-item">
              <div class="info-icon">📧</div>
              <div class="info-content">
                <h4>通知方式</h4>
                <p>找到资源后，我们会通过您提供的联系方式通知您</p>
              </div>
            </div>
            <div class="info-item">
              <div class="info-icon">🔍</div>
              <div class="info-content">
                <h4>搜索建议</h4>
                <p>建议先尝试使用搜索功能，可能已经有您需要的资源</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { throttle } from '../utils/debounce'

export default {
  name: 'Request',
  data() {
    return {
      form: {
        resourceName: '',
        resourceType: '',
        description: '',
        contact: ''
      },
      submitting: false,
      submitSuccess: false
    }
  },
  methods: {
    submitRequest: throttle(async function() {
      if (!this.validateForm()) {
        return
      }

      this.submitting = true

      try {
        // 模拟API调用
        await new Promise(resolve => setTimeout(resolve, 1500))
        
        // 模拟提交成功
        console.log('提交的求助信息:', this.form)
        this.submitSuccess = true
        this.resetForm()
      } catch (error) {
        console.error('提交失败:', error)
        alert('提交失败，请稍后重试')
      } finally {
        this.submitting = false
      }
    }, 1000),
    
    validateForm() {
      if (!this.form.resourceName.trim()) {
        alert('请输入资源名称')
        return false
      }
      
      if (!this.form.contact.trim()) {
        alert('请输入联系方式')
        return false
      }
      
      // 简单的联系方式验证
      const contact = this.form.contact.trim()
      const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
      const phoneRegex = /^1[3-9]\d{9}$/
      
      if (!emailRegex.test(contact) && !phoneRegex.test(contact)) {
        alert('请输入有效的邮箱地址或手机号码')
        return false
      }
      
      return true
    },
    
    resetForm() {
      this.form = {
        resourceName: '',
        resourceType: '',
        description: '',
        contact: ''
      }
    }
  }
}
</script>

<style scoped>
.request-page {
  padding: 40px 0;
}

.request-container {
  max-width: 800px;
  margin: 0 auto;
}

.request-header {
  text-align: center;
  margin-bottom: 40px;
}

.page-title {
  font-size: 32px;
  font-weight: 700;
  margin-bottom: 10px;
  color: #2c3e50;
}

.page-subtitle {
  font-size: 16px;
  color: #7f8c8d;
  line-height: 1.6;
}

.request-form-container {
  background: white;
  border-radius: 12px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
  padding: 40px;
  margin-bottom: 40px;
  position: relative;
}

.request-form {
  display: flex;
  flex-direction: column;
  gap: 24px;
}

.form-group {
  display: flex;
  flex-direction: column;
}

.form-label {
  font-weight: 600;
  margin-bottom: 8px;
  color: #2c3e50;
  font-size: 14px;
}

.form-input,
.form-select,
.form-textarea {
  padding: 12px 16px;
  border: 1px solid #ddd;
  border-radius: 8px;
  font-size: 16px;
  transition: all 0.3s ease;
  font-family: inherit;
  background-color: #ffffff;
}

.form-input:focus,
.form-select:focus,
.form-textarea:focus {
  outline: none;
  border-color: #bbb;
  background-color: #fdfdfd;
}

.form-textarea {
  resize: vertical;
  min-height: 100px;
}

.form-hint {
  font-size: 12px;
  color: #6c757d;
  margin-top: 4px;
}

.form-actions {
  display: flex;
  gap: 15px;
  margin-top: 10px;
}

.submit-btn {
  flex: 1;
  padding: 15px 20px;
  background: #6c757d;
  color: white;
  border: none;
  border-radius: 8px;
  font-size: 16px;
  font-weight: 600;
  cursor: pointer;
  transition: all 0.3s ease;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
}

.submit-btn:hover:not(:disabled) {
  background: #5a6268;
  transform: translateY(-1px);
}

.submit-btn:disabled {
  background: #6c757d;
  cursor: not-allowed;
  transform: none;
}

.reset-btn {
  padding: 15px 20px;
  background: #6c757d;
  color: white;
  border: none;
  border-radius: 8px;
  font-size: 16px;
  cursor: pointer;
  transition: all 0.3s ease;
  min-width: 100px;
}

.reset-btn:hover:not(:disabled) {
  background: #545b62;
}

.reset-btn:disabled {
  background: #adb5bd;
  cursor: not-allowed;
}

.loading-spinner {
  width: 16px;
  height: 16px;
  border: 2px solid transparent;
  border-top: 2px solid currentColor;
  border-radius: 50%;
  animation: spin 1s linear infinite;
}

@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}

/* 成功消息 */
.success-message {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: white;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  text-align: center;
  border-radius: 12px;
  padding: 40px;
}

.success-icon {
  font-size: 64px;
  margin-bottom: 20px;
}

.success-message h3 {
  font-size: 24px;
  margin-bottom: 10px;
  color: #2c3e50;
}

.success-message p {
  color: #7f8c8d;
  margin-bottom: 20px;
  line-height: 1.6;
}

/* 求助说明 */
.request-info {
  background: white;
  border-radius: 12px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
  padding: 40px;
}

.request-info h3 {
  font-size: 24px;
  margin-bottom: 30px;
  color: #2c3e50;
  text-align: center;
}

.info-list {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
  gap: 24px;
}

.info-item {
  display: flex;
  gap: 16px;
  align-items: flex-start;
}

.info-icon {
  font-size: 24px;
  flex-shrink: 0;
  margin-top: 4px;
}

.info-content h4 {
  font-size: 16px;
  font-weight: 600;
  margin-bottom: 8px;
  color: #2c3e50;
}

.info-content p {
  color: #7f8c8d;
  line-height: 1.5;
  font-size: 14px;
}

/* 移动端适配 */
@media (max-width: 768px) {
  .request-form-container {
    padding: 30px 20px;
  }
  
  .page-title {
    font-size: 28px;
  }
  
  .form-actions {
    flex-direction: column;
  }
  
  .reset-btn {
    min-width: auto;
  }
  
  .info-list {
    grid-template-columns: 1fr;
  }
  
  .info-item {
    flex-direction: column;
    text-align: center;
  }
  
  .info-icon {
    margin-top: 0;
  }
}

@media (max-width: 480px) {
  .request-page {
    padding: 20px 0;
  }
  
  .request-form-container,
  .request-info {
    padding: 20px 15px;
  }
  
  .success-message {
    padding: 20px;
  }
  
  .success-icon {
    font-size: 48px;
  }
}
</style>