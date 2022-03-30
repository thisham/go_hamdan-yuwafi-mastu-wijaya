package main

import (
	"gorm-mvc-api/src/database"
	"gorm-mvc-api/src/routes"
)

func init() {
	database.Migrate()
}

func main() {
	api := routes.New()
	api.Logger.Fatal(api.Start("0.0.0.0:8000"))
}
