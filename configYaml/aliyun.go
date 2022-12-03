package configYaml

type Aliyun struct {
	AccessKeyId     string `mapstructure:"access_key_id" json:"access_key_id" yaml:"access_key_id"`
	AccessKeySecret string `mapstructure:"access_key_secret" json:"access_key_secret" yaml:"access_key_secret"`
	SignName        string `mapstructure:"sign_name" json:"sign_name" yaml:"sign_name"`
	TemplateCode    string `mapstructure:"template_code" json:"template_code" yaml:"template_code"`
}
