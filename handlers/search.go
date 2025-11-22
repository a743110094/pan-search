package handlers

import (
	"pan-search-api/common"
	"pan-search-api/database"
	"pan-search-api/models"
	"strings"

	"github.com/gin-gonic/gin"
)

// GetSearchSuggestions 获取搜索建议
// @Summary 获取搜索建议
// @Description 获取搜索关键词建议
// @Tags search
// @Accept json
// @Produce json
// @Param q query string true "搜索关键词"
// @Success 200 {object} common.Response{data=models.SearchSuggestionResponse}
// @Router /search/suggestions [get]
func GetSearchSuggestions(c *gin.Context) {
	var req models.SearchSuggestionRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		common.BadRequest(c, "请求参数错误")
		return
	}

	if req.Q == "" {
		common.Success(c, models.SearchSuggestionResponse{
			Suggestions: []string{},
		})
		return
	}

	// 从搜索记录中获取热门建议
	var suggestions []string

	// 1. 从搜索记录中获取相关关键词
	var searchRecords []models.SearchRecord
	database.DB.Select("DISTINCT keyword").
		Where("keyword LIKE ?", req.Q+"%").
		Order("result_count DESC").
		Limit(5).
		Find(&searchRecords)

	for _, record := range searchRecords {
		suggestions = append(suggestions, record.Keyword)
	}

	// 2. 从资源标题中获取建议
	var resources []models.Resource
	database.DB.Select("DISTINCT title").
		Where("title LIKE ? AND valid = ?", "%"+req.Q+"%", true).
		Order("view_count DESC").
		Limit(5).
		Find(&resources)

	for _, resource := range resources {
		// 提取关键词相关的部分
		title := resource.Title
		if strings.Contains(title, req.Q) {
			// 简单的关键词提取逻辑
			words := strings.Fields(title)
			for _, word := range words {
				if strings.Contains(word, req.Q) && len(word) > len(req.Q) {
					suggestion := req.Q + word[len(req.Q):]
					if !contains(suggestions, suggestion) && len(suggestions) < 10 {
						suggestions = append(suggestions, suggestion)
					}
				}
			}
		}
	}

	// 3. 如果建议不足，添加一些通用建议
	if len(suggestions) < 5 {
		commonSuggestions := getCommonSuggestions(req.Q)
		for _, suggestion := range commonSuggestions {
			if !contains(suggestions, suggestion) && len(suggestions) < 10 {
				suggestions = append(suggestions, suggestion)
			}
		}
	}

	// 确保不超过10个建议
	if len(suggestions) > 10 {
		suggestions = suggestions[:10]
	}

	response := models.SearchSuggestionResponse{
		Suggestions: suggestions,
	}

	common.Success(c, response)
}

// 检查切片是否包含元素
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// 获取通用建议
func getCommonSuggestions(query string) []string {
	commonSuffixes := []string{"合集", "资源", "下载", "最新", "高清", "免费", "完整版", "国语", "英语"}
	var suggestions []string

	for _, suffix := range commonSuffixes {
		suggestion := query + suffix
		suggestions = append(suggestions, suggestion)
	}

	return suggestions
}