// Package menu 模型
package menu

import (
	"github.com/Gopherlinzy/gin-vue3-admin/app/models"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/database"
)

type Menu struct {
	models.BaseModel

	Name       string `json:"name"`
	Permission string `json:"permissions"`
	RouterName string `json:"router_name"`
	RouterPath string `json:"router_path"`
	FatherID   uint64 `json:"father_id,omitempty"`
	VuePath    string `json:"vue_path,omitempty"`
	Status     bool   `json:"status,omitempty"`

	//Children []Menu `json:"children" gorm:"-"`
	//Roles    []role.Role `gorm:"many2many:roles_menus;"`

	models.CommonTimestampsField
}

func (menu *Menu) Create() {
	database.Gohub_DB.Create(&menu)
}

func (menu *Menu) Save() (rowsAffected int64) {
	result := database.Gohub_DB.Save(&menu)
	return result.RowsAffected
}

func (menu *Menu) Delete() (rowsAffected int64) {
	result := database.Gohub_DB.Delete(&menu)
	return result.RowsAffected
}
