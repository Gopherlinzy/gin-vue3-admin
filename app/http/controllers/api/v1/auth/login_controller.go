package auth

import (
	"errors"
	v1 "github.com/Gopherlinzy/gin-vue3-admin/app/http/controllers/api/v1"
	"github.com/Gopherlinzy/gin-vue3-admin/app/models/user"
	"github.com/Gopherlinzy/gin-vue3-admin/app/requests"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/auth"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/jwt"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/response"
	"github.com/gin-gonic/gin"
)

// LoginController 用户控制器
type LoginController struct {
	v1.BaseApiController
}

// LoginByPhone 手机登录
func (lc *LoginController) LoginByPhone(c *gin.Context) {
	// 1. 验证表单
	request := requests.LoginByPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.LoginByPhone); !ok {
		return
	}

	// 2. 尝试登录
	user, err := auth.LoginByPhone(request.Phone)
	if err != nil {
		// 失败，显示错误提示
		response.Error(c, err, "账号不存在")
	} else {
		if !user.Status {
			// 失败，显示错误提示
			response.LoginError(c, errors.New("账号异常"), "账号当前处于冻结状态，登录失败")
			return
		}
		// 登录成功
		token := jwt.NewJWT().IssueToken(user.GetStringID(), user.Name)

		response.JSON(c, gin.H{
			"data":  user,
			"token": token,
		})
	}
}

// LoginByPassword 多种方法登录，支持手机号、email 和用户名
func (lc *LoginController) LoginByPassword(c *gin.Context) {
	request := requests.LoginByPasswordRequest{}
	if ok := requests.Validate(c, &request, requests.LoginByPassword); !ok {
		return
	}

	userModel, err := auth.Attempt(request.LoginID, request.Password)
	if err != nil {
		// 失败，显示错误提示
		response.Error(c, err, "账号不存在或密码错误")
	} else {
		if !userModel.Status {
			// 失败，显示错误提示
			response.LoginError(c, errors.New("账号异常"), "账号当前处于冻结状态，登录失败")
			return
		}
		// 登录成功
		token := jwt.NewJWT().IssueToken(userModel.GetStringID(), userModel.Name)

		response.JSON(c, gin.H{
			"data":        userModel,
			"permissions": user.GetMenus(userModel.Name),
			"token":       token,
		})
	}
}

// RefreshToken 刷新 Access Token
func (lc *LoginController) RefreshToken(c *gin.Context) {

	token, err := jwt.NewJWT().RefreshToken(c)

	if err != nil {
		response.Error(c, err, "令牌刷新失败")
	} else {
		response.JSON(c, gin.H{
			"token": token,
		})
	}
}
