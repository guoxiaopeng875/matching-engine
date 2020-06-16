package enum

type OrderSide string

func (o OrderSide) String() string {
	return string(o)
}

const (
	// 买单
	SideBuy OrderSide = "buy"
	// 卖单
	SideSell OrderSide = "sell"
)

func (o OrderSide) IsValid() bool {
	return o == SideBuy || o == SideSell
}
