package configYaml

type JWT struct {
	ExpireTime      int64 `mapstructure:"expire_time" json:"expire_time" yaml:"expire_time"`
	MaxRefreshTime  int64 `mapstructure:"max_refresh_time" json:"max_refresh_time" yaml:"max_refresh_time"`
	DebugExpireTime int64 `mapstructure:"debug_expire_time" json:"debug_expire_time" yaml:"debug_expire_time"`
}
