package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/deevarindu/hacktiv8_assignments/assignment2/httpserver/config"
	"github.com/deevarindu/hacktiv8_assignments/assignment2/httpserver/models"
	"github.com/gorilla/mux"
)

var NewOrder models.Order

func GetAllOrders(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()
	var orders []models.Order
	db.Find(&orders)
	json.NewEncoder(w).Encode(orders)
}

func GetOrder(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()
	params := mux.Vars(r)
	var order models.Order
	db.First(&order, params["id"])
	json.NewEncoder(w).Encode(order)
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()
	json.NewDecoder(r.Body).Decode(&NewOrder)
	db.Create(&NewOrder)
	json.NewEncoder(w).Encode(NewOrder)
}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()
	params := mux.Vars(r)
	var order models.Order
	db.First(&order, params["id"])
	json.NewDecoder(r.Body).Decode(&order)
	db.Save(&order)
	json.NewEncoder(w).Encode(order)
}

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()
	params := mux.Vars(r)
	var order models.Order
	db.Delete(&order, params["id"])
	json.NewEncoder(w).Encode(order)
}
