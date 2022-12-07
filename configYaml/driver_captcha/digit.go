package driver_captcha

type Digit struct {
	Maxskew  float64 `mapstructure:"maxskew" json:"maxskew" yaml:"maxskew"`
	DotCount int     `mapstructure:"dotcount" json:"dotcount" yaml:"dotcount"`
}
