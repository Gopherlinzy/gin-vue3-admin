package configYaml

import "github.com/Gopherlinzy/gohub/configYaml/mailhog"

type Mail struct {
	STMP mailhog.STMP `mapstructure:"stmp" json:"stmp" yaml:"stmp"`
	FROM mailhog.FROM `mapstructure:"from" json:"from" yaml:"from"`
}
