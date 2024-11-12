package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    string `json:"quantity"`
	OrderID     uint   `json:"order_id"`           // Foreign key to reference Order
	Order       Order  `gorm:"foreignKey:OrderID"` // GORM relation to Order
}
