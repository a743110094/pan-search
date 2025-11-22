/**
 * 防抖函数
 * @param {Function} func 要防抖的函数
 * @param {number} wait 等待时间(毫秒)
 * @param {boolean} immediate 是否立即执行
 * @returns {Function} 防抖后的函数
 */
export function debounce(func, wait, immediate = false) {
  let timeoutId
  
  return function executedFunction(...args) {
    const context = this
    
    const later = function() {
      timeoutId = null
      if (!immediate) {
        func.apply(context, args)
      }
    }
    
    const callNow = immediate && !timeoutId
    
    clearTimeout(timeoutId)
    timeoutId = setTimeout(later, wait)
    
    if (callNow) {
      func.apply(context, args)
    }
  }
}

/**
 * 节流函数
 * @param {Function} func 要节流的函数
 * @param {number} limit 时间限制(毫秒)
 * @returns {Function} 节流后的函数
 */
export function throttle(func, limit) {
  let inThrottle
  let lastFunc
  let lastRan
  
  return function(...args) {
    const context = this
    
    if (!inThrottle) {
      func.apply(context, args)
      lastRan = Date.now()
      inThrottle = true
    } else {
      clearTimeout(lastFunc)
      lastFunc = setTimeout(function() {
        if ((Date.now() - lastRan) >= limit) {
          func.apply(context, args)
          lastRan = Date.now()
        }
      }, limit - (Date.now() - lastRan))
    }
  }
}