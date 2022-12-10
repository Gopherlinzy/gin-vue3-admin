// Package role 模型
package role

import (
	"github.com/Gopherlinzy/gin-vue3-admin/app/models"
	"github.com/Gopherlinzy/gin-vue3-admin/app/models/api"
	"github.com/Gopherlinzy/gin-vue3-admin/app/models/menu"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/database"
	"strconv"
)

type Role struct {
	models.BaseModel

	RoleName string `json:"role_name,omitempty"`
	Des      string `json:"des,omitempty"`
	Status   bool   `json:"status,omitempty"`

	Menus []menu.Menu `gorm:"many2many:roles_menus;"`
	Apis  []api.Api   `gorm:"many2many:roles_apis;"`

	models.CommonTimestampsField
}

func (role *Role) Create() {
	database.Gohub_DB.Create(&role)
}

func (role *Role) Save() (rowsAffected int64) {
	result := database.Gohub_DB.Save(&role)
	return result.RowsAffected
}

func (role *Role) Delete() (rowsAffected int64) {
	result := database.Gohub_DB.Delete(&role)
	return result.RowsAffected
}

func (role *Role) GetAssociationsMenus() (data []menu.Menu) {
	var menus []menu.Menu
	database.Gohub_DB.Model(&role).Association("Menus").Find(&menus)
	return menus

	return
}

func (role *Role) AssociationClear(tableName string) (err error) {
	err = database.Gohub_DB.Model(&role).Association(tableName).Clear()
	return
}

func (role *Role) AppendAssociation(tableName string, AssIDS []string) (err error) {
	role.AssociationClear(tableName)
	if tableName == "Menus" {
		var ass []menu.Menu
		for _, v := range AssIDS {
			id, _ := strconv.Atoi(v)
			ass = append(ass, menu.Menu{BaseModel: models.BaseModel{ID: uint64(id)}})
		}
		err = database.Gohub_DB.Model(&role).Association(tableName).Append(&ass)
	} else if tableName == "Apis" {
		var ass []api.Api
		for _, v := range AssIDS {
			id, _ := strconv.Atoi(v)
			ass = append(ass, api.Api{BaseModel: models.BaseModel{ID: uint64(id)}})
		}
		err = database.Gohub_DB.Model(&role).Association(tableName).Append(&ass)
	}
	return
}
