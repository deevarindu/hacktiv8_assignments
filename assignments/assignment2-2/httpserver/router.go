package httpserver

import (
	"github.com/deevarindu/hacktiv8_assignments/assignment2-2/httpserver/controllers"
	"github.com/gin-gonic/gin"
)

func CreateRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", controllers.HealthCheck)
	router.POST("/order", controllers.CreateOrder)
	router.GET("/orders", controllers.GetOrders)
	router.DELETE("/order/:id", controllers.DeleteOrder)

	return router
}
