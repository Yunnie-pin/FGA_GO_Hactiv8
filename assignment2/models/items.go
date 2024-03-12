package models

// import (
// 	"errors"
// 	// "fmt"
// 	// "time"
// 	"gorm.io/gorm"
// )

type Item struct {
	ItemID      int    `gorm:"primaryKey" json:"item_id"`
	ItemCode    string `gorm:"unique type:varchar(100)" json:"itemCode"`
	Description string `gorm:"type:varchar(100)" json:"description"`
	Quantity    int    `gorm:"type:integer" json:"quantity"`
	OrderID     int    `json:"order_id"`
}

// func (p *Item) BeforeCreate(tx *gorm.DB) (err error) {
// 	if p.Quantity < 0 {
// 		return errors.New("KUANTITAS TIDAK BOLEH KURANG DARI 0")
// 	}
// 	return
// }
