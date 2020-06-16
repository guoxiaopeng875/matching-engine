package engine

import (
	"github.com/guoxiaopeng875/matching-engine/enum"
	"github.com/guoxiaopeng875/matching-engine/log"
	"github.com/guoxiaopeng875/matching-engine/middleware/cache"
	"github.com/guoxiaopeng875/matching-engine/middleware/mq"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"time"
)

func Run(symbol string, price decimal.Decimal) {
	lastTradePrice := price
	book := &orderBook{}
	book.init()
	log.Logger.Info("engine is running", zap.String("symbol", symbol))
	for {
		order, ok := <-ChanMap[symbol]
		if !ok {
			log.Logger.Info("engine is closed", zap.String("symbol", symbol))
			delete(ChanMap, symbol)
			cache.Clear(symbol)
			return
		}
		log.Logger.Info("engine %s receive an order: %s", zap.String("symbol", symbol), zap.String("order", order.ToJSON()))
		switch order.Action {
		case enum.ActionCreate:
			dealCreate(&order, book, &lastTradePrice)
		case enum.ActionCancel:
			dealCancel(&order, book)
		}
	}
}

func dealCreate(order *Order, book *orderBook, lastTradePrice *decimal.Decimal) {
	switch order.Type {
	case enum.TypeLimit:
		dealLimit(order, book, lastTradePrice)
		//case enum.TypeLimitIoc:
		//	dealLimitIoc(order, book, lastTradePrice)
		//case enum.TypeMarket:
		//	dealMarket(order, book, lastTradePrice)
		//case enum.TypeMarketTop5:
		//	dealMarketTop5(order, book, lastTradePrice)
		//case enum.TypeMarketTop10:
		//	dealMarketTop10(order, book, lastTradePrice)
		//case enum.TypeMarketOpponent:
		//	dealMarketOpponent(order, book, lastTradePrice)
	}
}

func dealLimit(order *Order, book *orderBook, lastTradePrice *decimal.Decimal) {
	switch order.Side {
	case enum.SideBuy:
		dealBuyLimit(order, book, lastTradePrice)
	case enum.SideSell:
		dealSellLimit(order, book, lastTradePrice)
	}
}

func dealSellLimit(order *Order, book *orderBook, lastTradePrice *decimal.Decimal) {
LOOP:
	headOrder := book.getHeadBuyOrder()
	if headOrder == nil || order.Price.GreaterThan(headOrder.Price) {
		book.addBuyOrder(order)
		log.Logger.Info("a sell order has added to the orderbook", zap.String("symbol", order.Symbol), zap.String("order", order.ToJSON()))
		return
	}
	matchTrade(headOrder, order, book, lastTradePrice)
	if order.Amount.IsPositive() {
		goto LOOP
	}
}

func dealBuyLimit(order *Order, book *orderBook, lastTradePrice *decimal.Decimal) {
LOOP:
	headOrder := book.getHeadSellOrder()
	if headOrder == nil || order.Price.LessThan(headOrder.Price) {
		book.addBuyOrder(order)
		log.Logger.Info("a buy order has added to the orderbook", zap.String("symbol", order.Symbol), zap.String("order", order.ToJSON()))
		return
	}
	matchTrade(headOrder, order, book, lastTradePrice)
	if order.Amount.IsPositive() {
		goto LOOP
	}
}

func matchTrade(headOrder, order *Order, book *orderBook, lastTradePrice *decimal.Decimal) {
	trade := &Trade{
		MakerId:   headOrder.OrderId,
		TakerId:   order.OrderId,
		TakerSide: order.Side,
		Price:     *lastTradePrice,
		Timestamp: time.Now().UnixNano(),
	}
	// 挂单数量小于吃单数量
	if headOrder.Amount.LessThanOrEqual(order.Amount) {
		trade.Amount = headOrder.Amount
		order.Amount.Sub(headOrder.Amount)
		removeOrder(headOrder, book)
	} else {
		trade.Amount = order.Amount
		headOrder.Amount.Sub(order.Amount)
		order.Amount = decimal.Zero
	}
	mq.SendTrade(order.Symbol, trade.ToMap())
}

func removeOrder(order *Order, book *orderBook) bool {
	var ok bool
	switch order.Side {
	case enum.SideBuy:
		ok = book.removeBuyOrder(order)
	case enum.SideSell:
		ok = book.removeSellOrder(order)
	}
	cache.RemoveOrder(order.ToMap())
	return ok
}

func dealCancel(order *Order, book *orderBook) {
	ok := removeOrder(order, book)
	mq.SendCancelResult(order.Symbol, order.OrderId, ok)
	log.Logger.Info("order cancel result", zap.String("symbol", order.Symbol), zap.String("orderId", order.OrderId), zap.Bool("ok", ok))
}
