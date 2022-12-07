// Package routes 注册路由
package routes

import (
	controllers "github.com/Gopherlinzy/gin-vue3-admin/app/http/controllers/api/v1"
	"github.com/Gopherlinzy/gin-vue3-admin/app/http/controllers/api/v1/auth"
	"github.com/Gopherlinzy/gin-vue3-admin/app/http/middlewares"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/configYaml"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RegisterAPIRoutes 注册 API 相关路由
func RegisterAPIRoutes(r *gin.Engine) {

	var v1 *gin.RouterGroup
	if len(configYaml.Gohub_Config.App.APIDomain) == 0 {
		configYaml.Gohub_Config.App.Url += "/api"
		v1 = r.Group("/api/v1")
	} else {
		v1 = r.Group("/v1")
	}
	// 设置静态路径，可以通过url访问图片
	v1.StaticFS("uploads", http.Dir("./public/uploads"))

	// 全局限流中间件：每小时限流。这里是所有 API （根据 IP）请求加起来。
	// 作为参考 Github API 每小时最多 60 个请求（根据 IP）。
	// 测试时，可以调高一点。
	v1.Use(middlewares.LimitIP("1000-H"))
	{
		authGroup := v1.Group("/auth")
		// 限流中间件：每小时限流，作为参考 Github API 每小时最多 60 个请求（根据 IP）
		// 测试时，可以调高一点
		authGroup.Use(middlewares.LimitIP("500-H"))
		{
			// 登录
			lgc := new(auth.LoginController)
			lgcGroup := authGroup.Group("/login")
			{
				lgcGroup.POST("/using-phone", middlewares.GuestJWT(), lgc.LoginByPhone)
				lgcGroup.POST("/using-password", middlewares.GuestJWT(), lgc.LoginByPassword)
				lgcGroup.POST("/refresh-token", middlewares.AuthJWT(), middlewares.CasbinAPI(), lgc.RefreshToken)
			}

			// 重置密码
			pwc := new(auth.PasswordController)
			pwcGroup := authGroup.Group("/password-reset", middlewares.GuestJWT())
			{
				pwcGroup.POST("/using-email", pwc.ResetByEmail)
				pwcGroup.POST("/using-phone", pwc.ResetByPhone)
			}

			// 注册用户
			suc := new(auth.SignupController)
			sucGroup := authGroup.Group("/signup")
			{
				sucGroup.POST("/using-phone", suc.SignupUsingPhone)
				sucGroup.POST("/using-email", suc.SignupUsingEmail)
				sucGroup.POST("/phone/exist", middlewares.LimitPerRoute("60-H"), suc.IsPhoneExist)
				sucGroup.POST("/email/exist", middlewares.LimitPerRoute("60-H"), suc.IsEmailExist)

			}

			// 发送验证码
			vcc := new(auth.VerifyCodeController)
			vccGroup := authGroup.Group("/verify-codes")
			{
				vccGroup.POST("/phone", middlewares.LimitPerRoute("20-H"), vcc.SendUsingPhone)
				vccGroup.POST("/email", middlewares.LimitPerRoute("20-H"), vcc.SendUsingEmail)
				// 图片验证码
				vccGroup.GET("/captcha", middlewares.LimitPerRoute("50-H"), vcc.ShowCaptcha)

			}

			// users CRUD接口
			uc := new(controllers.UsersController)
			// 获取当前用户
			v1.GET("/user", middlewares.AuthJWT(), uc.CurrentUser)
			usersGroup := v1.Group("/users", middlewares.AuthJWT(), middlewares.CasbinAPI())
			{
				usersGroup.GET("", uc.Index)
				usersGroup.POST("", uc.Store)
				usersGroup.POST("/id", uc.GetUser)
				usersGroup.POST("/reset", uc.ResetPassword)
				usersGroup.PUT("", uc.Update)
				usersGroup.PUT("/status", uc.UpdateUserStatus)
				usersGroup.PUT("/role", uc.UpdateUserRole)
				usersGroup.PUT("/email", uc.UpdateEmail)
				usersGroup.PUT("/phone", uc.UpdatePhone)
				usersGroup.PUT("/password", uc.UpdatePassword)
				usersGroup.PUT("/avatar", uc.UpdateAvatar)
				usersGroup.DELETE("", uc.DeleteUser)
			}

			// category CRUD接口
			cgc := new(controllers.CategoriesController)
			cgcGroup := v1.Group("/categories", middlewares.AuthJWT(), middlewares.CasbinAPI())
			{
				cgcGroup.GET("", cgc.Index)
				cgcGroup.POST("", cgc.Store)
				cgcGroup.PUT("", cgc.Update)
				cgcGroup.DELETE("", cgc.Delete)
			}

			// topic CRUD接口
			tpc := new(controllers.TopicsController)
			tpcGroup := v1.Group("/topics", middlewares.AuthJWT())
			{
				tpcGroup.GET("", middlewares.CasbinAPI(), tpc.Index)
				tpcGroup.POST("/id", tpc.Show)
				tpcGroup.POST("", middlewares.CasbinAPI(), tpc.Store)
				tpcGroup.PUT("", middlewares.CasbinAPI(), tpc.Update)
				tpcGroup.DELETE("", middlewares.CasbinAPI(), tpc.Delete)
			}

			// link 接口
			lsc := new(controllers.LinksController)
			lscGroup := v1.Group("/links")
			{
				lscGroup.GET("", lsc.Index)
			}

			// role 接口
			rsc := new(controllers.RolesController)
			rscGroup := v1.Group("/roles", middlewares.AuthJWT(), middlewares.CasbinAPI())
			{
				rscGroup.GET("", rsc.Index)
				rscGroup.POST("", rsc.Store)
				rscGroup.POST("/id", rsc.Show)
				rscGroup.PUT("", rsc.Update)
				rscGroup.PUT("/status", rsc.UpdateRoleStatus)
				rscGroup.DELETE("", rsc.Delete)
			}
		}
	}
}
