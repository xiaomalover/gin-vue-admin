package initialize

import (
	"gin-vue-admin/global"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

func Redis() {
	redisCfg := global.GvaConfig.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping().Result()
	if err != nil {
		global.GvaLog.Error("redis connect ping failed, err:", zap.Any("err", err))
	} else {
		global.GvaLog.Info("redis connect ping response:", zap.String("pong",pong))
		global.GvaRedis = client
	}
}
