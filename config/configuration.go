package config

import (
	"fmt"
	"github.com/caarlos0/env"
)

type Configuration struct {
	Port  int  `env:"PORT" envDefault:"3000"`
}

func GetConfiguration() Configuration {
	configuration := Configuration{}
	err := env.Parse(&configuration)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	return configuration
}