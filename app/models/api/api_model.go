// Package api 模型
package api

import (
	"github.com/Gopherlinzy/gin-vue3-admin/app/models"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/database"
	"gorm.io/gorm/clause"
)

type Api struct {
	models.BaseModel

	Path        string `json:"path"`
	Description string `json:"description,omitempty"`
	ApiGroup    string `json:"api_group"`
	Method      string `json:"method"`

	//Roles []role.Role `gorm:"many2many:roles_apis;"`

	models.CommonTimestampsField
}

func (api *Api) Create() {
	database.Gohub_DB.Clauses(clause.OnConflict{UpdateAll: true}).Create(&api)
}

func (api *Api) Save() (rowsAffected int64) {
	result := database.Gohub_DB.Save(&api)
	return result.RowsAffected
}

func (api *Api) Delete() (rowsAffected int64) {
	result := database.Gohub_DB.Delete(&api)
	return result.RowsAffected
}
