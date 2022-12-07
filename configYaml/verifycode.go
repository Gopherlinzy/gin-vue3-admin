package configYaml

type VerifyCode struct {
	CodeLength int `mapstructure:"code_length" json:"code_length" yaml:"code_length"`
	ExpireTime int `mapstructure:"expire_time" json:"expire_time" yaml:"expire_time"`

	DebugExpireTime int    `mapstructure:"debug_expire_time" json:"debug_expire_time" yaml:"debug_expire_time"`
	DebugCode       string `mapstructure:"debug_code" json:"debug_code" yaml:"debug_code"`

	DebugPhonePrefix string `mapstructure:"debug_phone_prefix" json:"debug_phone_prefix" yaml:"debug_phone_prefix"`
	DebugEmailPrefix string `mapstructure:"debug_email_suffix" json:"debug_email_suffix" yaml:"debug_email_suffix"`
}
