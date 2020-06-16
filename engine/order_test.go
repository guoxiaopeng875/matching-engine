package engine

import (
	"github.com/guoxiaopeng875/matching-engine/enum"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func TestOrder_FromToMap(t *testing.T) {
	o := &Order{
		Action:    enum.ActionCreate,
		Symbol:    "abc",
		OrderId:   "123",
		Side:      enum.SideBuy,
		Type:      enum.TypeLimit,
		Amount:    decimal.NewFromInt(11),
		Price:     decimal.NewFromFloat(33.3),
		Timestamp: time.Now().UnixNano(),
	}
	//fmt.Println(o.ToJSON())
	var o2 Order
	m := o.ToMap()
	m["timestamp"] = strconv.FormatInt(o.Timestamp, 10)
	o2.FromMap(m)
	assert.True(t, reflect.DeepEqual(o, &o2))
}
