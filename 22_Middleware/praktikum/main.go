package main

import (
	"fmt"
	"middleware-api/src/configs"
	"middleware-api/src/database"
	"middleware-api/src/routes"
)

func init() {
	database.Migrate()
	// database.Demigrate()
}

func main() {
	api := routes.New()
	serverPort := configs.GetServerConfig().Port
	api.Logger.Fatal(api.Start(fmt.Sprintf(":%s", serverPort)))
}
