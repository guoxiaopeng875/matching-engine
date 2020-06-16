package enum

type OrderAction string

func (o OrderAction) String() string {
	return string(o)
}

const (
	// 下单操作
	ActionCreate OrderAction = "create"
	// 取消订单操作
	ActionCancel OrderAction = "cancel"
)

func (o OrderAction) IsValid() bool {
	return o == ActionCancel || o == ActionCreate
}
