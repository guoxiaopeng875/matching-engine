package config

import (
	"fmt"
	"github.com/dipperin/go-ms-toolkit/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInit(t *testing.T) {
	Init(TestPath)
	assert.NotNil(t, Conf)
	assert.NotNil(t, Conf.Redis)
	assert.Equal(t, "localhost:6379", Conf.Redis.Addr)
	Conf.ErrCodes = map[int]string{
		1: "成功",
	}
	fmt.Println(json.StringifyJson(Conf))
}
