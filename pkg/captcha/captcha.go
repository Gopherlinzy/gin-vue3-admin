// Package captcha 处理图片验证码逻辑
package captcha

import (
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/app"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/configYaml"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/helpers"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/redis"
	"github.com/mojocn/base64Captcha"
	"image/color"
	"sync"
)

type Captcha struct {
	Base64Captcha *base64Captcha.Captcha
}

// once 确保 internalCaptcha 对象只初始化一次
var once sync.Once

// internalCaptcha 内部使用的 Captcha 对象
var internalCaptcha *Captcha

// NewCaptcha 单例模式获取
func NewCaptcha() *Captcha {
	once.Do(func() {
		// 初始化 Captcha 对象
		internalCaptcha = &Captcha{}

		// 使用全局 Redis 对象, 并配置存储 Key 的前缀
		store := RedisStore{
			RedisClient: redis.Gohub_Redis,
			KeyPrefix:   configYaml.Gohub_Config.App.Name + ":captcha:",
		}

		x := helpers.RandomInt(2) + 1
		switch x {
		case 1:
			driverDigit := newDriverDigit(
				configYaml.Gohub_Config.Captcha.Height,
				configYaml.Gohub_Config.Captcha.Width,
				configYaml.Gohub_Config.Captcha.Length,
				configYaml.Gohub_Config.Captcha.Digit.Maxskew,
				configYaml.Gohub_Config.Captcha.Digit.DotCount,
			)
			// 实例化 base64captcha, 并赋值给内部使用的 internalCaptcha 对象
			internalCaptcha.Base64Captcha = base64Captcha.NewCaptcha(driverDigit, &store)
		case 2:
			driverString := newDriverString(
				configYaml.Gohub_Config.Captcha.Height,
				configYaml.Gohub_Config.Captcha.Width,
				helpers.RandomInt(20),
				helpers.RandomInt(8),
				configYaml.Gohub_Config.Captcha.Length,
				configYaml.Gohub_Config.Captcha.String.Source,
				&color.RGBA{
					R: uint8(helpers.RandomInt(255)),
					G: uint8(helpers.RandomInt(255)),
					B: uint8(helpers.RandomInt(255)),
					A: uint8(helpers.RandomInt(255)),
				},
				nil,
				nil,
			)
			// 实例化 base64captcha, 并赋值给内部使用的 internalCaptcha 对象
			internalCaptcha.Base64Captcha = base64Captcha.NewCaptcha(driverString, &store)
		}
	})
	return internalCaptcha
}

func newDriverDigit(height int, width int, length int, maxSkew float64, dotCount int) *base64Captcha.DriverDigit {
	return base64Captcha.NewDriverDigit(height, width, length, maxSkew, dotCount)
}

func newDriverString(height int, width int, noiseCount int, showLineOptions int, length int, source string,
	bgColor *color.RGBA, fontsStorage base64Captcha.FontsStorage, fonts []string) *base64Captcha.DriverString {
	return base64Captcha.NewDriverString(height, width, noiseCount, showLineOptions, length, source, bgColor, fontsStorage, fonts)
}

// GenerateCaptcha 生成图片验证码
func (c *Captcha) GenerateCaptcha() (id string, b64s string, err error) {
	return c.Base64Captcha.Generate()
}

// VerifyCaptcha 验证验证码是否正确
func (c *Captcha) VerifyCaptcha(id string, answer string) (match bool) {
	// 方便本地和 API 测试
	if !app.IsProduction() && id == configYaml.Gohub_Config.Captcha.TestingKey {
		return true
	}
	// 第三个参数是验证后是否删除，我们选择 false
	// 这样方便用户多次提交，防止表单提交错误需要多次输入图片验证码
	return c.Base64Captcha.Verify(id, answer, false)
}
