package main

import (
	"log"
	"strconv"
	"pan-search-api/config"
	"pan-search-api/database"
	"pan-search-api/routes"

	"github.com/gin-gonic/gin"
)

// @title 网盘资源搜索平台 API
// @version 1.0
// @description 网盘资源搜索平台后端API接口
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Bearer token认证，格式: "Bearer {token}"

func main() {
	// 初始化配置
	config.Init()

	// 初始化数据库
	if err := database.Init(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.Close()

	// 设置Gin模式
	if config.GlobalConfig.App.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// 设置路由
	router := routes.SetupRouter()

	// 启动服务器
	port := ":" + strconv.Itoa(config.GlobalConfig.App.Port)
	log.Printf("Server starting on port %s", port)
	log.Printf("Swagger documentation available at http://localhost%s/swagger/index.html", port)

	if err := router.Run(port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}