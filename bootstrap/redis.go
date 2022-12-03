package bootstrap

import (
	"fmt"
	"github.com/Gopherlinzy/gohub/pkg/configYaml"
	"github.com/Gopherlinzy/gohub/pkg/redis"
)

// SetupRedis 初始化 Redis
func SetupRedis() {
	redis.ConnectRedis(
		fmt.Sprintf("%v:%v", configYaml.Gohub_Config.Redis.Host,
			configYaml.Gohub_Config.Redis.Port),
		configYaml.Gohub_Config.Redis.UserName,
		configYaml.Gohub_Config.Redis.PassWord,
		configYaml.Gohub_Config.Redis.Database,
	)
}
