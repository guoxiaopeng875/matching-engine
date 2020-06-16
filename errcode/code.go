package errcode

var (
	OK                 = New(0)
	EngineNotFound     = New(1)
	OrderExist         = New(2)
	OrderNotFound      = New(3)
	EngineExist        = New(4)
	UnknownOrderAction = New(5)
	InvalidParams      = New(6)
)
