package models

type Order struct {
	Id           int    `gorm:"primaryKey"`
	CustomerName string `json:"customer_name"`
	OrderedAt    string `json:"order_date"`
}
