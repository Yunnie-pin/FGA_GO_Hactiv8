package main

import (
	// "assignment2/common"
	"assignment2/controllers"
	"assignment2/database"
	"assignment2/repositories"
	"assignment2/routers"
	"assignment2/services"
	"fmt"
)

func main() {
	database.StartDB()
	port := ":8080"

	orderRepo := repositories.NewOrderRepo(database.GetDB())
	orderService := services.NewOrderService(orderRepo)

	itemRepo := repositories.NewItemRepo(database.GetDB())
	itemService := services.NewItemService(itemRepo)

	orderController := controllers.NewOrderController(orderService, itemService)

	routers.SetupRouter(orderController).Run(port)

	fmt.Printf("Server is running")

}
