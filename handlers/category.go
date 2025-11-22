package handlers

import (
	"pan-search-api/common"
	"pan-search-api/database"
	"pan-search-api/models"

	"github.com/gin-gonic/gin"
)

// GetCategories 获取分类列表
// @Summary 获取分类列表
// @Description 获取所有资源分类
// @Tags categories
// @Accept json
// @Produce json
// @Success 200 {object} common.Response{data=[]models.CategoryResponse}
// @Router /categories [get]
func GetCategories(c *gin.Context) {
	var categories []models.Category
	if err := database.DB.Where("is_active = ?", true).
		Order("sort_order ASC").
		Find(&categories).Error; err != nil {
		common.InternalServerError(c, "查询失败")
		return
	}

	// 获取每个分类的资源数量
	var categoryList []models.CategoryResponse
	for _, category := range categories {
		var count int64
		database.DB.Model(&models.Resource{}).
			Where("category_id = ? AND valid = ?", category.ID, true).
			Count(&count)

		categoryList = append(categoryList, models.CategoryResponse{
			Value: category.Value,
			Label: category.Label,
			Count: int(count),
			Icon:  category.Icon,
		})
	}

	common.Success(c, map[string]interface{}{
		"list": categoryList,
	})
}