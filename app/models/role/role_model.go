// Package role 模型
package role

import (
	"github.com/Gopherlinzy/gohub/app/models"
	"github.com/Gopherlinzy/gohub/pkg/database"
)

type Role struct {
	models.BaseModel

	RoleName string `json:"role_name,omitempty"`
	Des      string `json:"des,omitempty"`
	Status   bool   `json:"status,omitempty"`

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
