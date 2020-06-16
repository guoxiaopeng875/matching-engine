package errcode

import (
	"fmt"
	"sync/atomic"
)

var (
	_messages atomic.Value         // NOTE: stored map[string]map[int]string
	_codes    = map[int]struct{}{} // register codes.
)

// Init init ecode message map.
func Init(cm map[int]string) {
	_messages.Store(cm)
}

type Errcode struct {
	msg  string
	code code
}

func (e *Errcode) Code() int {
	return int(e.code)
}

func (e *Errcode) String() string {
	if cm, ok := _messages.Load().(map[int]string); ok {
		if msg, ok := cm[e.Code()]; ok {
			return msg
		}
	}
	return "unknown code"
}

func (e *Errcode) Error() string {
	return e.String()
}

func (e *Errcode) IsOK() bool {
	return e == OK
}

type code int

func New(code int) *Errcode {
	return &Errcode{code: add(code)}
}

func add(e int) code {
	if _, ok := _codes[e]; ok {
		panic(fmt.Sprintf("ecode: %d already exist", e))
	}
	_codes[e] = struct{}{}
	return code(e)
}
