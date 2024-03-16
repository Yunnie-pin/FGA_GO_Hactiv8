package handlers

import (
	"assignment3/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetWindWaterData(c *gin.Context) {
	windWaterData := services.GenerateData()
	responseData := windWaterData.ToResponse()

	// Menggunakan gin.IndentedJSON untuk menghasilkan JSON yang terstruktur.
	c.IndentedJSON(http.StatusOK, gin.H{
		"data": responseData,
	})
}
