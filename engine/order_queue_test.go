package engine

import (
	"github.com/guoxiaopeng875/matching-engine/enum"
	"testing"
)

func TestOrderQueue_addOrder(t *testing.T) {
	q := new(orderQueue)
	q.init(enum.DESC)
	q.addOrder(&Order{})
}
