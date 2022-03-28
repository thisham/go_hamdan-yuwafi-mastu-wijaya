package routes

import (
	"static-api/src/controllers/users"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	route := echo.New()

	route.GET("/users", users.GetAllUserController)
	route.GET("/users/:id", users.GetUserController)
	route.POST("/users", users.CreateUserController)
	route.PUT("/users/:id", users.UpdateUserController)
	route.DELETE("/users/:id", users.DeleteUserController)
	return route
}
