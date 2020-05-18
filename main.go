package main

import (
	"github.com/guoxiaopeng875/matching-engine/config"
	"github.com/guoxiaopeng875/matching-engine/engine"
	"github.com/guoxiaopeng875/matching-engine/log"
	"github.com/guoxiaopeng875/matching-engine/middleware"
	"github.com/guoxiaopeng875/matching-engine/process"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"net/http"
)

func init() {
	config.Init()
	log.Init()
	engine.Init()
	middleware.Init()
	process.Init()
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/openMatching", handler.OpenMatching)
	mux.HandleFunc("/closeMatching", handler.CloseMatching)
	mux.HandleFunc("/handleOrder", handler.HandleOrder)

	log.Logger.Info("HTTP ListenAndServe at port %s", zap.String("port", viper.GetString("server.port")))
	if err := http.ListenAndServe(viper.GetString("server.port"), mux); err != nil {
		panic(err)
	}
}
