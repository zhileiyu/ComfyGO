package client

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Endpoint `toml:"endpoint"`
}

type Endpoint struct {
	BaseUrl string `toml:"base_url"`
	Port    int    `toml:"port"`
}

func loadConfig(fileName string) (Config, error) {
	var c Config
	_, err := toml.DecodeFile(fileName, &c)
	if err != nil {
		return c, err
	}
	return c, nil
}
