package configYaml

type Log struct {
	Level     string `mapstructure:"level" json:"level" yaml:"level"`
	Type      string `mapstructure:"type" json:"type" yaml:"type"`
	FileName  string `mapstructure:"filename" json:"filename" yaml:"filename"`
	MaxSize   int    `mapstructure:"max_size" json:"max_size" yaml:"max_size"`
	MaxBackup int    `mapstructure:"max_backup" json:"max_backup" yaml:"max_backup"`
	MaxAge    int    `mapstructure:"max_age" json:"max_age" yaml:"max_age"`
	Compress  bool   `mapstructure:"compress" json:"compress" yaml:"compress"`
}
