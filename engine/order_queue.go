package engine

import (
	"container/list"
	"github.com/guoxiaopeng875/matching-engine/enum"
)

// orderQueue 订单队列
type orderQueue struct {
	// 价格排序的方向
	sortBy enum.SortDirection
	// 保存整个二维链表的所有订单，第一维以价格排序，第二维以时间排序
	parentList *list.List
	// Key 为价格、Value 为第二维订单链表的键值对
	elementMap map[string]*list.Element
}

func (q *orderQueue) init(sortBy enum.SortDirection) {
	q.sortBy = sortBy
	q.parentList = list.New()
	q.elementMap = make(map[string]*list.Element)
}

// addOrder 添加订单
func (q *orderQueue) addOrder(order *Order) {
	var (
		price = order.Price.String()
	)

	// 父链表长度为0
	if q.parentList.Len() == 0 {
		childList := list.New()
		elem := childList.PushFront(order)
		q.parentList.PushFrontList(childList)
		q.elementMap[price] = elem
		return
	}
	elem, empty := q.elementMap[price]
	// 判断子链表是否为空
	if !empty {
		// 将订单插入到子链表后面
		q.parentList.InsertAfter(order, elem)
		return
	}
	childList := list.New()
	childElem := childList.PushBack(order)
	// 父链表第一个元素
	p := q.parentList.Front()
	oldOrder := p.Value.(*Order)
	for {
		if q.canPush(oldOrder, order) {
			q.parentList.PushFrontList(childList)
			q.elementMap[price] = childElem
			return
		}
		if p == nil {
			q.parentList.PushBackList(childList)
			q.elementMap[price] = childElem
			return
		}
		p = p.Next()
	}
}

func (q *orderQueue) canPush(order, newOrder *Order) bool {
	return q.isDESCAndGreaterThan(order, newOrder) || q.isASCAndLessThan(order, newOrder)
}
func (q *orderQueue) isASCAndLessThan(order, newOrder *Order) bool {
	// 新订单价格小于订单价格
	return !q.isDESC() && newOrder.Price.LessThan(order.Price)
}

func (q *orderQueue) isDESCAndGreaterThan(order, newOrder *Order) bool {
	// 新订单价格大于订单价格
	return q.isDESC() && newOrder.Price.GreaterThan(order.Price)
}

func (q *orderQueue) isDESC() bool {
	return q.sortBy == enum.DESC
}

// getHeadOrder 读取头部订单
func (q *orderQueue) getHeadOrder() *Order {
	return q.parentList.Front().Value.(*Order)
}

// popHeadOrder 读取并删除头部订单
func (q *orderQueue) popHeadOrder() *Order {
	headOrder := q.getHeadOrder()
	elem := q.getOrderElement(headOrder)
	q.parentList.Remove(elem)
	return headOrder
}

func (q *orderQueue) getOrderElement(order *Order) *list.Element {
	return q.elementMap[order.Price.String()]
}

// removeOrder 移除订单
func (q *orderQueue) removeOrder(order *Order) {
	elem := q.getOrderElement(order)
	for {
		if elem == nil {
			return
		}
		if elem.Value.(*Order).Timestamp == order.Timestamp {
			q.parentList.Remove(elem)
			return
		}
		elem = elem.Next()
	}
}

// getDepthPrice 读取深度价格
func (q *orderQueue) getDepthPrice(depth int) (string, int) {
	if q.parentList.Len() == 0 {
		return "", 0
	}
	p := q.parentList.Front()
	i := 1
	for ; i < depth; i++ {
		t := p.Next()
		if t == nil {
			break
		}
		p = t
	}
	o := p.Value.(*list.List).Front().Value.(*Order)
	return o.Price.String(), i
}
