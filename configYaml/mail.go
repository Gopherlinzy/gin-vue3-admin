package configYaml

import "github.com/Gopherlinzy/gin-vue3-admin/configYaml/mailhog"

type Mail struct {
	STMP mailhog.STMP `mapstructure:"stmp" json:"stmp" yaml:"stmp"`
	FROM mailhog.FROM `mapstructure:"from" json:"from" yaml:"from"`
}
