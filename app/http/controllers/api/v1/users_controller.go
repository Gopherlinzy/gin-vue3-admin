package v1

import (
	"fmt"
	"github.com/Gopherlinzy/gohub/app/models/user"
	"github.com/Gopherlinzy/gohub/app/requests"
	"github.com/Gopherlinzy/gohub/pkg/auth"
	"github.com/Gopherlinzy/gohub/pkg/file"
	"github.com/Gopherlinzy/gohub/pkg/helpers"
	"github.com/Gopherlinzy/gohub/pkg/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UsersController struct {
	BaseApiController
}

// CurrentUser 当前登录用户信息
func (ctrl *UsersController) CurrentUser(c *gin.Context) {
	userModel := auth.CurrentUser(c)
	response.Data(c, userModel)
}

// Index 所有用户
func (ctrl *UsersController) Index(c *gin.Context) {
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	data, pager := user.Paginate(c, 5)
	response.JSON(c, gin.H{
		"data":  data,
		"pager": pager,
	})
}

func (ctrl *UsersController) Store(c *gin.Context) {
	request := requests.UserStoreRequest{}
	if ok := requests.Validate(c, &request, requests.UserStore); !ok {
		return
	}

	status, _ := strconv.ParseBool(request.Status)

	userModel := user.User{
		Name:         request.Name,
		Email:        request.Email,
		Phone:        request.Phone,
		Password:     request.Password,
		City:         request.City,
		Introduction: request.Introduction,
		Status:       status,
	}

	userModel.Create()
	if userModel.ID > 0 {
		response.Created(c, userModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *UsersController) Update(c *gin.Context) {
	request := requests.UserUpdateRequest{}
	if ok := requests.Validate(c, &request, requests.UserUpdate); !ok {
		return
	}

	status, _ := strconv.ParseBool(request.Status)

	currentUser := user.Get(request.ID)
	currentUser.Name = request.Name
	currentUser.Email = request.Email
	currentUser.Phone = request.Phone
	currentUser.Password = request.Password
	currentUser.City = helpers.IsNull(currentUser.City, request.City)
	currentUser.Introduction = helpers.IsNull(currentUser.Introduction, request.Introduction)
	currentUser.Status = status

	rowsAffected := currentUser.Save()
	if rowsAffected > 0 {
		response.Data(c, currentUser)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}

}

func (ctrl *UsersController) UpdateEmail(c *gin.Context) {
	request := requests.UserUpdateEmailRequest{}
	if ok := requests.Validate(c, &request, requests.UserUpdateEmail); !ok {
		return
	}

	currentUser := auth.CurrentUser(c)
	currentUser.Email = request.Email
	rowsAffected := currentUser.Save()

	if rowsAffected > 0 {
		response.Success(c)
	} else {
		// 失败，显示错误提示
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *UsersController) UpdatePhone(c *gin.Context) {
	request := requests.UserUpdatePhoneRequest{}
	if ok := requests.Validate(c, &request, requests.UserUpdatePhone); !ok {
		return
	}

	currentUser := auth.CurrentUser(c)
	currentUser.Phone = request.Phone
	rowsAffected := currentUser.Save()

	if rowsAffected > 0 {
		response.Success(c)
	} else {
		// 失败，显示错误提示
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *UsersController) UpdatePassword(c *gin.Context) {
	request := requests.UserUpdatePasswordRequest{}
	if ok := requests.Validate(c, &request, requests.UserUpdatePassword); !ok {
		return
	}

	currentUser := auth.CurrentUser(c)
	// 验证原始密码是否正确
	_, err := auth.Attempt(currentUser.Name, request.Password)
	if err != nil {
		// 失败，显示错误提示
		response.Unauthorized(c, "原密码不正确")
	} else {
		// 更新密码为新密码
		currentUser.Password = request.NewPassword
		currentUser.Save()

		response.Success(c)
	}
}

// UpdateAvatar 上传头像
func (ctrl *UsersController) UpdateAvatar(c *gin.Context) {
	request := requests.UserUpdateAvatarRequest{}
	if ok := requests.Validate(c, &request, requests.UserUpdateAvatar); !ok {
		return
	}
	avatar, err := file.SaveUploadAvatar(c, request.Avatar)
	if err != nil {
		response.Abort500(c, "上传头像失败，请稍后尝试~")
		return
	}

	currentUser := auth.CurrentUser(c)
	currentUser.Avatar = avatar
	currentUser.Save()

	response.Data(c, currentUser)
}

// StoreUserRole 给用户添加角色
func (ctrl *UsersController) StoreUserRole(c *gin.Context) {

	request := requests.StoreUserRoleRequest{}
	if ok := requests.Validate(c, &request, requests.StoreUserRole); !ok {
		return
	}

	userModel := user.Get(request.ID)

	userRoleModel := user.UserRole{
		Name:     userModel.Name,
		RoleName: request.RoleName,
	}

	err := userRoleModel.UpdateRole()
	if err != nil {
		response.Abort500(c, "创建失败，请稍后尝试~")
	} else {
		response.Created(c, userModel)
	}
}

func (ctrl *UsersController) DeleteUser(c *gin.Context) {

	// 表单验证
	request := requests.UserDeleteResetRequest{}

	if bindOk := requests.Validate(c, &request, requests.UserDeleteReset); !bindOk {
		return
	}
	fmt.Println(request)

	userModel := user.Get(request.ID)

	rowsAffected := userModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "删除失败，请稍后尝试~")
}

// ResetPassword 重置密码为 123456
func (ctrl *UsersController) ResetPassword(c *gin.Context) {
	// 表单验证
	request := requests.UserDeleteResetRequest{}

	if bindOk := requests.Validate(c, &request, requests.UserDeleteReset); !bindOk {
		return
	}

	userModel := user.Get(request.ID)
	fmt.Println(userModel)

	userModel.Password = "123456"
	rowsAffected := userModel.Save()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "重置密码失败，请稍后尝试~")
}
