package engine

var ChanMap map[string]chan Order

// Init 初始化撮合引擎
func Init() {
	ChanMap = make(map[string]chan Order)
}
