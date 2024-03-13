package routers

import (
	"assignment2/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(controller *controllers.OrderController) *gin.Engine {
	router := gin.Default()

	router.POST("/orders", controller.CreateNewOrder)
	router.GET("/orders", controller.GetAllOrdersWithItems)
	router.DELETE("/orders/:id", controller.DeleteOrder)
	router.PUT("/orders/:id", controller.UpdateOrder)

	return router
}
