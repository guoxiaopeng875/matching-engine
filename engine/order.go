package engine

import (
	"github.com/dipperin/go-ms-toolkit/json"
	"github.com/guoxiaopeng875/matching-engine/enum"
	"github.com/guoxiaopeng875/matching-engine/errcode"
	"github.com/shopspring/decimal"
	"strconv"
)

// Order 委托单
type Order struct {
	// 声明对委托单要进行哪种操作, 下单(create)和撤单(cancel)
	Action enum.OrderAction `json:"action"`
	// symbol 指定该委托单所属的交易对
	Symbol string `json:"symbol"`
	// 该委托单的唯一标识
	OrderId string `json:"orderId"`
	// 买入(buy)/卖出(sell)
	Side enum.OrderSide `json:"side"`
	// 交易类型
	// 限价交易(limit)或市价交易(market)等
	Type enum.OrderType `json:"type"`
	// 购买数量
	Amount decimal.Decimal `json:"amount"`
	// 购买金额
	Price decimal.Decimal `json:"price"`
	// 订单时间
	Timestamp int64 `json:"timestamp"`
}

func (o Order) IsValid() *errcode.Errcode {
	if !o.Action.IsValid() || o.Symbol == "" || o.OrderId == "" || !o.Side.IsValid() || !o.Type.IsValid() || o.Amount.IsNegative() || o.Price.IsNegative() {
		return errcode.InvalidParams
	}
	return errcode.OK
}

func (o *Order) FromMap(oMap map[string]interface{}) {
	o.Action = enum.OrderAction(oMap["action"].(string))
	o.Symbol = oMap["symbol"].(string)
	o.OrderId = oMap["orderId"].(string)
	o.Side = enum.OrderSide(oMap["side"].(string))
	o.Type = enum.OrderType(oMap["type"].(string))
	o.Amount, _ = decimal.NewFromString(oMap["amount"].(string))
	o.Price, _ = decimal.NewFromString(oMap["price"].(string))
	ts := oMap["timestamp"]
	switch ts.(type) {
	case string:
		o.Timestamp, _ = strconv.ParseInt(ts.(string), 10, 64)
	case int64:
		o.Timestamp = ts.(int64)
	}
}

func (o *Order) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"action":    o.Action.String(),
		"symbol":    o.Symbol,
		"orderId":   o.OrderId,
		"side":      o.Side.String(),
		"type":      o.Type.String(),
		"amount":    o.Amount.String(),
		"price":     o.Price.String(),
		"timestamp": o.Timestamp,
	}
}

func (o *Order) ToJSON() string {
	return json.StringifyJson(o)
}
