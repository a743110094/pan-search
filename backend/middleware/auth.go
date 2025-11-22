package middleware

import (
	"pan-search-api/common"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware 认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			common.Unauthorized(c, "缺少认证token")
			c.Abort()
			return
		}

		// 检查Bearer token格式
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			common.Unauthorized(c, "token格式错误")
			c.Abort()
			return
		}

		token := parts[1]
		if token == "" {
			common.Unauthorized(c, "token不能为空")
			c.Abort()
			return
		}

		// TODO: 验证token有效性
		// 这里可以集成JWT验证或其他认证方式
		// 暂时简单验证token是否存在
		if !isValidToken(token) {
			common.Unauthorized(c, "无效的token")
			c.Abort()
			return
		}

		// 将用户信息存储到上下文中
		// 这里可以解析token获取用户信息
		userID := extractUserIDFromToken(token)
		c.Set("userID", userID)

		c.Next()
	}
}

// isValidToken 验证token有效性
func isValidToken(token string) bool {
	// TODO: 实现实际的token验证逻辑
	// 这里简单返回true用于演示
	return token != "invalid"
}

// extractUserIDFromToken 从token中提取用户ID
func extractUserIDFromToken(token string) string {
	// TODO: 实现实际的token解析逻辑
	// 这里返回一个模拟的用户ID
	return "user_001"
}

// OptionalAuthMiddleware 可选认证中间件
func OptionalAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.Next()
			return
		}

		// 检查Bearer token格式
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.Next()
			return
		}

		token := parts[1]
		if token == "" {
			c.Next()
			return
		}

		// 验证token有效性
		if isValidToken(token) {
			userID := extractUserIDFromToken(token)
			c.Set("userID", userID)
		}

		c.Next()
	}
}

// AdminAuthMiddleware 管理员认证中间件
func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 先进行普通认证
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			common.Unauthorized(c, "缺少认证token")
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			common.Unauthorized(c, "token格式错误")
			c.Abort()
			return
		}

		token := parts[1]
		if token == "" {
			common.Unauthorized(c, "token不能为空")
			c.Abort()
			return
		}

		if !isValidToken(token) {
			common.Unauthorized(c, "无效的token")
			c.Abort()
			return
		}

		// TODO: 检查用户是否为管理员
		// 这里简单验证，实际应该检查用户角色
		userID := extractUserIDFromToken(token)
		if !isAdmin(userID) {
			common.Forbidden(c, "需要管理员权限")
			c.Abort()
			return
		}

		c.Set("userID", userID)
		c.Set("isAdmin", true)

		c.Next()
	}
}

// isAdmin 检查用户是否为管理员
func isAdmin(userID string) bool {
	// TODO: 实现实际的管理员检查逻辑
	// 这里简单返回true用于演示
	return userID == "admin_001"
}