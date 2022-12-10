package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type ApiSaveRequest struct {
	ID string `valid:"id" json:"id"`

	Path        string `json:"path" valid:"path"`
	ApiGroup    string `json:"api_group" valid:"api_group"`
	Description string `json:"description" valid:"description"`
	Method      string `json:"method" valid:"method"`
}

func ApiSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"id":          []string{"numeric", "exists:apis,id"},
		"path":        []string{"required", "min_cn:3", "max_cn:255", "not_exists:apis,path"},
		"api_group":   []string{"required", "min_cn:3", "max_cn:255"},
		"description": []string{"min_cn:3", "max_cn:255"},
		"method":      []string{"required", "min_cn:3", "max_cn:20"},
	}
	messages := govalidator.MapData{
		"id": []string{
			"numeric:id必须为数字",
			"exists:id必须存在",
		},
		"path": []string{
			"required:API路径为必填项",
			"min_cn:名称长度需至少 3 个字",
			"max_cn:名称长度不能超过 255 个字",
			"not_exists:名称已存在",
		},
		"api_group": []string{
			"required:API分组为必填项",
			"min_cn:描述长度需至少 3 个字",
			"max_cn:描述长度不能超过 255 个字",
		},
		"description": []string{
			"min_cn:描述长度需至少 3 个字",
			"max_cn:描述长度不能超过 255 个字",
		},
		"method": []string{
			"required:请求为必填项",
			"min_cn:描述长度需至少 3 个字",
			"max_cn:描述长度不能超过 20 个字",
		},
	}
	return validate(data, rules, messages)
}

type ApiIDRequest struct {
	ID string `valid:"id" json:"id"`
}

func ApiID(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"id": []string{"numeric", "exists:apis,id"},
	}
	messages := govalidator.MapData{
		"id": []string{
			"numeric:id必须为数字",
			"exists:id必须存在",
		},
	}
	return validate(data, rules, messages)
}
