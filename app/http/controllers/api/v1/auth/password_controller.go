// Package auth 处理用户注册、登录、密码重置
package auth

import (
	v1 "github.com/Gopherlinzy/gin-vue3-admin/app/http/controllers/api/v1"
	"github.com/Gopherlinzy/gin-vue3-admin/app/models/user"
	"github.com/Gopherlinzy/gin-vue3-admin/app/requests"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/response"
	"github.com/gin-gonic/gin"
)

// PasswordController 用户控制器
type PasswordController struct {
	v1.BaseApiController
}

// ResetByPhone 使用手机和验证码重置密码
func (pc *PasswordController) ResetByPhone(c *gin.Context) {
	// 1. 验证表单
	request := requests.ResetByPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.ResetByPhone); !ok {
		return
	}

	// 2. 更新密码
	userModel := user.GetByPhone(request.Phone)
	if userModel.ID == 0 {
		response.Abort404(c)
	} else {
		userModel.Password = request.Password
		userModel.Save()

		response.Success(c)
	}
}

// ResetByEmail 使用手机和验证码重置密码
func (pc *PasswordController) ResetByEmail(c *gin.Context) {
	// 1. 验证表单
	request := requests.ResetByEmailRequest{}
	if ok := requests.Validate(c, &request, requests.ResetByEmail); !ok {
		return
	}

	// 2. 更新密码
	userModel := user.GetByMulti(request.Email)
	if userModel.ID == 0 {
		response.Abort404(c)
	} else {
		userModel.Password = request.Password
		userModel.Save()

		response.Success(c)
	}
}
