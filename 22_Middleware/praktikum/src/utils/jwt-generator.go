package utils

import (
	"middleware-api/src/configs"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	UserId = "adb182ce-dumm-dumm-dumm-dummycreated"
)

type JwtCustomClaims struct {
	jwt.StandardClaims
}

func GenerateJwt() (string, error) {
	claims := JwtCustomClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			Id:        UserId,
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return rawToken.SignedString([]byte(configs.GetJwtSecret().SecretKey))
}
