package routes

import (
	"assignment3/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() {
	router := gin.Default()

	router.GET("/", handlers.GetWindWaterData)

	router.Run(":5555")
}
