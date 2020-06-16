package cache

import (
	"github.com/guoxiaopeng875/matching-engine/config"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func init() {
	config.Init(config.TestPath)
	Init()
}

func TestSymbol(t *testing.T) {
	SaveSymbol("a_b")
	SaveSymbol("a_c")
	symbols := GetSymbols()
	assert.Equal(t, 2, len(symbols))
	assert.Equal(t, "a_c", symbols[0])
	assert.Equal(t, "a_b", symbols[1])
	RemoveSymbol("a_b")
	symbols = GetSymbols()
	assert.Equal(t, 1, len(symbols))
}

func TestPrice(t *testing.T) {
	var (
		p1 = decimal.NewFromInt(1)
		p2 = decimal.NewFromInt(2)
	)
	SavePrice("aa", p1)
	SavePrice("bb", p2)
	price := GetPrice("aa")
	assert.True(t, price.Equal(p1))
	RemovePrice("aa")
	price = GetPrice("aa")
	assert.Equal(t, decimal.Zero, price)
}

func TestOrder(t *testing.T) {
	SaveOrder(map[string]interface{}{
		"action":    "create",
		"symbol":    "aa",
		"orderId":   "123",
		"side":      "sell",
		"type":      "aa",
		"amount":    "10",
		"price":     "33.33",
		"timestamp": time.Now().UnixNano(),
	})
	o2 := map[string]interface{}{
		"action":    "create",
		"symbol":    "aa",
		"orderId":   "1234",
		"side":      "buy",
		"type":      "aa",
		"amount":    "10",
		"price":     "33.33",
		"timestamp": time.Now().UnixNano(),
	}
	SaveOrder(o2)
	assert.False(t, OrderExist("aa", "000", "create"))
	assert.True(t, OrderExist("aa", "123", "create"))
	orderIds := GetOrderIdsWithAction("aa")
	assert.Equal(t, 2, len(orderIds))
	orderId := orderIds[1]
	order := GetOrder("aa", orderId)
	for k, v := range order {
		if k == "timestamp" {
			continue
		}
		assert.Equal(t, o2[k], v)
	}
	Clear("aa")
	orderIds = GetOrderIdsWithAction("aa")
	assert.Equal(t, 0, len(orderIds))
}
