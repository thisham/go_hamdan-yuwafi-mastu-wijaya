package main

import (
	"middleware-api/src/database"
	"middleware-api/src/routes"
)

func init() {
	database.Migrate()
	// database.Demigrate()
}

func main() {
	api := routes.New()
	api.Logger.Fatal(api.Start(":8000"))
}
