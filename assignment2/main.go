package main

import (
	"assignment2/database"
	"assignment2/routers"
	"fmt"
)

func main() {
	database.StartDB()
	fmt.Println("Connected to database")

	var PORT = "8080"

	routers.StartServer().Run(":" + PORT)
}
