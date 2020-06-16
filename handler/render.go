package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/guoxiaopeng875/matching-engine/errcode"
	"net/http"
)

type baseResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
}

func renderJSON(ctx *gin.Context, err *errcode.Errcode, data ...interface{}) {
	if err.IsOK() {
		ctx.JSON(http.StatusOK, &baseResponse{
			Success: true,
			Data:    data,
			Code:    err.Code(),
			Msg:     err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusBadRequest, &baseResponse{
		Success: false,
		Data:    data,
		Code:    err.Code(),
		Msg:     err.Error(),
	})
}
