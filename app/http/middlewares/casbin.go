package middlewares

import (
	"github.com/Gopherlinzy/gin-vue3-admin/app/models/role"
	casbins "github.com/Gopherlinzy/gin-vue3-admin/pkg/casbin"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/response"
	"github.com/gin-gonic/gin"
)

// CasbinAPI 在用户登陆后判断他是否拥有权限
func CasbinAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		sub, exists := c.MustGet("current_user_name").(string)
		if exists {
			r := casbins.NewCasbin().GetRolesForUser(sub)[0]
			if r == "" {
				response.NoPolicyRequest(c, "你没有权限")
				c.Abort()
			}
			// 获取请求的PATH
			obj := c.Request.URL.Path
			// 获取请求方法
			act := c.Request.Method
			//fmt.Println("--------", sub, obj, act)

			// 存在这条policy
			success := casbins.NewCasbin().Enforce(r, obj, act)
			// 并且角色状态为true
			status := role.GetBy("role_name", r).Status
			if !success || !status {
				response.NoPolicyRequest(c)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
