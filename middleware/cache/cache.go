package cache

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
