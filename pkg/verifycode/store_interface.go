package verifycode

type Store interface {
	// Set 保存验证码
	Set(id string, value string) bool

	// Get 获取验证码
	Get(id string, clear bool) string

	// Verify 检查验证码
	Verify(id, answer string, clear bool) bool
}
