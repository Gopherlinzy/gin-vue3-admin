package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type RoleRequest struct {
	ID       string `valid:"id" json:"id"`
	RoleName string `json:"role_name,omitempty" valid:"role_name"`
	Des      string `json:"des,omitempty" valid:"des"`
}

func RoleSave(data interface{}, c *gin.Context) map[string][]string {

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

type RoleDeleteRequest struct {
	ID string `valid:"id" json:"id"`
}

func RoleDelete(data interface{}, c *gin.Context) map[string][]string {

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
