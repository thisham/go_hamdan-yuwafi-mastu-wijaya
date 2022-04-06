package middlewares

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func VerifyAuthentication() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		TokenLookup: "cookie:token",
		SigningKey:  []byte("secret"),
		Claims:      jwt.MapClaims{},
		ContextKey:  "token",
	})
}
