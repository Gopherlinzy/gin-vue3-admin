package mail

type Driver interface {
	// 检查验证码
	Send(mail Email, config map[string]string) bool
}
