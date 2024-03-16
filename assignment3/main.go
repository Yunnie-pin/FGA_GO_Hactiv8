package main

import (
	"assignment3/routes"
	"assignment3/services"
)

func main() {
	go services.StartUpdatingData()

	routes.SetupRoutes()
}
