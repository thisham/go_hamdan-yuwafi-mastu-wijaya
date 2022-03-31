package routes

import (
	"gorm-mvc-api/src/controllers/users"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	route := echo.New()

	route.GET("/users", users.GetAllUsers)
	route.POST("/users", users.CreateUser)

	return route
}
