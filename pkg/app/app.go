// Package app 应用信息
package app

import (
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/configYaml"
	"time"
)

func IsLocal() bool {
	return configYaml.Gohub_Config.App.Env == "local"
}

func IsProduction() bool {
	return configYaml.Gohub_Config.App.Env == "production"
}

func IsTesting() bool {
	return configYaml.Gohub_Config.App.Env == "testing"
}

// TimeNowInTimezone 获取当前时间，支持时区
func TimeNowInTimezone() time.Time {
	chinaTimezone, _ := time.LoadLocation(configYaml.Gohub_Config.App.TimeZone)
	return time.Now().In(chinaTimezone)
}

// URL 传参 path 拼接站点的 URL
func URL(path string) string {
	return configYaml.Gohub_Config.App.Url + path
}

// V1URL 拼接带 v1 标示 URL
func V1URL(path string) string {
	return "/v1/" + path
}
