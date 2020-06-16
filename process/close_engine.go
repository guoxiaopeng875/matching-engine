package process

import (
	"github.com/guoxiaopeng875/matching-engine/engine"
	"github.com/guoxiaopeng875/matching-engine/errcode"
)

func CloseEngine(symbol string) *errcode.Errcode {
	if engine.ChanMap[symbol] == nil {
		return errcode.EngineNotFound
	}

	close(engine.ChanMap[symbol])

	return errcode.OK
}
