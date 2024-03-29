package routes

import (
	"gorm-mvc-api/src/controllers/books"
	"gorm-mvc-api/src/controllers/users"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	route := echo.New()

	// users
	route.GET("/users", users.GetAllUsers)
	route.POST("/users", users.CreateUser)
	route.GET("/users/:id", users.GetUser)
	route.PUT("/users/:id", users.UpdateUser)
	route.DELETE("/users/:id", users.DeleteUser)

	// books
	route.GET("/books", books.GetAllBooks)
	route.POST("/books", books.CreateBook)
	route.GET("/books/:id", books.GetBook)
	route.PUT("/books/:id", books.UpdateBook)
	route.DELETE("/books/:id", books.DeleteBook)

	return route
}
