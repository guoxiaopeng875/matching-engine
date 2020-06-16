package process

import (
	"github.com/guoxiaopeng875/matching-engine/engine"
	"github.com/guoxiaopeng875/matching-engine/errcode"
	"github.com/guoxiaopeng875/matching-engine/middleware/cache"
	"github.com/shopspring/decimal"
)

func NewEngine(symbol string, price decimal.Decimal) *errcode.Errcode {
	if engine.ChanMap[symbol] != nil {
		return errcode.EngineExist
	}

	engine.ChanMap[symbol] = make(chan engine.Order, 100)
	go engine.Run(symbol, price)

	cache.SaveSymbol(symbol)
	cache.SavePrice(symbol, price)

	return errcode.OK
}
