package mailhog

type FROM struct {
	Address string `mapstructure:"address" json:"address" yaml:"address"`
	Name    string `mapstructure:"name" json:"name" yaml:"name"`
}
