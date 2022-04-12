package utils

import (
	"middleware-api/src/configs"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtCustomClaims struct {
	jwt.StandardClaims
}

func GenerateJwt(userId string) (string, error) {
	claims := JwtCustomClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			Id:        userId,
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return rawToken.SignedString([]byte(configs.GetJwtSecret().SecretKey))
}
