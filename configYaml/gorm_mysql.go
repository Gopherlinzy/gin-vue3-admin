package configYaml

type MySQL struct {
	GeneralDB `yaml:",inline" mapstructure:",squash"`
}
