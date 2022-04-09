package configs

import "github.com/caarlos0/env/v6"

type JwtSecret struct {
	SecretKey string `env:"JWT_SECRET"`
}

func GetJwtSecret() JwtSecret {
	config := JwtSecret{}
	env.Parse(&config)
	return config
}
