package models

import (
	"time"
)

// Category 分类模型
type Category struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Value     string    `gorm:"size:50;uniqueIndex;not null" json:"value"`
	Label     string    `gorm:"size:50;not null" json:"label"`
	Icon      string    `gorm:"size:20;not null" json:"icon"`
	SortOrder int       `gorm:"default:0" json:"sort_order"`
	IsActive  bool      `gorm:"default:true" json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Resource 资源模型
type Resource struct {
	ID            string    `gorm:"primaryKey;size:32" json:"id"`
	Title         string    `gorm:"size:255;not null" json:"title"`
	Description   string    `gorm:"type:text" json:"description"`
	Size          string    `gorm:"size:20;not null" json:"size"`
	Type          string    `gorm:"size:50;not null" json:"type"`
	CategoryID    uint      `gorm:"not null" json:"category_id"`
	Category      Category  `gorm:"foreignKey:CategoryID" json:"category"`
	Source        string    `gorm:"size:100;not null" json:"source"`
	DownloadURL   string    `gorm:"size:500;not null" json:"download_url"`
	ExtractCode   string    `gorm:"size:20" json:"extract_code"`
	FileCount     uint      `gorm:"default:1" json:"file_count"`
	ViewCount     uint      `gorm:"default:0" json:"view_count"`
	DownloadCount uint      `gorm:"default:0" json:"download_count"`
	Valid         bool      `gorm:"default:true" json:"valid"`
	ExpireTime    time.Time `json:"expire_time"`
	UploadTime    time.Time `json:"upload_time"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Tags          []ResourceTag `gorm:"foreignKey:ResourceID" json:"tags"`
}

// ResourceTag 资源标签模型
type ResourceTag struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	ResourceID string    `gorm:"size:32;not null" json:"resource_id"`
	TagName    string    `gorm:"size:50;not null" json:"tag_name"`
	CreatedAt  time.Time `json:"created_at"`
}

// HelpRequest 求助请求模型
type HelpRequest struct {
	ID              string    `gorm:"primaryKey;size:32" json:"id"`
	ResourceName    string    `gorm:"size:255;not null" json:"resource_name"`
	ResourceType    string    `gorm:"size:50" json:"resource_type"`
	Description     string    `gorm:"type:text" json:"description"`
	Contact         string    `gorm:"size:100;not null" json:"contact"`
	ContactType     string    `gorm:"type:ENUM('email','phone');not null" json:"contact_type"`
	Status          string    `gorm:"type:ENUM('pending','processing','completed','rejected');default:'pending'" json:"status"`
	AdminNotes      string    `gorm:"type:text" json:"admin_notes"`
	CompletedTime   time.Time `json:"completed_time"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// DownloadRecord 下载记录模型
type DownloadRecord struct {
	ID           string    `gorm:"primaryKey;size:32" json:"id"`
	ResourceID   string    `gorm:"size:32;not null" json:"resource_id"`
	UserID       string    `gorm:"size:32" json:"user_id"`
	UserAgent    string    `gorm:"type:text" json:"user_agent"`
	IPAddress    string    `gorm:"size:45" json:"ip_address"`
	DownloadTime time.Time `json:"download_time"`
	CreatedAt    time.Time `json:"created_at"`
}

// SearchRecord 搜索记录模型
type SearchRecord struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Keyword     string    `gorm:"size:255;not null" json:"keyword"`
	Category    string    `gorm:"size:50" json:"category"`
	SortBy      string    `gorm:"size:50" json:"sort_by"`
	UserID      string    `gorm:"size:32" json:"user_id"`
	IPAddress   string    `gorm:"size:45" json:"ip_address"`
	UserAgent   string    `gorm:"type:text" json:"user_agent"`
	ResultCount uint      `gorm:"default:0" json:"result_count"`
	SearchTime  time.Time `json:"search_time"`
	CreatedAt   time.Time `json:"created_at"`
}

// User 用户模型
type User struct {
	ID           string    `gorm:"primaryKey;size:32" json:"id"`
	Username     string    `gorm:"size:50;uniqueIndex;not null" json:"username"`
	Email        string    `gorm:"size:100;uniqueIndex;not null" json:"email"`
	Phone        string    `gorm:"size:20;uniqueIndex" json:"phone"`
	PasswordHash string    `gorm:"size:255;not null" json:"-"`
	Avatar       string    `gorm:"size:500" json:"avatar"`
	Status       string    `gorm:"type:ENUM('active','inactive','banned');default:'active'" json:"status"`
	LastLoginAt  time.Time `json:"last_login_at"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// SystemConfig 系统配置模型
type SystemConfig struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	ConfigKey   string    `gorm:"size:100;uniqueIndex;not null" json:"config_key"`
	ConfigValue string    `gorm:"type:text;not null" json:"config_value"`
	ConfigType  string    `gorm:"type:ENUM('string','number','boolean','json');default:'string'" json:"config_type"`
	Description string    `gorm:"size:255" json:"description"`
	IsPublic    bool      `gorm:"default:false" json:"is_public"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// API请求和响应模型

// SearchRequest 搜索请求
type SearchRequest struct {
	Q        string `form:"q" binding:"required" json:"q"`
	Category string `form:"category" json:"category"`
	Sort     string `form:"sort" json:"sort"`
	Page     int    `form:"page" json:"page"`
	PageSize int    `form:"pageSize" json:"pageSize"`
}

// HelpRequestCreate 创建求助请求
type HelpRequestCreate struct {
	ResourceName string `json:"resourceName" binding:"required"`
	ResourceType string `json:"resourceType"`
	Description  string `json:"description"`
	Contact      string `json:"contact" binding:"required"`
	ContactType  string `json:"contactType" binding:"required,oneof=email phone"`
}

// DownloadRecordCreate 创建下载记录
type DownloadRecordCreate struct {
	UserID    string `json:"userId"`
	UserAgent string `json:"userAgent"`
	IP        string `json:"ip"`
}

// SearchSuggestionRequest 搜索建议请求
type SearchSuggestionRequest struct {
	Q string `form:"q" binding:"required" json:"q"`
}

// API响应模型

// ResourceResponse 资源响应
type ResourceResponse struct {
	ID            string    `json:"id"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	Size          string    `json:"size"`
	Type          string    `json:"type"`
	Category      string    `json:"category"`
	Source        string    `json:"source"`
	DownloadURL   string    `json:"downloadUrl"`
	ExtractCode   string    `json:"extractCode"`
	UploadTime    time.Time `json:"uploadTime"`
	FileCount     uint      `json:"fileCount"`
	ViewCount     uint      `json:"viewCount"`
	DownloadCount uint      `json:"downloadCount"`
	Tags          []string  `json:"tags"`
	Valid         bool      `json:"valid"`
	ExpireTime    time.Time `json:"expireTime"`
}

// HotResourceResponse 热门资源响应
type HotResourceResponse struct {
	ID          string `json:"id"`
	Rank        int    `json:"rank"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Size        string `json:"size"`
	Type        string `json:"type"`
	Category    string `json:"category"`
	SearchCount uint   `json:"searchCount"`
	Trend       string `json:"trend"`
	TrendValue  int    `json:"trendValue"`
}

// CategoryResponse 分类响应
type CategoryResponse struct {
	Value string `json:"value"`
	Label string `json:"label"`
	Count int    `json:"count"`
	Icon  string `json:"icon"`
}

// HelpRequestResponse 求助请求响应
type HelpRequestResponse struct {
	RequestID           string `json:"requestId"`
	Status              string `json:"status"`
	SubmitTime          time.Time `json:"submitTime"`
	EstimatedProcessTime string `json:"estimatedProcessTime"`
}

// DownloadResponse 下载响应
type DownloadResponse struct {
	DownloadID  string    `json:"downloadId"`
	ResourceID  string    `json:"resourceId"`
	DownloadURL string    `json:"downloadUrl"`
	ExpireTime  time.Time `json:"expireTime"`
}

// SearchSuggestionResponse 搜索建议响应
type SearchSuggestionResponse struct {
	Suggestions []string `json:"suggestions"`
}