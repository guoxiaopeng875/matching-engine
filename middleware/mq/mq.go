package mq

import (
	"github.com/go-redis/redis"
	"github.com/guoxiaopeng875/matching-engine/middleware/cache"
)

func SendCancelResult(symbol, orderId string, ok bool) {
	values := map[string]interface{}{"orderId": orderId, "ok": ok}
	a := &redis.XAddArgs{
		Stream:       "matching:cancelresults:" + symbol,
		MaxLenApprox: 1000,
		Values:       values,
	}
	cache.RedisClient.XAdd(a)
}

func SendTrade(symbol string, trade map[string]interface{}) {
	a := &redis.XAddArgs{
		Stream:       "matching:trades:" + symbol,
		MaxLenApprox: 1000,
		Values:       trade,
	}
	cache.RedisClient.XAdd(a)
}
