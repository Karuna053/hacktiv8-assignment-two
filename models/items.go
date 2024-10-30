package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    string `json:"quantity"`
	OrderID     Order  `gorm:"foreignKey:Id"`
}
