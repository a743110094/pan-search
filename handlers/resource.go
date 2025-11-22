package handlers

import (
	"pan-search-api/common"
	"pan-search-api/database"
	"pan-search-api/models"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SearchResources 搜索资源
// @Summary 搜索资源
// @Description 根据关键词和分类搜索网盘资源
// @Tags resources
// @Accept json
// @Produce json
// @Param q query string true "搜索关键词"
// @Param category query string false "资源分类 (all/movie/tv/music/software/document/other)"
// @Param sort query string false "排序方式 (relevance/time/size)"
// @Param page query int false "页码，默认1"
// @Param pageSize query int false "每页数量，默认10"
// @Success 200 {object} common.Response{data=common.PaginatedResponse{list=[]models.ResourceResponse}}
// @Router /resources/search [get]
func SearchResources(c *gin.Context) {
	var req models.SearchRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		common.BadRequest(c, "请求参数错误")
		return
	}

	// 设置默认值
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}
	if req.Sort == "" {
		req.Sort = "relevance"
	}

	// 构建查询
	db := database.DB.Model(&models.Resource{}).
		Preload("Category").
		Preload("Tags").
		Where("valid = ?", true)

	// 关键词搜索
	if req.Q != "" {
		keywords := strings.Split(req.Q, " ")
		for _, keyword := range keywords {
			if keyword != "" {
				db = db.Where("title LIKE ? OR description LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
			}
		}
	}

	// 分类筛选
	if req.Category != "" && req.Category != "all" {
		db = db.Joins("JOIN categories ON resources.category_id = categories.id").
			Where("categories.value = ?", req.Category)
	}

	// 排序
	switch req.Sort {
	case "time":
		db = db.Order("upload_time DESC")
	case "size":
		// 这里需要将size转换为数字进行排序，暂时按字符串排序
		db = db.Order("size DESC")
	default: // relevance
		db = db.Order("view_count DESC, download_count DESC")
	}

	// 获取总数
	var total int64
	if err := db.Count(&total).Error; err != nil {
		common.InternalServerError(c, "查询失败")
		return
	}

	// 分页查询
	offset := (req.Page - 1) * req.PageSize
	var resources []models.Resource
	if err := db.Offset(offset).Limit(req.PageSize).Find(&resources).Error; err != nil {
		common.InternalServerError(c, "查询失败")
		return
	}

	// 转换为响应格式
	var resourceList []models.ResourceResponse
	for _, resource := range resources {
		tags := make([]string, len(resource.Tags))
		for i, tag := range resource.Tags {
			tags[i] = tag.TagName
		}

		resourceList = append(resourceList, models.ResourceResponse{
			ID:            resource.ID,
			Title:         resource.Title,
			Description:   resource.Description,
			Size:          resource.Size,
			Type:          resource.Type,
			Category:      resource.Category.Label,
			Source:        resource.Source,
			DownloadURL:   resource.DownloadURL,
			ExtractCode:   resource.ExtractCode,
			UploadTime:    resource.UploadTime,
			FileCount:     resource.FileCount,
			ViewCount:     resource.ViewCount,
			DownloadCount: resource.DownloadCount,
			Tags:          tags,
			Valid:         resource.Valid,
			ExpireTime:    resource.ExpireTime,
		})
	}

	// 记录搜索日志（异步）
	go recordSearchLog(c, req, int(total))

	common.SuccessWithPagination(c, resourceList, req.Page, req.PageSize, int(total))
}

// GetHotResources 获取热门推荐
// @Summary 获取热门推荐
// @Description 获取热门资源推荐列表
// @Tags resources
// @Accept json
// @Produce json
// @Param limit query int false "返回数量，默认10"
// @Success 200 {object} common.Response{data=[]models.HotResourceResponse}
// @Router /resources/hot [get]
func GetHotResources(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}
	if limit > 50 {
		limit = 50
	}

	var resources []models.Resource
	if err := database.DB.Model(&models.Resource{}).
		Preload("Category").
		Where("valid = ?", true).
		Order("view_count DESC, download_count DESC").
		Limit(limit).
		Find(&resources).Error; err != nil {
		common.InternalServerError(c, "查询失败")
		return
	}

	var hotResources []models.HotResourceResponse
	for i, resource := range resources {
		// 简单计算趋势（这里使用随机值模拟）
		trend := "stable"
		trendValue := 0
		if i%3 == 0 {
			trend = "up"
			trendValue = 15
		} else if i%3 == 1 {
			trend = "down"
			trendValue = -8
		}

		hotResources = append(hotResources, models.HotResourceResponse{
			ID:          resource.ID,
			Rank:        i + 1,
			Title:       resource.Title,
			Description: resource.Description,
			Size:        resource.Size,
			Type:        resource.Type,
			Category:    resource.Category.Label,
			SearchCount: resource.ViewCount,
			Trend:       trend,
			TrendValue:  trendValue,
		})
	}

	common.Success(c, map[string]interface{}{
		"list": hotResources,
	})
}

// RecordDownload 记录资源下载
// @Summary 记录资源下载
// @Description 记录资源下载行为
// @Tags resources
// @Accept json
// @Produce json
// @Param id path string true "资源ID"
// @Param body body models.DownloadRecordCreate true "下载记录信息"
// @Security BearerAuth
// @Success 200 {object} common.Response{data=models.DownloadResponse}
// @Router /resources/{id}/download [post]
func RecordDownload(c *gin.Context) {
	resourceID := c.Param("id")
	if resourceID == "" {
		common.BadRequest(c, "资源ID不能为空")
		return
	}

	var req models.DownloadRecordCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		common.BadRequest(c, "请求参数错误")
		return
	}

	// 检查资源是否存在
	var resource models.Resource
	if err := database.DB.First(&resource, "id = ?", resourceID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			common.NotFound(c, "资源不存在")
		} else {
			common.InternalServerError(c, "查询失败")
		}
		return
	}

	// 创建下载记录
	downloadRecord := models.DownloadRecord{
		ID:           generateID(),
		ResourceID:   resourceID,
		UserID:       req.UserID,
		UserAgent:    req.UserAgent,
		IPAddress:    req.IP,
		DownloadTime: time.Now(),
	}

	if err := database.DB.Create(&downloadRecord).Error; err != nil {
		common.InternalServerError(c, "创建下载记录失败")
		return
	}

	// 更新资源的下载计数
	database.DB.Model(&resource).Update("download_count", gorm.Expr("download_count + ?", 1))

	response := models.DownloadResponse{
		DownloadID:  downloadRecord.ID,
		ResourceID:  resourceID,
		DownloadURL: resource.DownloadURL,
		ExpireTime:  time.Now().Add(24 * time.Hour), // 24小时后过期
	}

	common.Success(c, response)
}

// 记录搜索日志
func recordSearchLog(c *gin.Context, req models.SearchRequest, resultCount int) {
	searchRecord := models.SearchRecord{
		Keyword:     req.Q,
		Category:    req.Category,
		SortBy:      req.Sort,
		UserID:      "", // 可以从上下文中获取
		IPAddress:   c.ClientIP(),
		UserAgent:   c.GetHeader("User-Agent"),
		ResultCount: uint(resultCount),
		SearchTime:  time.Now(),
	}

	database.DB.Create(&searchRecord)
}

// 生成ID（简化实现）
func generateID() string {
	// TODO: 实现更复杂的ID生成逻辑
	return "dl_" + strconv.FormatInt(time.Now().UnixNano(), 10)
}