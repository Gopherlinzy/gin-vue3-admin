package configYaml

type System struct {
	// 应用名称
	Name string `mapstructure:"name" json:"name" yaml:"name"`
	// 当前环境，用以区分多环境，一般为 local, stage, production, test
	Env string `mapstructure:"env" json:"env" yaml:"env"`
	// 加密会话、JWT 加密
	Key string `mapstructure:"key" json:"key" yaml:"key"`
	// 是否进入调试模式
	Debug bool `mapstructure:"debug" json:"debug" yaml:"debug"`
	// 用以生成链接
	Url string `mapstructure:"url" json:"url" yaml:"url"`
	// 应用服务端口
	Port string `mapstructure:"port" json:"port" yaml:"port"`
	// 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
	DbType string `mapstructure:"db-type" json:"db-type" yaml:"db-type"`

	TimeZone  string `mapstructure:"timezone" json:"timezone" yaml:"timezone"`
	APIDomain string `mapstructure:"api_domain" json:"api_domain" yaml:"api_domain"`
}
