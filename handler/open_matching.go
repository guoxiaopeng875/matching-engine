package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/guoxiaopeng875/matching-engine/errcode"
	"github.com/guoxiaopeng875/matching-engine/process"
	"github.com/shopspring/decimal"
)

type openMatchingParams struct {
	Symbol string          `json:"symbol"`
	Price  decimal.Decimal `json:"price"`
}

func (p openMatchingParams) isValid() *errcode.Errcode {
	if p.Symbol == "" || p.Price.IsNegative() {
		return errcode.InvalidParams
	}
	return errcode.OK
}

func OpenMatching() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var params openMatchingParams
		if ctx.BindJSON(&params) != nil {
			renderJSON(ctx, errcode.InvalidParams)
			return
		}
		if err := params.isValid(); !err.IsOK() {
			renderJSON(ctx, err)
			return
		}
		renderJSON(ctx, process.NewEngine(params.Symbol, params.Price))
	}
}
