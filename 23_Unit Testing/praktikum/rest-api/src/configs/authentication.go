package configs

type JwtSecret struct {
	SecretKey string
}

func GetJwtSecret() JwtSecret {
	config := JwtSecret{
		SecretKey: GetServerConfig().JwtSecretKey,
	}
	return config
}
