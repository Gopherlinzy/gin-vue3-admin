// Package link 模型
package link

import (
	"github.com/Gopherlinzy/gin-vue3-admin/app/models"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/database"
)

type Link struct {
	models.BaseModel

	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`

	models.CommonTimestampsField
}

func (link *Link) Create() {
	database.Gohub_DB.Create(&link)
}

func (link *Link) Save() (rowsAffected int64) {
	result := database.Gohub_DB.Save(&link)
	return result.RowsAffected
}

func (link *Link) Delete() (rowsAffected int64) {
	result := database.Gohub_DB.Delete(&link)
	return result.RowsAffected
}
