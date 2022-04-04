package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	UserId    = "adb182ce-dumm-dumm-dumm-dummycreated"
	UserName  = "Daniel"
	SecretKey = "secret"
)

type JwtCustomClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

func GenerateJwt() (string, error) {
	claims := JwtCustomClaims{
		UserName,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			Id:        UserId,
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return rawToken.SignedString([]byte(SecretKey))
}
