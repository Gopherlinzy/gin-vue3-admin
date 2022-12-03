package configYaml

type GeneralDB struct {
	// 服务器地址:端口
	Host string `mapstructure:"host" json:"host" yaml:"host"`
	// :端口
	Port string `mapstructure:"port" json:"port" yaml:"port"`
	// 数据库名
	Database string `mapstructure:"database" json:"database" yaml:"database"`
	// 数据库用户名
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	// 数据库密码
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Charset  string `mapstructure:"charset" json:"charset" yaml:"charset"`

	// 空闲中的最大连接数
	MaxIdleConections int `mapstructure:"max_idle_connections" json:"max_idle_connections" yaml:"max_idle_connections"`
	// 打开到数据库的最大连接数
	MaxOpenConnections int `mapstructure:"max_open_connections" json:"max_open_connections" yaml:"max_open_connections"`
	MaxLifeSeconds     int `mapstructure:"max_life_seconds" json:"max_life_seconds" yaml:"max_life_seconds"`
}
