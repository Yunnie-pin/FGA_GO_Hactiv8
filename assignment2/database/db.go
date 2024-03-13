package database

import (
	"assignment2/common"
	"assignment2/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	config := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",
		common.GetDBHost(),
		common.GetDBPort(),
		common.GetDBUsername(),
		common.GetDBPassword(),
		common.GetDBName())

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database : ", err)
	}

	//untuk membuat tabel migrasi
	db.Debug().AutoMigrate(&models.Order{}, &models.Item{})

	fmt.Println("Connected to database")

}

func GetDB() *gorm.DB {
	return db
}
