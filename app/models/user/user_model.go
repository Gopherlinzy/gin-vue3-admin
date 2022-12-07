package user

import (
	"errors"
	"github.com/Gopherlinzy/gin-vue3-admin/app/models"
	casbins "github.com/Gopherlinzy/gin-vue3-admin/pkg/casbin"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/database"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/hash"
)

// User 用户模型
type User struct {
	models.BaseModel

	Name string `json:"name,omitempty"`

	Email    string `json:"email,omitempty"` // json:"-" json解析器忽略此字段
	Phone    string `json:"phone,omitempty"`
	Password string `json:"-"`

	City         string `json:"city,omitempty"`
	Introduction string `json:"introduction,omitempty"`
	Avatar       string `json:"avatar,omitempty"`

	Status bool `json:"status,omitempty"`

	RoleName string `json:"role_name,omitempty"`

	models.CommonTimestampsField
}

// Create 创建用户，通过 User.ID 来判断是否创建成功
func (userModel *User) Create() {
	database.Gohub_DB.Create(&userModel)
}

// ComparePassword 密码是否正确
func (userModel *User) ComparePassword(_password string) bool {
	return hash.BcryptCheck(_password, userModel.Password)
}

func (userModel *User) Save() (rowsAffected int64) {
	result := database.Gohub_DB.Save(&userModel)
	return result.RowsAffected
}

func (userRoleModel *User) AddRole() error {
	return casbins.NewCasbin().AddUserRole(userRoleModel.Name, userRoleModel.RoleName)
}

// UpdateRole 更新用户的所属角色
func (userRoleModel *User) UpdateRole(oldName, newRole string) error {
	cs := casbins.NewCasbin()
	oldRole := userRoleModel.RoleName
	userRoleModel.RoleName = newRole
	rowsAffected := userRoleModel.Save()
	if rowsAffected == 0 {
		return errors.New("更新用户的所属角色失败")
	}
	err := cs.UpdateUserRole(oldName, userRoleModel.Name, oldRole, newRole)
	if err != nil {
		return err
	}
	err = cs.Enforcer.InvalidateCache()
	if err != nil {
		return err
	}
	return err
}

func (user *User) Delete() (rowsAffected int64) {
	result := database.Gohub_DB.Delete(&user)
	return result.RowsAffected
}
