package models

import "time"

type Order struct {
	OrderId      int       `gorm:"primaryKey;autoIncrement" json:"order_id"`
	CustomerName string    `json:"customer_name"`
	OrderedAt    time.Time `json:"ordered_at"`
	Items        []Item    `gorm:"foreignkey:OrderID" json:"items"`
}
