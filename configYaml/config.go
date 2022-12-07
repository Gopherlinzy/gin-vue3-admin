package configYaml

type Server struct {
	App   System `mapstructure:"app" json:"app" yaml:"app"`
	MySQL MySQL  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Redis Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`
	JWT   JWT    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`

	VerifyCode VerifyCode `mapstructure:"verifycode" json:"verifycode" yaml:"verifycode"`
	Captcha    Captcha    `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	SMS        SMS        `mapstructure:"sms" json:"sms" yaml:"sms"`
	Mail       Mail       `mapstructure:"mail" json:"mail" yaml:"mail"`

	Paging Paging `mapstructure:"paging" json:"paging" yaml:"paging"`
	Log    Log    `mapstructure:"log" json:"log" yaml:"log"`
}
