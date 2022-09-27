package services

import (
	"github.com/deevarindu/hacktiv8_assignments/assignment2-2/httpserver/repositories/models"
	"github.com/jinzhu/gorm"
)

func CreateOrder(db *gorm.DB, order *models.Order) error {
	err := db.Create(&order).Error
	if err != nil {
		return err
	}
	return nil
}

func GetOrders(db *gorm.DB, orders *[]models.Order) error {
	err := db.Find(&orders).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteOrder(db *gorm.DB, order *models.Order) error {
	err := db.Delete(&order).Error
	if err != nil {
		return err
	}
	return nil
}
