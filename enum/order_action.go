package enum

type OrderAction string

const (
	// 下单操作
	ActionCreate OrderAction = "create"
	// 取消订单操作
	ActionCancel OrderAction = "cancel"
)
