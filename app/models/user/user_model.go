package user

import (
	"github.com/Gopherlinzy/gohub/app/models"
	casbins "github.com/Gopherlinzy/gohub/pkg/casbin"
	"github.com/Gopherlinzy/gohub/pkg/database"
	"github.com/Gopherlinzy/gohub/pkg/hash"
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

	models.CommonTimestampsField
}

type UserRole struct {
	Name     string `json:"name,omitempty"`
	RoleName string `json:"role_name,omitempty"`
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

func (userRoleModel *UserRole) AddRole() error {
	return casbins.NewCasbin().AddUserRole(userRoleModel.Name, userRoleModel.RoleName)
}

func (userRoleModel *UserRole) UpdateRole() error {
	cs := casbins.NewCasbin()
	role := cs.GetRolesForUser(userRoleModel.Name)
	if len(role) == 0 {
		userRoleModel.AddRole()
	}
	err := cs.UpdateUserRole(userRoleModel.Name, role[0], userRoleModel.RoleName)
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
