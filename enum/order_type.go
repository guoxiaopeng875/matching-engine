package enum

type OrderType string

const (
	// 限价交易
	TypeLimit OrderType = "limit"
	// 市价交易
	TypeMarket OrderType = "market"
)
