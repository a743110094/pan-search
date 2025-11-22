package routes

import (
	"pan-search-api/handlers"
	"pan-search-api/middleware"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRouter 设置路由
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// 添加中间件
	router.Use(middleware.CORSMiddleware())

	// Swagger文档
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API路由组
	api := router.Group("/api/v1")
	{
		// 搜索相关接口（无需认证）
		search := api.Group("/search")
		{
			search.GET("/suggestions", handlers.GetSearchSuggestions)
		}

		// 资源相关接口
		resources := api.Group("/resources")
		{
			resources.GET("/search", handlers.SearchResources)
			resources.GET("/hot", handlers.GetHotResources)
			resources.POST("/:id/download", middleware.AuthMiddleware(), handlers.RecordDownload)
		}

		// 分类相关接口（无需认证）
		categories := api.Group("/categories")
		{
			categories.GET("", handlers.GetCategories)
		}

		// 求助请求接口（无需认证）
		requests := api.Group("/requests")
		{
			requests.POST("", handlers.SubmitHelpRequest)
		}

		// 管理接口（需要管理员认证）
		admin := api.Group("/admin")
		admin.Use(middleware.AdminAuthMiddleware())
		{
			// 这里可以添加管理相关的接口
			// admin.GET("/help-requests", handlers.AdminGetHelpRequests)
			// admin.PUT("/help-requests/:id", handlers.AdminUpdateHelpRequest)
		}
	}

	// 健康检查
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			"timestamp": time.Now().UnixMilli(),
		})
	})

	return router
}