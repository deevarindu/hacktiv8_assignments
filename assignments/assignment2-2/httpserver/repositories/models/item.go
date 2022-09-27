package models

type Item struct {
	ItemId      int    `gorm:"primaryKey;autoIncrement" json:"item_id"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderId     int    `json:"order_id"`
}