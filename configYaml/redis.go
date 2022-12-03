package configYaml

type Redis struct {
	Host          string `mapstructure:"host" json:"host" yaml:"host"`
	Port          string `mapstructure:"port" json:"port" yaml:"port"`
	UserName      string `mapstructure:"username" json:"username" yaml:"username"`
	PassWord      string `mapstructure:"password" json:"password" yaml:"password"`
	Database      int    `mapstructure:"database" json:"database" yaml:"database"`
	DatabaseCache int    `mapstructure:"database_cache" json:"database_cache" yaml:"database_cache"`
}
