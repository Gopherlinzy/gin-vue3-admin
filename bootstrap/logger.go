package bootstrap

import (
	"github.com/Gopherlinzy/gohub/pkg/configYaml"
	"github.com/Gopherlinzy/gohub/pkg/logger"
)

// SetupLogger 初始化 Logger
func SetupLogger() {
	logger.InitLogger(
		configYaml.Gohub_Config.Log.FileName,
		configYaml.Gohub_Config.Log.MaxSize,
		configYaml.Gohub_Config.Log.MaxBackup,
		configYaml.Gohub_Config.Log.MaxAge,
		configYaml.Gohub_Config.Log.Compress,
		configYaml.Gohub_Config.Log.Type,
		configYaml.Gohub_Config.Log.Level,
	)
}
