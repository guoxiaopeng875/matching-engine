package cache

import (
	"github.com/go-redis/redis"
	"github.com/shopspring/decimal"
)

func SaveSymbol(symbol string) {
	key := "matching:symbols"
	RedisClient.SAdd(key, symbol)
}

func RemoveSymbol(symbol string) {
	key := "matching:symbols"
	RedisClient.SRem(key, symbol)
}

func GetSymbols() []string {
	key := "matching:symbols"
	return RedisClient.SMembers(key).Val()
}

func SavePrice(symbol string, price decimal.Decimal) {
	key := "matching:price:" + symbol
	RedisClient.Set(key, price.String(), 0)
}

func GetPrice(symbol string) decimal.Decimal {
	key := "matching:price:" + symbol
	priceStr := RedisClient.Get(key).Val()
	result, err := decimal.NewFromString(priceStr)
	if err != nil {
		result = decimal.Zero
	}
	return result
}

func RemovePrice(symbol string) {
	key := "matching:price:" + symbol
	RedisClient.Del(key)
}

func SaveOrder(order map[string]interface{}) {
	symbol := order["symbol"].(string)
	orderId := order["orderId"].(string)
	timestamp := order["timestamp"].(int64)
	action := order["action"].(string)

	key := "matching:order:" + symbol + ":" + orderId + ":" + action
	RedisClient.HMSet(key, order)

	key = "matching:orderids:" + symbol
	z := redis.Z{
		Score:  float64(timestamp),
		Member: orderId + ":" + action,
	}
	RedisClient.ZAdd(key, z)
}

func GetOrder(symbol, orderIdWithAction string) map[string]interface{} {
	key := "matching:order:" + symbol + ":" + orderIdWithAction
	fields := []string{
		"action",
		"symbol",
		"orderId",
		"side",
		"type",
		"amount",
		"price",
		"timestamp",
	}
	vals := RedisClient.HMGet(key, fields...).Val()
	if len(vals) == 0 {
		return nil
	}
	order := make(map[string]interface{}, len(vals))
	for i, val := range vals {
		field := fields[i]
		order[field] = val
	}
	return order
}

func RemoveOrder(order map[string]interface{}) {
	var (
		symbol  = order["symbol"].(string)
		orderId = order["orderId"].(string)
		action  = order["action"].(string)
	)

	removeOrder(symbol, orderId+action)
}

func removeOrder(symbol, orderIdWithAction string) {
	key := "matching:order:" + symbol + ":" + orderIdWithAction
	RedisClient.Del(key)
}

// orderId:action
func GetOrderIdsWithAction(symbol string) []string {
	key := "matching:orderids:" + symbol
	return RedisClient.ZRange(key, 0, -1).Val()
}

func RemoveOrderIdsWithAction(symbol string) {
	key := "matching:orderids:" + symbol
	RedisClient.Del(key)
}

// OrderExist 判断订单是否存在
func OrderExist(symbol, orderId, action string) bool {
	key := "matching:order:" + symbol + ":" + orderId + ":" + action
	return RedisClient.Exists(key).Val() > 0
}

// Clear 清除缓存
func Clear(symbol string) {
	RemovePrice(symbol)
	RemoveSymbol(symbol)
	orderIds := GetOrderIdsWithAction(symbol)
	for _, orderId := range orderIds {
		removeOrder(symbol, orderId)
	}
	RemoveOrderIdsWithAction(symbol)
}
