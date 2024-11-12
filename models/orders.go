package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CustomerName string    `json:"customer_name"`
	OrderedAt    time.Time `json:"order_date"`
	Items        []Item    `gorm:"foreignKey:OrderID"` // One-to-many relationship with Items
}
