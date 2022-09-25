package models

import (
	"time"

	"github.com/deevarindu/hacktiv8_assignments/assignment2/httpserver/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Order struct {
	OrderId      int       `gorm:"primaryKey;autoIncrement" json:"order_id"`
	CustomerName string    `json:"customer_name"`
	OrderedAt    time.Time `json:"ordered_at"`
	Items        []Item    `gorm:"foreignkey:OrderID" json:"items"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Order{})
}

func (order *Order) CreateOrder() (*Order, error) {
	if err := db.Create(order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func (order *Order) GetOrder() (*Order, error) {
	if err := db.Where("order_id = ?", order.OrderId).First(order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func (order *Order) UpdateOrder() (*Order, error) {
	if err := db.Save(order).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func (order *Order) DeleteOrder() (*Order, error) {
	if err := db.Delete(order).Error; err != nil {
		return nil, err
	}
	return order, nil
}	

func (order *Order) GetAllOrders() (*[]Order, error) {
	var orders []Order
	if err := db.Find(&orders).Error; err != nil {
		return nil, err
	}
	return &orders, nil
}

func (order *Order) GetOrderItems() (*[]Item, error) {
	var items []Item
	if err := db.Where("order_id = ?", order.OrderId).Find(&items).Error; err != nil {
		return nil, err
	}
	return &items, nil
}

