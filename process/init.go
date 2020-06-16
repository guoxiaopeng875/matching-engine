package process

import (
	"github.com/guoxiaopeng875/matching-engine/engine"
	"github.com/guoxiaopeng875/matching-engine/middleware/cache"
)

func Init() {
	symbols := cache.GetSymbols()
	for _, symbol := range symbols {
		price := cache.GetPrice(symbol)
		if err := NewEngine(symbol, price); !err.IsOK() {
			panic(err)
		}
		orderIds := cache.GetOrderIdsWithAction(symbol)
		for _, orderId := range orderIds {
			mapOrder := cache.GetOrder(symbol, orderId)
			order := engine.Order{}
			order.FromMap(mapOrder)
			engine.ChanMap[order.Symbol] <- order
		}
	}
}
