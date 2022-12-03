package configYaml

import (
	"fmt"
	"github.com/Gopherlinzy/gohub/configYaml"
	"github.com/fsnotify/fsnotify"
	viperlib "github.com/spf13/viper"
	"os"
)

var viper *viperlib.Viper

var Gohub_Config configYaml.Server

func init() {
	// 1. 初始化 Viper 库
	viper = viperlib.New()
	// 2. 配置类型，支持 "json", "toml", "yaml", "yml", "properties",
	//             "props", "prop", "env", "dotenv"
	viper.SetConfigType("yaml")
	// 3. 环境变量配置文件查找的路径，相对于 main.go
	viper.AddConfigPath(".")
	// 4. 设置环境变量前缀，用以区分 Go 的系统环境变量
	viper.SetEnvPrefix("appenv")
	// 5. 读取环境变量（支持 flags）
	viper.AutomaticEnv()
}

// InitConfig 初始化配置信息, 完成对环境变量以及 config 信息的加载
func InitConfig(yaml string) {
	// 1. 加载环境变量
	loadYaml(yaml)
	// 2. 注册配置信息
	loadConfig()
}

func loadYaml(yamSuffix string) {
	// 默认加载 .yaml 文件，如果有传参 --yaml=name 的话，加载 .yaml.name 文件
	yamlPath := "config.yaml"
	if len(yamSuffix) > 0 {
		filePath := yamlPath + "." + yamSuffix
		if _, err := os.Stat(filePath); err != nil {
			// 如 .yaml.testing 或 .yaml.stage
			yamlPath = filePath
		}
	}

	// 加载 yaml
	viper.SetConfigName(yamlPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// 监控 .yaml 文件,变更时加载
	viper.WatchConfig()
}

func loadConfig() {
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)
		if err := viper.Unmarshal(&Gohub_Config); err != nil {
			fmt.Println(err)
		}
	})
	if err := viper.Unmarshal(&Gohub_Config); err != nil {
		fmt.Println(err)
	}
}
