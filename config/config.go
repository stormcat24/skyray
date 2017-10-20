package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Port string `envconfig:"port" default:"5514"`
}

func Parse() *Config {
	appConf := &Config{}
	err := envconfig.Process("skyray", appConf)
	if err != nil {
		panic(err)
	}
	return appConf
}
