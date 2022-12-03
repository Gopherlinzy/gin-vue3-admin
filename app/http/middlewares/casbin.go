package middlewares

import (
	casbins "github.com/Gopherlinzy/gohub/pkg/casbin"
	"github.com/Gopherlinzy/gohub/pkg/response"
	"github.com/gin-gonic/gin"
)

// CasbinAPI 在用户登陆后判断他是否拥有权限
func CasbinAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		sub, exists := c.MustGet("current_user_name").(string)
		if exists {
			//sub := casbins.NewCasbin().GetRolesForUser(userName)[0]sub := casbins.NewCasbin().GetRolesForUser(userName)[0]
			// 获取请求的PATH
			obj := c.Request.URL.Path
			// 获取请求方法
			act := c.Request.Method
			//fmt.Println("--------", sub, obj, act)

			success := casbins.NewCasbin().Enforce(sub, obj, act)
			if !success {
				response.NoPolicyRequest(c)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
