// Package role 模型
package role

import (
	"github.com/Gopherlinzy/gin-vue3-admin/app/models"
	"github.com/Gopherlinzy/gin-vue3-admin/app/models/api"
	"github.com/Gopherlinzy/gin-vue3-admin/app/models/menu"
	casbins "github.com/Gopherlinzy/gin-vue3-admin/pkg/casbin"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/database"
	"gorm.io/gorm/clause"
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
	database.Gohub_DB.Clauses(clause.OnConflict{UpdateAll: true}).Create(&role)
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
}

func (role *Role) AssociationClear(tableName string) (err error) {
	err = database.Gohub_DB.Model(&role).Association(tableName).Clear()
	return
}

func (role *Role) AppendAssociation(tableName string, AssPolicyS []string) (err error) {
	role.AssociationClear(tableName)
	if tableName == "Menus" {
		var ass []menu.Menu
		for _, v := range AssPolicyS {
			id, _ := strconv.Atoi(v)
			ass = append(ass, menu.Menu{BaseModel: models.BaseModel{ID: uint64(id)}})
		}
		err = database.Gohub_DB.Model(&role).Association(tableName).Append(&ass)
	} else if tableName == "Apis" {
		var ass, apis []api.Api
		rules := make([][]string, 0, len(AssPolicyS))
		database.Gohub_DB.Select("id", "path", "method").Find(&apis)
		for _, v := range AssPolicyS {
			id, _ := strconv.Atoi(v)
			rules = append(rules, []string{role.RoleName, apis[id-1].Path, apis[id-1].Method})
			ass = append(ass, api.Api{BaseModel: models.BaseModel{ID: uint64(id)}})
		}
		cs := casbins.NewCasbin()
		cs.DeleteRole(role.RoleName)
		cs.AddPolicies(rules)
		err = database.Gohub_DB.Model(&role).Association(tableName).Append(&ass)
	}
	return
}
