package mail

import (
	"fmt"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/configYaml"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/logger"
	emailPKG "github.com/jordan-wright/email"
	"net/smtp"
)

// SMTP 实现 email.Driver interface
type SMTP struct{}

// Send 实现 email.Driver interface 的 Send 方法
func (s *SMTP) Send(email Email, config map[string]string) bool {
	e := emailPKG.NewEmail()

	e.From = fmt.Sprintf("%v <%v>", email.From.Name, email.From.Address)
	e.To = email.To
	e.Bcc = email.Bcc
	e.Cc = email.Cc
	e.Subject = email.Subject
	e.Text = email.Text
	e.HTML = email.HTML

	logger.DebugJSON("发送邮件", "发件详情", e)

	err := e.Send(
		fmt.Sprintf("%v:%v", config["host"], config["port"]),
		smtp.PlainAuth(
			"",
			configYaml.Gohub_Config.Mail.STMP.UserName,
			configYaml.Gohub_Config.Mail.STMP.PassWord,
			configYaml.Gohub_Config.Mail.STMP.Host,
		),
	)

	if err != nil {
		logger.ErrorString("发送邮件", "发送出错", err.Error())
		return false
	}

	logger.DebugString("发送邮件", "发送成功", "")
	return true
}
