package mailhog

type STMP struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
	UserName string `mapstructure:"username" json:"username" yaml:"username"`
	PassWord string `mapstructure:"password" json:"password" yaml:"password"`
}
