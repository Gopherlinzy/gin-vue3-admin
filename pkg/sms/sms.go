package sms

import (
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/configYaml"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/helpers"
	"sync"
)

// Message 是短信结构体
type Message struct {
	Template string
	Data     map[string]string

	Content string
}

// SMS 是短信操作类
type SMS struct {
	Driver Driver
}

var once sync.Once

// internalSMS 内部使用的 SMS 对象
var internalSMS *SMS

// NewSMS 单例模式获取
func NewSMS() *SMS {
	once.Do(func() {
		internalSMS = &SMS{
			Driver: &Aliyun{},
		}
	})
	return internalSMS
}

func (sms *SMS) Send(phone string, message Message) bool {
	return sms.Driver.Send(phone, message, helpers.StructToMap(&configYaml.Gohub_Config.SMS.Aliyun))
}
