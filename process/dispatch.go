package process

import (
	"github.com/guoxiaopeng875/matching-engine/engine"
	"github.com/guoxiaopeng875/matching-engine/enum"
	"github.com/guoxiaopeng875/matching-engine/errcode"
	"github.com/guoxiaopeng875/matching-engine/middleware/cache"
	"time"
)

func Dispatch(order engine.Order) *errcode.Errcode {
	if engine.ChanMap[order.Symbol] == nil {
		return errcode.EngineNotFound
	}
	switch order.Action {
	case enum.ActionCreate:
		if cache.OrderExist(order.Symbol, order.OrderId, order.Action.String()) {
			return errcode.OrderExist
		}
	case enum.ActionCancel:
		if !cache.OrderExist(order.Symbol, order.OrderId, enum.ActionCreate.String()) {
			return errcode.OrderNotFound
		}
	default:
		return errcode.UnknownOrderAction
	}
	order.Timestamp = time.Now().UnixNano() / 1e3
	cache.SaveOrder(order.ToMap())
	engine.ChanMap[order.Symbol] <- order
	return errcode.OK
}
