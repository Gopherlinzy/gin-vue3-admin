package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type MenuSaveRequest struct {
	ID string `valid:"id" json:"id"`

	Name       string `json:"name" valid:"name"`
	Permission string `json:"permission" valid:"permission"`
	RouterName string `json:"router_name" valid:"router_name"`
	RouterPath string `json:"router_path" valid:"router_path"`
	FatherID   string `json:"father_id,omitempty" valid:"father_id"`
	VuePath    string `json:"vue_path,omitempty" valid:"vue_path"`
	Status     string `valid:"status" json:"status"`
}

func MenuSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"id":          []string{"numeric", "exists:menus,id"},
		"name":        []string{"required", "not_exists:menus,name"},
		"permission":  []string{"required", "not_exists:menus,permission"},
		"router_name": []string{"required"},
		"router_path": []string{"required"},
		"father_id":   []string{"numeric", "exists:menus,id,0"},
		"vue_path":    []string{"required"},
		"status":      []string{"required", "bool"},
	}
	messages := govalidator.MapData{
		"id": []string{
			"numeric:id必须为数字",
			"exists:id必须存在",
		},
		"name": []string{
			"required:名称为必填项",
			"not_exists:名称已存在",
		},
		"permission": []string{
			"required:权限为必填项",
			"not_exists:权限已存在",
		},
		"router_name": []string{
			"required:路由名称为必填项",
		},
		"router_path": []string{
			"required:路由路径为必填项",
		},
		"father_id": []string{
			"numeric:父节点必须为数字",
			"exists:父节点必须是存在的id",
		},
		"vue_path": []string{
			"required:文件路径为必填项",
		},
		"status": []string{
			"required:status为必填项",
			"bool:必须得为布尔类型:true, false, 1, 0, \"1\" and \"0\"",
		},
	}
	return validate(data, rules, messages)
}

type MenuIDRequest struct {
	ID string `valid:"id" json:"id"`
}

func MenuID(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"id": []string{"numeric", "exists:menus,id"},
	}
	messages := govalidator.MapData{
		"id": []string{
			"numeric:id必须为数字",
			"exists:id必须存在",
		},
	}
	return validate(data, rules, messages)
}
