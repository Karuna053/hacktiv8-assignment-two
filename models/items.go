package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderID     uint   `json:"orderID"`                     // Foreign key to reference Order
	Order       Order  `gorm:"foreignKey:OrderID" json:"-"` // Prevent recursive JSON nesting
}
