package verifycode

import (
	"github.com/Gopherlinzy/gohub/pkg/app"
	"github.com/Gopherlinzy/gohub/pkg/configYaml"
	"github.com/Gopherlinzy/gohub/pkg/redis"
	"time"
)

// RedisStore 实现 verifycode.Store interface
type RedisStore struct {
	RedisClient *redis.RedisClient
	KeyPrefix   string
}

func (s *RedisStore) Set(key string, value string) bool {
	expireTime := time.Minute * time.Duration(configYaml.Gohub_Config.VerifyCode.ExpireTime)

	// 方便本地环境测试
	if app.IsLocal() {
		expireTime = time.Minute * time.Duration(configYaml.Gohub_Config.VerifyCode.DebugExpireTime)
	}

	return s.RedisClient.Set(s.KeyPrefix+key, value, expireTime)
}

func (s *RedisStore) Get(key string, clear bool) string {
	key = s.KeyPrefix + key
	val := s.RedisClient.Get(key)
	if clear {
		s.RedisClient.Del(key)
	}
	return val
}

func (s *RedisStore) Verify(key, answer string, clear bool) bool {
	v := s.Get(key, clear)
	return v == answer
}
