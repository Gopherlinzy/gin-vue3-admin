package sms

import (
	"encoding/json"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/configYaml"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/logger"
	aliyunsmsclient "github.com/KenmyZhang/aliyun-communicate"
)

// Aliyun 实现 sms.Driver interface
type Aliyun struct {
}

// Send 实现 sms.Driver interface 的 Send 方法
func (s *Aliyun) Send(phone string, message Message, config map[string]string) bool {
	smsClient := aliyunsmsclient.New("http://dysmsapi.aliyuncs.com/")

	templateParam, err := json.Marshal(message.Data)
	if err != nil {
		logger.ErrorString("短信[阿里云]", "解析绑定错误", err.Error())
		return false
	}

	logger.DebugJSON("短信[阿里云]", "配置信息", config)

	result, err := smsClient.Execute(
		configYaml.Gohub_Config.SMS.Aliyun.AccessKeyId,
		configYaml.Gohub_Config.SMS.Aliyun.AccessKeySecret,
		phone,
		configYaml.Gohub_Config.SMS.Aliyun.SignName,
		message.Template,
		string(templateParam),
	)

	logger.DebugJSON("短信[阿里云]", "发送失败", smsClient.Request)
	logger.DebugJSON("短信[阿里云]", "接口响应", result)

	if err != nil {
		logger.ErrorString("短信[阿里云]", "发信失败", err.Error())
		return false
	}

	resultJSON, err := json.Marshal(result)
	if err != nil {
		logger.ErrorString("短信[阿里云]", "解析响应 JSON 错误", err.Error())
		return false
	}

	if result.IsSuccessful() {
		logger.DebugString("短信[阿里云]", "发信成功", "")
		return true
	} else {
		logger.ErrorString("短信[阿里云]", "服务商返回错误", string(resultJSON))
		return false
	}
}
