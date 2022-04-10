package configs

import "github.com/caarlos0/env/v6"

type ServerConfig struct {
	Port int `env:"PORT" envDefault:"8000"`
}

func GetServerConfig() ServerConfig {
	config := ServerConfig{}
	env.Parse(&config)
	return config
}
