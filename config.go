package awesomeProject

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Token string `yaml:"token"`
}

func ParseConfig(pathConfig string) (config *Config, err error) {
	var rawConfig []byte
	if rawConfig, err = os.ReadFile(pathConfig); err != nil {
		return
	}
	if err = yaml.Unmarshal(rawConfig, &config); err != nil {
		return
	}
	return config, err
}
