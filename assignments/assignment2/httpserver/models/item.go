package models

import (
	"github.com/deevarindu/hacktiv8_assignments/assignment2/httpserver/config"
)

type Item struct {
	ItemId      int    `gorm:"primaryKey;autoIncrement" json:"item_id"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderId     int    `json:"order_id"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Item{})
}

func (item *Item) CreateItem() (*Item, error) {
	if err := db.Create(item).Error; err != nil {
		return nil, err
	}
	return item, nil
}

func (item *Item) GetItem() (*Item, error) {
	if err := db.Where("item_id = ?", item.ItemId).First(item).Error; err != nil {
		return nil, err
	}
	return item, nil
}

func (item *Item) UpdateItem() (*Item, error) {
	if err := db.Save(item).Error; err != nil {
		return nil, err
	}
	return item, nil
}

func (item *Item) DeleteItem() (*Item, error) {
	if err := db.Delete(item).Error; err != nil {
		return nil, err
	}
	return item, nil
}