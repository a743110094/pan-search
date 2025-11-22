package handlers

import (
	"pan-search-api/common"
	"pan-search-api/database"
	"pan-search-api/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// SubmitHelpRequest 提交资源求助
// @Summary 提交资源求助
// @Description 提交资源求助请求
// @Tags requests
// @Accept json
// @Produce json
// @Param body body models.HelpRequestCreate true "求助请求信息"
// @Success 200 {object} common.Response{data=models.HelpRequestResponse}
// @Router /requests [post]
func SubmitHelpRequest(c *gin.Context) {
	var req models.HelpRequestCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		common.BadRequest(c, "请求参数错误")
		return
	}

	// 创建求助请求
	helpRequest := models.HelpRequest{
		ID:           generateRequestID(),
		ResourceName: req.ResourceName,
		ResourceType: req.ResourceType,
		Description:  req.Description,
		Contact:      req.Contact,
		ContactType:  req.ContactType,
		Status:       "pending",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := database.DB.Create(&helpRequest).Error; err != nil {
		common.InternalServerError(c, "提交求助请求失败")
		return
	}

	response := models.HelpRequestResponse{
		RequestID:           helpRequest.ID,
		Status:              helpRequest.Status,
		SubmitTime:          helpRequest.CreatedAt,
		EstimatedProcessTime: "1-3个工作日",
	}

	common.Success(c, response)
}

// 生成求助请求ID
func generateRequestID() string {
	return "req_" + strconv.FormatInt(time.Now().UnixNano(), 10)
}