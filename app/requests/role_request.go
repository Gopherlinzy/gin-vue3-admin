package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type RoleStoreRequest struct {
	RoleName string `json:"role_name,omitempty" valid:"role_name"`
	Des      string `json:"des,omitempty" valid:"des"`
}

func RoleStore(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"role_name": []string{"required", "between:3,20", "not_exists:roles,role_name"},
		"des":       []string{"min_cn:3", "max_cn:255"},
	}
	messages := govalidator.MapData{
		"role_name": []string{
			"required:角色名为必填项",
			"between:角色名长度需在 3~20 之间",
			"not_exists:角色名已存在",
		},
		"des": []string{
			"min_cn:描述长度需至少 3 个字",
			"max_cn:描述长度不能超过 255 个字",
		},
	}
	return validate(data, rules, messages)
}

type RoleUpdateRequest struct {
	ID       string `valid:"id" json:"id"`
	RoleName string `json:"role_name,omitempty" valid:"role_name"`
	Des      string `json:"des,omitempty" valid:"des"`
}

func RoleUpdate(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"id":        []string{"numeric", "exists:roles,id"},
		"role_name": []string{"required", "between:3,20"},
		"des":       []string{"min_cn:3", "max_cn:255"},
	}
	messages := govalidator.MapData{
		"id": []string{
			"numeric:id必须为数字",
			"exists:id必须存在",
		},
		"role_name": []string{
			"required:角色名为必填项",
			"between:角色名长度需在 3~20 之间",
		},
		"des": []string{
			"min_cn:描述长度需至少 3 个字",
			"max_cn:描述长度不能超过 255 个字",
		},
	}
	return validate(data, rules, messages)
}

type RoleIDRequest struct {
	ID string `valid:"id" json:"id"`
}

func RoleID(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"id": []string{"numeric", "exists:roles,id"},
	}
	messages := govalidator.MapData{
		"id": []string{
			"numeric:id必须为数字",
			"exists:id必须存在",
		},
	}
	return validate(data, rules, messages)
}

type UpdateRoleStatusRequest struct {
	ID     string `valid:"id" json:"id"`
	Status string `valid:"status" json:"status"`
}

func UpdateRoleStatus(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"id":     []string{"numeric", "exists:roles,id"},
		"status": []string{"required", "bool"},
	}
	messages := govalidator.MapData{
		"id": []string{
			"numeric:id必须为数字",
			"exists:id必须存在",
		},
		"status": []string{
			"required:status为必填项",
			"bool:必须得为布尔类型:true, false, 1, 0, \"1\" and \"0\"",
		},
	}
	return validate(data, rules, messages)
}
