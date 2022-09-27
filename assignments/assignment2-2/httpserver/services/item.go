package services

import (
	"github.com/deevarindu/hacktiv8_assignments/assignment2-2/httpserver/repositories/models"
	"github.com/jinzhu/gorm"
)

func CreateItem(db *gorm.DB, item *models.Item) error {
	err := db.Create(&item).Error
	if err != nil {
		return err
	}
	return nil
}

func GetItems(db *gorm.DB) ([]models.Item, error) {
	var items []models.Item
	err := db.Find(&items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}