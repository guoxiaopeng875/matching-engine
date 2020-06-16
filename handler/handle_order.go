package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/guoxiaopeng875/matching-engine/engine"
	"github.com/guoxiaopeng875/matching-engine/errcode"
	"github.com/guoxiaopeng875/matching-engine/process"
)

type handleOrderParams struct {
	Order engine.Order `json:"order"`
}

func (p handleOrderParams) isValid() *errcode.Errcode {
	return p.Order.IsValid()
}

func HandleOrder() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var params handleOrderParams
		if ctx.BindJSON(&params) != nil {
			renderJSON(ctx, errcode.InvalidParams)
			return
		}
		if err := params.isValid(); !err.IsOK() {
			renderJSON(ctx, err)
			return
		}
		renderJSON(ctx, process.Dispatch(params.Order))
	}
}
