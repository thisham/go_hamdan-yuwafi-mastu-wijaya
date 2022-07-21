package routes

import (
	"middleware-api/src/controllers/authentications"
	"middleware-api/src/controllers/books"
	"middleware-api/src/controllers/users"
	"middleware-api/src/middlewares"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	route := echo.New()
	middlewares.Logger(route)

	// implement auth middleware
	authRoutes := route.Group("")
	authRoutes.Use(middlewares.VerifyAuthentication())

	// token-generate
	route.POST("/login", authentications.LoginHandler)

	// users
	route.GET("/users", users.GetAllUsers)
	authRoutes.POST("/users", users.CreateUser)
	authRoutes.GET("/users/:id", users.GetUser)
	authRoutes.PUT("/users/:id", users.UpdateUser)
	authRoutes.DELETE("/users/:id", users.DeleteUser)

	// books
	authRoutes.GET("/books", books.GetAllBooks)
	authRoutes.POST("/books", books.CreateBook)
	authRoutes.GET("/books/:id", books.GetBook)
	authRoutes.PUT("/books/:id", books.UpdateBook)
	authRoutes.DELETE("/books/:id", books.DeleteBook)

	return route
}
