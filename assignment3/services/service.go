package services

import (
	"assignment3/models"
	"math/rand"
	"time"
)

func GenerateRandomNumber() int {
	return rand.Intn(16)
}

func GenerateData() models.Data {
	waterQuality, windSpeed := GenerateRandomNumber(), GenerateRandomNumber()

	waterStatus := "Aman"

	if waterQuality < 5 {
		waterStatus = "Aman"
	} else if waterQuality >= 5 && waterQuality <= 8 {
		waterStatus = "Siaga"
	} else {
		waterStatus = "Bahaya"
	}

	windStatus := "Aman"
	if windSpeed < 6 {
		windStatus = "Aman"
	} else if windSpeed >= 6 && windSpeed <= 15 {
		windStatus = "Siaga"
	} else {
		windStatus = "Bahaya"
	}

	return models.Data{
		WindSpeed:    windSpeed,
		WaterQuality: waterQuality,
		WindStatus:   windStatus,
		WaterStatus:  waterStatus,
	}
}

func StartUpdatingData() {
	for {
		GenerateData()
		time.Sleep(15 * time.Second)
	}
}
