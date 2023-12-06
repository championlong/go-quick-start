package initialize

import (
	"context"

	"github.com/championlong/go-quick-start/internal/app/global"
	"github.com/championlong/go-quick-start/pkg/log"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

func Redis() {
	redisCfg := global.GVA_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Error("redis connect ping failed, err:", zap.Error(err))
	} else {
		log.Info("redis connect ping response:", zap.String("pong", pong))
		global.GVA_REDIS = client
	}
}
