package configYaml

import "github.com/Gopherlinzy/gin-vue3-admin/configYaml/driver_captcha"

type Captcha struct {
	Height          int    `mapstructure:"height" json:"height" yaml:"height"`
	Width           int    `mapstructure:"width" json:"width" yaml:"width"`
	Length          int    `mapstructure:"length" json:"length" yaml:"length"`
	ExpireTime      int    `mapstructure:"expire_time" json:"expire_time" yaml:"expire_time"`
	DebugExpireTime int    `mapstructure:"debug_expire_time" json:"debug_expire_time" yaml:"debug_expire_time"`
	TestingKey      string `mapstructure:"testing_key" json:"testing_key" yaml:"testing_key"`

	Digit  driver_captcha.Digit  `mapstructure:"digit" json:"digit" yaml:"digit"`
	String driver_captcha.String `mapstructure:"string" json:"string" yaml:"string"`
}
