package enum

type OrderType string

func (o OrderType) String() string {
	return string(o)
}

const (
	// 限价交易
	TypeLimit OrderType = "limit"
	// 市价交易
	TypeMarket OrderType = "market"
)

func (o OrderType) IsValid() bool {
	return o == TypeLimit || o == TypeMarket
}
