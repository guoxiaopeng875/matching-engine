package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/guoxiaopeng875/matching-engine/errcode"
	"github.com/guoxiaopeng875/matching-engine/process"
)

type closeMatchingParams struct {
	Symbol string `json:"symbol"`
}

func (p closeMatchingParams) isValid() *errcode.Errcode {
	if p.Symbol == "" {
		return errcode.InvalidParams
	}
	return errcode.OK
}

func CloseMatching() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var params closeMatchingParams
		if ctx.BindJSON(&params) != nil {
			renderJSON(ctx, errcode.InvalidParams)
			return
		}
		if err := params.isValid(); !err.IsOK() {
			renderJSON(ctx, err)
			return
		}
		renderJSON(ctx, process.CloseEngine(params.Symbol))
	}
}
