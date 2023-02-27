package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"go-web/web_app2/settings"
)

var rdb *redis.Client

// 初始化链接
func Init(cfg *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize,
	})
	_, err = rdb.Ping().Result()
	return err
}

func Close() {
	_ = rdb.Close()
}
