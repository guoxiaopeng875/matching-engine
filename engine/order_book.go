package engine

import "github.com/guoxiaopeng875/matching-engine/enum"

type orderBook struct {
	// 买单队列
	buyOrderQueue *orderQueue
	// 卖单队列
	sellOrderQueue *orderQueue
}

func (ob *orderBook) init() {
	ob.buyOrderQueue = new(orderQueue)
	ob.sellOrderQueue = new(orderQueue)
	// 买单队列是按照价格降序的
	ob.buyOrderQueue.init(enum.DESC)
	// 卖单队列则是按照价格升序的
	ob.sellOrderQueue.init(enum.ASC)
}

// addBuyOrder 增加买单委托单
func (ob *orderBook) addBuyOrder(order *Order) {
	ob.buyOrderQueue.addOrder(order)
}

// addSellOrder 增加卖单委托单
func (ob *orderBook) addSellOrder(order *Order) {
	ob.sellOrderQueue.addOrder(order)
}

// getHeadBuyOrder 获取头部买单委托单
func (ob *orderBook) getHeadBuyOrder() *Order {
	return ob.buyOrderQueue.getHeadOrder()
}

// getHeadSellOrder 获取头部卖单委托单
func (ob *orderBook) getHeadSellOrder() *Order {
	return ob.sellOrderQueue.getHeadOrder()
}

// popHeadBuyOrder 获取并移除头部买单委托单
func (ob *orderBook) popHeadBuyOrder() *Order {
	return ob.buyOrderQueue.popHeadOrder()
}

// popHeadSellOrder 获取并移除头部卖单委托单
func (ob *orderBook) popHeadSellOrder() *Order {
	return ob.sellOrderQueue.popHeadOrder()
}

// removeBuyOrder 移除买单委托单
func (ob *orderBook) removeBuyOrder(order *Order) {
	ob.buyOrderQueue.removeOrder(order)
}

// removeSellOrder 移除卖单委托单
func (ob *orderBook) removeSellOrder(order *Order) {
	ob.sellOrderQueue.removeOrder(order)
}
