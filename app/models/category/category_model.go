// Package category 模型
package category

import (
	"github.com/Gopherlinzy/gin-vue3-admin/app/models"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/database"
	"gorm.io/gorm/clause"
)

type Category struct {
	models.BaseModel

	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`

	models.CommonTimestampsField
}

func (category *Category) Create() {
	database.Gohub_DB.Clauses(clause.OnConflict{UpdateAll: true}).Create(&category)
}

func (category *Category) Save() (rowsAffected int64) {
	result := database.Gohub_DB.Save(&category)
	return result.RowsAffected
}

func (category *Category) Delete() (rowsAffected int64) {
	result := database.Gohub_DB.Delete(&category)
	return result.RowsAffected
}
