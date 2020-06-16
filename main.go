package main

import (
	"github.com/gin-gonic/gin"
	"github.com/guoxiaopeng875/matching-engine/config"
	"github.com/guoxiaopeng875/matching-engine/engine"
	"github.com/guoxiaopeng875/matching-engine/errcode"
	"github.com/guoxiaopeng875/matching-engine/handler"
	"github.com/guoxiaopeng875/matching-engine/log"
	"github.com/guoxiaopeng875/matching-engine/middleware"
	"github.com/guoxiaopeng875/matching-engine/process"
)

func init() {
	config.Init(config.TestPath)
	log.Init()
	errcode.Init(config.Conf.ErrCodes)
	engine.Init()
	middleware.Init()
	process.Init()
}

func main() {
	e := gin.Default()
	e.POST("/open_matching", handler.OpenMatching())
	e.POST("/close_matching", handler.CloseMatching())
	e.POST("/handle_order", handler.HandleOrder())
	if err := e.Run(":8989"); err != nil {
		panic(err)
	}
}
