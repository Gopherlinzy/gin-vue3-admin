// Package verifycode 用以发送手机验证码和邮箱验证码
package verifycode

import (
	"fmt"
	"github.com/Gopherlinzy/gohub/pkg/app"
	"github.com/Gopherlinzy/gohub/pkg/configYaml"
	"github.com/Gopherlinzy/gohub/pkg/helpers"
	"github.com/Gopherlinzy/gohub/pkg/logger"
	"github.com/Gopherlinzy/gohub/pkg/mail"
	"github.com/Gopherlinzy/gohub/pkg/redis"
	"github.com/Gopherlinzy/gohub/pkg/sms"
	"strings"
	"sync"
)

type VerifyCode struct {
	Store Store
}

var once sync.Once
var internalVerifyCode *VerifyCode

func NewVerifyCode() *VerifyCode {
	once.Do(func() {
		internalVerifyCode = &VerifyCode{
			Store: &RedisStore{
				RedisClient: redis.Gohub_Redis,
				KeyPrefix:   configYaml.Gohub_Config.App.Name + ":verifycode:",
			},
		}
	})
	return internalVerifyCode
}

// SendSMS 发送短信验证码，调用示例：
//         verifycode.NewVerifyCode().SendSMS(request.Phone)
func (vc *VerifyCode) SendSMS(phone string) bool {
	// 生成验证码
	code := vc.generateVerifyCode(phone)

	// 方便本地和 API 自动测试
	if !app.IsProduction() && strings.HasPrefix(phone, configYaml.Gohub_Config.VerifyCode.DebugPhonePrefix) {
		return true
	}

	return sms.NewSMS().Send(phone, sms.Message{
		Template: configYaml.Gohub_Config.SMS.Aliyun.TemplateCode,
		Data:     map[string]string{"code": code},
	})
}

// SendEmail 发送邮件验证码，调用示例：
//         verifycode.NewVerifyCode().SendEmail(request.Email)
func (vc *VerifyCode) SendEmail(email string) error {
	// 生成验证码
	code := vc.generateVerifyCode(email)

	// 方便本地和 API 自动测试
	if !app.IsProduction() && strings.HasPrefix(email, configYaml.Gohub_Config.VerifyCode.DebugEmailPrefix) {
		return nil
	}

	content := fmt.Sprintf("<h1>您的 Email 验证码是 %v </h1>", code)
	// 发送邮件
	mail.NewMailer().Send(mail.Email{
		From: mail.From{
			Address: configYaml.Gohub_Config.Mail.FROM.Address,
			Name:    configYaml.Gohub_Config.Mail.FROM.Name,
		},
		To:      []string{email},
		Subject: "Email 验证码",
		HTML:    []byte(content),
	})

	return nil
}

// CheckAnswer 检查用户提交的验证码是否正确，key 可以是手机号或者 Email
func (vc *VerifyCode) CheckAnswer(key string, answer string) bool {

	logger.DebugJSON("验证码", "检查验证码", map[string]string{key: answer})

	// 方便开发，在非生产环境下，具备特殊前缀的手机号和 Email后缀，会直接验证成功
	if !app.IsProduction() &&
		(strings.HasSuffix(key, configYaml.Gohub_Config.VerifyCode.DebugEmailPrefix) ||
			strings.HasPrefix(key, configYaml.Gohub_Config.VerifyCode.DebugPhonePrefix)) {
		return true
	}

	return vc.Store.Verify(key, answer, false)
}

// generateVerifyCode 生成验证码，并放置于 Redis 中
func (vc *VerifyCode) generateVerifyCode(key string) string {
	// 生成随机码
	code := helpers.RandomNumber(configYaml.Gohub_Config.VerifyCode.CodeLength)

	// 为方便开发，本地环境使用固定验证码
	if app.IsLocal() {
		code = configYaml.Gohub_Config.VerifyCode.DebugCode
	}

	logger.DebugJSON("验证码", "生成验证码", map[string]string{key: code})
	// 将验证码及 KEY（邮箱或手机号）存放到 Redis 中并设置过期时间
	vc.Store.Set(key, code)
	return code
}
