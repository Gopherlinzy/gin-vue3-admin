package requests

import (
	"github.com/Gopherlinzy/gin-vue3-admin/app/requests/validators"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/auth"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"mime/multipart"
)

type UserStoreRequest struct {
	Name         string `valid:"name" json:"name"`
	Email        string `json:"email,omitempty" valid:"email"`
	Phone        string `json:"phone,omitempty" valid:"phone"`
	Password     string `valid:"password" json:"password,omitempty"`
	City         string `valid:"city" json:"city"`
	Introduction string `valid:"introduction" json:"introduction"`

	RoleName string `valid:"role_name" json:"role_name,omitempty"`
}

func UserStore(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"name": []string{"required", "alpha_num", "between:3,20", "not_exists:users,name", "not_exists:roles,name"},
		"email": []string{
			"required", "min:4",
			"max:30",
			"email",
			"not_exists:users,email",
		},
		"phone": []string{
			"required",
			"digits:11",
			"not_exists:users,phone",
		},
		"password":     []string{"required", "min:6"},
		"introduction": []string{"min_cn:4", "max_cn:240"},
		"city":         []string{"min_cn:2", "max_cn:20"},
		"role_name":    []string{"required", "between:3,20", "exists:roles,role_name"},
	}
	messages := govalidator.MapData{
		"name": []string{
			"required:用户名为必填项",
			"alpha_num:用户名格式错误，只允许数字和英文",
			"between:用户名长度需在 3~20 之间",
			"not_exists:用户名已被占用",
			"not_exists:用户名与角色名重名",
		},
		"email": []string{
			"required:Email 为必填项",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
			"not_exists:Email 已被占用",
		},
		"phone": []string{
			"required:手机号为必填项，参数名称 phone",
			"digits:手机号长度必须为 11 位的数字",
			"not_exists:Phone 已被占用",
		},
		"password": []string{
			"required:密码为必填项",
			"min:密码长度需大于 6",
		},
		"introduction": []string{
			"min_cn:描述长度需至少 4 个字",
			"max_cn:描述长度不能超过 240 个字",
		},
		"city": []string{
			"min_cn:城市需至少 2 个字",
			"max_cn:城市不能超过 20 个字",
		},
		"role_name": []string{
			"required:角色名为必填项",
			"between:角色名长度需在 3~20 之间",
			"exists:角色名必须存在",
		},
	}
	return validate(data, rules, messages)
}

type UserUpdateRequest struct {
	ID           string `valid:"id" json:"id"`
	Name         string `valid:"name" json:"name"`
	Email        string `json:"email,omitempty" valid:"email"`
	Phone        string `json:"phone,omitempty" valid:"phone"`
	Password     string `valid:"password" json:"password,omitempty"`
	City         string `valid:"city" json:"city"`
	Introduction string `valid:"introduction" json:"introduction"`
	RoleName     string `valid:"role_name" json:"role_name,omitempty"`
}

func UserUpdate(data interface{}, c *gin.Context) map[string][]string {

	// 查询用户名重复时，过滤掉当前用户 ID
	uid := data.(*UserUpdateRequest).ID
	rules := govalidator.MapData{
		"id": []string{"numeric", "exists:users,id"},
		"email": []string{
			"required", "min:4",
			"max:30",
			"email",
			"not_exists:users,email," + uid,
		},
		"phone": []string{
			"required",
			"digits:11",
			"not_exists:users,phone," + uid,
		},
		"password":     []string{"required", "min:6"},
		"name":         []string{"required", "alpha_num", "between:3,20", "not_exists:users,name," + uid},
		"introduction": []string{"min_cn:4", "max_cn:240"},
		"city":         []string{"min_cn:2", "max_cn:20"},
		"role_name":    []string{"required", "between:3,20", "exists:roles,role_name"},
	}
	messages := govalidator.MapData{
		"id": []string{
			"numeric:id必须为数字",
			"exists:id必须存在",
		},
		"name": []string{
			"required:用户名为必填项",
			"alpha_num:用户名格式错误，只允许数字和英文",
			"between:用户名长度需在 3~20 之间",
			"not_exists:用户名已被占用",
		},
		"email": []string{
			"required:Email 为必填项",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
			"not_exists:Email 已被占用",
		},
		"phone": []string{
			"required:手机号为必填项，参数名称 phone",
			"digits:手机号长度必须为 11 位的数字",
			"not_exists:Phone 已被占用",
		},
		"password": []string{
			"required:密码为必填项",
			"min:密码长度需大于 6",
		},
		"introduction": []string{
			"min_cn:描述长度需至少 4 个字",
			"max_cn:描述长度不能超过 240 个字",
		},
		"city": []string{
			"min_cn:城市需至少 2 个字",
			"max_cn:城市不能超过 20 个字",
		},
		"role_name": []string{
			"required:角色名为必填项",
			"between:角色名长度需在 3~20 之间",
			"exists:角色名必须存在",
		},
	}
	return validate(data, rules, messages)
}

type UserUpdateEmailRequest struct {
	ID         string `valid:"id" json:"id"`
	Email      string `json:"email,omitempty" valid:"email"`
	VerifyCode string `json:"verify_code,omitempty" valid:"verify_code"`
}

func UserUpdateEmail(data interface{}, c *gin.Context) map[string][]string {

	currentUser := auth.CurrentUser(c)
	rules := govalidator.MapData{
		"id": []string{"numeric", "exists:users,id"},
		"email": []string{

			"required", "min:4",
			"max:30",
			"email",
			"not_exists:users,email," + currentUser.GetStringID(),
			"not_in:" + currentUser.Email,
		},
		"verify_code": []string{"required", "digits:6"},
	}

	messages := govalidator.MapData{
		"id": []string{
			"numeric:id必须为数字",
			"exists:id必须存在",
		},
		"email": []string{
			"required:Email 为必填项",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
			"not_exists:Email 已被占用",
			"not_in:新的 Email 与老 Email 一致",
		},
		"verify_code": []string{
			"required:验证码答案必填",
			"digits:验证码长度必须为 6 位的数字",
		},
	}

	errs := validate(data, rules, messages)
	_data := data.(*UserUpdateEmailRequest)
	errs = validators.ValidateVerifyCode(_data.Email, _data.VerifyCode, errs)

	return errs
}

type UserUpdatePhoneRequest struct {
	Phone      string `json:"phone,omitempty" valid:"phone"`
	VerifyCode string `json:"verify_code,omitempty" valid:"verify_code"`
}

func UserUpdatePhone(data interface{}, c *gin.Context) map[string][]string {

	currentUser := auth.CurrentUser(c)
	rules := govalidator.MapData{
		"phone": []string{
			"required",
			"digits:11",
			"not_exists:users,phone," + currentUser.GetStringID(),
			"not_in:" + currentUser.Phone,
		},
		"verify_code": []string{"required", "digits:6"},
	}

	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项，参数名称 phone",
			"digits:手机号长度必须为 11 位的数字",
			"not_exists:Phone 已被占用",
			"not_in:新的 Phone 与老 Phone 一致",
		},
		"verify_code": []string{
			"required:验证码答案必填",
			"digits:验证码长度必须为 6 位的数字",
		},
	}

	errs := validate(data, rules, messages)
	_data := data.(*UserUpdatePhoneRequest)
	errs = validators.ValidateVerifyCode(_data.Phone, _data.VerifyCode, errs)

	return errs
}

type UserUpdatePasswordRequest struct {
	Password           string `valid:"password" json:"password,omitempty"`
	NewPassword        string `valid:"new_password" json:"new_password,omitempty"`
	NewPasswordConfirm string `valid:"new_password_confirm" json:"new_password_confirm,omitempty"`
}

func UserUpdatePassword(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"password":             []string{"required", "min:6"},
		"new_password":         []string{"required", "min:6"},
		"new_password_confirm": []string{"required", "min:6"},
	}

	messages := govalidator.MapData{
		"password": []string{
			"required:密码为必填项",
			"min:密码长度需大于 6",
		},
		"new_password": []string{
			"required:密码为必填项",
			"min:密码长度需大于 6",
		},
		"new_password_confirm": []string{
			"required:确认密码框为必填项",
			"min:确认密码长度需大于 6",
		},
	}

	// 确保 comfirm 密码正确
	errs := validate(data, rules, messages)
	_data := data.(*UserUpdatePasswordRequest)
	errs = validators.ValidatePasswordConfirm(_data.NewPassword, _data.NewPasswordConfirm, errs)

	return errs
}

type UserUpdateAvatarRequest struct {
	Avatar *multipart.FileHeader `valid:"avatar" form:"avatar"`
}

func UserUpdateAvatar(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		// size 的单位为 bytes
		// - 1024 bytes 为 1kb
		// - 1048576 bytes 为 1mb
		// - 20971520 bytes 为 20mb
		"file:avatar": []string{"ext:png,jpg,jpeg", "size:20971520", "required"},
	}
	messages := govalidator.MapData{
		"file:avatar": []string{
			"ext:ext头像只能上传 png, jpg, jpeg 任意一种的图片",
			"size:头像文件最大不能超过 20MB",
			"required:必须上传图片",
		},
	}

	return validateFile(c, data, rules, messages)
}

type StoreUserRoleRequest struct {
	ID       string `valid:"id" json:"id"`
	RoleName string `json:"role_name,omitempty" valid:"role_name"`
}

func StoreUserRole(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"id":        []string{"numeric", "exists:users,id"},
		"role_name": []string{"required", "exists:roles,role_name"},
	}

	messages := govalidator.MapData{
		"id": []string{
			"numeric:id必须为数字",
			"exists:id必须存在",
		},
		"role_name": []string{
			"required:角色名为必填项",
			"exists:角色名必须存在",
		},
	}

	return validate(data, rules, messages)
}

type UserIDRequest struct {
	ID string `valid:"id" json:"id"`
}

func UserID(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"id": []string{"numeric", "exists:users,id"},
	}
	messages := govalidator.MapData{
		"id": []string{
			"numeric:id必须为数字",
			"exists:id必须存在",
		},
	}
	return validate(data, rules, messages)
}

type UpdateUserStatusRequest struct {
	ID     string `valid:"id" json:"id"`
	Status string `valid:"status" json:"status"`
}

func UpdateUserStatus(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"id":     []string{"numeric", "exists:users,id"},
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
