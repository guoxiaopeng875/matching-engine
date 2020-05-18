package engine

import (
	"github.com/guoxiaopeng875/matching-engine/enum"
	"github.com/shopspring/decimal"
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
