// Package requests 处理请求数据和表单验证
package requests

import (
	"fmt"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"net/http"
)

// ValidatorFunc 验证函数类型
type ValidatorFunc func(interface{}, *gin.Context) map[string][]string

// Validate 控制器里调用示例：
//
//	if ok := requests.Validate(c, &requests.UserSaveRequest{}, requests.UserSave); !ok {
//	    return
//	}
func Validate(c *gin.Context, request interface{}, handler ValidatorFunc) bool {
	// 1. 解析请求，支持 JSON 数据、表单请求和 URL Query
	if err := c.ShouldBind(request); err != nil {
		// 解释失败，返回 422 状态码和错误信息
		response.BadRequest(c, err,
			"请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。")
		// 打印错误信息
		fmt.Println(err.Error())
		// 出错了, 中断请求
		return false
	}

	// 表单验证 判断验证是否通过
	if errs := handler(request, c); len(errs) > 0 {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "请求验证不通过，具体请查看 errors",
			"errors":  errs,
		})
		return false
	}
	return true
}

func validate(data interface{}, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {
	// 配置初始化
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		TagIdentifier: "valid", // 模型中的 Struct 标签标识符
		Messages:      messages,
	}

	// 开始验证
	return govalidator.New(opts).ValidateStruct()
}

func validateFile(c *gin.Context, data interface{}, rules govalidator.MapData,
	messages govalidator.MapData) map[string][]string {
	opts := govalidator.Options{
		Request:       c.Request,
		Rules:         rules,
		Messages:      messages,
		TagIdentifier: "valid",
	}
	// 调用 govalidator 的 Validate 方法来验证文件
	return govalidator.New(opts).Validate()
}
