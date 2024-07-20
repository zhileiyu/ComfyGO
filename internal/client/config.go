package client

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	EndPoint
}

type EndPoint struct {
	BaseUrl string
	Port    int
}

func loadConfig(fileName string) (Config, error) {
	var c Config
	_, err := toml.DecodeFile(fileName, &c)
	if err != nil {
		return c, err
	}
	return c, nil
}
