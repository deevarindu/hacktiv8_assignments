package controllers

import (
	"net/http"

	"github.com/deevarindu/hacktiv8_assignments/assignment2-2/httpserver/repositories/models"
	"github.com/deevarindu/hacktiv8_assignments/assignment2-2/httpserver/services"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func CreateOrder(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var order models.Order
	c.BindJSON(&order)

	err := services.CreateOrder(db, &order)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}

func GetOrders(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var orders []models.Order
	err := services.GetOrders(db, &orders)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, orders)
	}
}

func DeleteOrder(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var order models.Order
	c.BindJSON(&order)

	err := services.DeleteOrder(db, &order)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}
