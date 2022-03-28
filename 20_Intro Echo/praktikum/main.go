package main

import (
	"static-api/src/routes"
)

func main() {
	api := routes.New()
	api.Logger.Fatal(api.Start(":8000"))
}
