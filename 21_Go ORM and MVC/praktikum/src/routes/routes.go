package routes

import (
	"gorm-mvc-api/src/controllers/users"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	route := echo.New()

	route.GET("/users", users.GetAllUsers)
	route.POST("/users", users.CreateUser)
	route.GET("/users/:id", users.GetUser)
	route.PUT("/users/:id", users.UpdateUser)
	route.DELETE("/users/:id", users.DeleteUser)

	return route
}
