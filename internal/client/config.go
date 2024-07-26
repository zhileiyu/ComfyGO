package client

import (
	"github.com/BurntSushi/toml"
	"github.com/zhileiyu/comfyGO/internal/logger"
	"os"
)

type Config struct {
	Endpoint `toml:"endpoint"`
}

type Endpoint struct {
	BaseUrl string `toml:"base_url"`
	Port    int    `toml:"port"`
}

func loadConfig(fileName string) *Config {
	c := &Config{}
	_, err := toml.DecodeFile(fileName, c)
	if os.IsNotExist(err) {
		logger.Error("config file not exist")
		return nil
	} else if err != nil {
		logger.Error("decode config file fail " + err.Error())
		return nil
	}
	return c
}
