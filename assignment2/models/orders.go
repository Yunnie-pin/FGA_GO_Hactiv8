package models

import (
	// "fmt"
	"time"
)

type Order struct {
	OrderID      int    `gorm:"primaryKey" json:"order_id"`
	CustomerName string `gorm:"type:varchar(100)"`
	Items        []Item
	OrderedAt    time.Time
}
