package config

type Config struct {
	Redis *Redis
}

type Redis struct {
	Addr     string
	Password string
	DB       int
}

var Conf *Config

func Init() {
	Conf = new(Config)
}
