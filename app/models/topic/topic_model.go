// Package topic 模型
package topic

import (
	"github.com/Gopherlinzy/gin-vue3-admin/app/models"
	"github.com/Gopherlinzy/gin-vue3-admin/app/models/category"
	"github.com/Gopherlinzy/gin-vue3-admin/app/models/user"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/database"
	"gorm.io/gorm/clause"
)

type Topic struct {
	models.BaseModel

	Title      string `json:"title,omitempty"`
	Body       string `json:"body,omitempty"`
	UserID     string `json:"user_id,omitempty"`
	CategoryID string `json:"category_id,omitempty"`

	// 通过 user_id 关联用户
	User user.User `json:"user"`

	// 通过 category_id 关联分类
	Category category.Category `json:"category"`

	models.CommonTimestampsField
}

func (topic *Topic) Create() {
	database.Gohub_DB.Clauses(clause.OnConflict{UpdateAll: true}).Create(&topic)
}

func (topic *Topic) Save() (rowsAffected int64) {
	result := database.Gohub_DB.Save(&topic)
	return result.RowsAffected
}

func (topic *Topic) Delete() (rowsAffected int64) {
	result := database.Gohub_DB.Delete(&topic)
	return result.RowsAffected
}
