// Package bootstrap 启动程序功能
package bootstrap

import (
	"fmt"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/cache"
	"github.com/Gopherlinzy/gin-vue3-admin/pkg/configYaml"
)

// SetupCache 缓存
func SetupCache() {
	rds := cache.NewRedisStore(
		fmt.Sprintf("%v:%v",
			configYaml.Gohub_Config.Redis.Host,
			configYaml.Gohub_Config.Redis.Port,
		),
		configYaml.Gohub_Config.Redis.UserName,
		configYaml.Gohub_Config.Redis.PassWord,
		configYaml.Gohub_Config.Redis.DatabaseCache,
	)

	cache.InitWithCacheStore(rds)
}
