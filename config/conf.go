package config

import (
	"fmt"
	"github.com/dipperin/go-ms-toolkit/json"
	"github.com/guoxiaopeng875/matching-engine/log"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
	"path"
)

type Config struct {
	Redis    *Redis         `json:"redis"`
	ErrCodes map[int]string `json:"err_codes"`
}

type Redis struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

var (
	Conf     *Config
	TestPath = path.Join(os.Getenv("GOPATH"), "src/matching-engine/config")
)

func Init(path string) {
	log.Logger.Info("配置文件路径", zap.String("path", path))

	viper.SetConfigName("conf") // name of config file (without extension)
	viper.SetConfigType("json") // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(path)   // path to look for the config file in
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	Conf = new(Config)
	var (
		redis    Redis
		errCodes = make(map[int]string)
	)
	if err := json.ParseJson(json.StringifyJson(viper.Get("redis")), &redis); err != nil {
		panic(err)
	}
	if err := json.ParseJson(json.StringifyJson(viper.Get("err_codes")), &errCodes); err != nil {
		panic(err)
	}
	Conf.Redis = &redis
	Conf.ErrCodes = errCodes
}
