package cache

import (
	"github.com/go-redis/redis"
	"github.com/guoxiaopeng875/matching-engine/config"
	"github.com/guoxiaopeng875/matching-engine/log"
	"go.uber.org/zap"
)

var RedisClient *redis.Client

func Init() {
	conf := config.Conf.Redis
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     conf.Addr,
		Password: conf.Password, // no password set
		DB:       conf.DB,       // use default DB
	})

	_, err := RedisClient.Ping().Result()
	if err != nil {
		panic(err)
	}
	log.Logger.Info("Connected to redis", zap.String("addr", conf.Addr))
}
