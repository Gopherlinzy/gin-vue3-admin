// Package api 模型
package api

import (
	"github.com/Gopherlinzy/gin-vue3-admin/app/models"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/database"
)

type Api struct {
	models.BaseModel

	Path    string `json:"api_path"`
	Group   string `json:"api_group"`
	Des     string `json:"api_des,omitempty"`
	Request string `json:"request"`

	//Roles []role.Role `gorm:"many2many:roles_apis;"`

	models.CommonTimestampsField
}

func (api *Api) Create() {
	database.Gohub_DB.Create(&api)
}

func (api *Api) Save() (rowsAffected int64) {
	result := database.Gohub_DB.Save(&api)
	return result.RowsAffected
}

func (api *Api) Delete() (rowsAffected int64) {
	result := database.Gohub_DB.Delete(&api)
	return result.RowsAffected
}
