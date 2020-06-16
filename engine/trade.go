package engine

import (
	"github.com/dipperin/go-ms-toolkit/json"
	"github.com/guoxiaopeng875/matching-engine/enum"
	"github.com/shopspring/decimal"
)

// Trade 成交记录
type Trade struct {
	// 挂单id(本来挂在交易委托账本里的订单)
	MakerId string `json:"maker_id"`
	// 吃单id
	TakerId string `json:"taker_id"`
	// 吃单的买卖方向
	TakerSide enum.OrderSide `json:"taker_side"`
	// 成交数量
	Amount decimal.Decimal `json:"amount"`
	// 成交价格
	Price decimal.Decimal `json:"price"`
	// 成交时间
	Timestamp int64 `json:"timestamp"`
}

func (t *Trade) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"maker_id":   t.MakerId,
		"taker_id":   t.TakerId,
		"taker_side": t.TakerSide.String(),
		"amount":     t.Amount.String(),
		"price":      t.Price.String(),
		"timestamp":  t.Timestamp,
	}
}

func (t *Trade) ToJSON() string {
	return json.StringifyJson(t)
}
