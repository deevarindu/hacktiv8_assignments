package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/deevarindu/hacktiv8_assignments/assignment2/httpserver/config"
	"github.com/deevarindu/hacktiv8_assignments/assignment2/httpserver/models"
	"github.com/gorilla/mux"
)

var NewItem models.Item

func GetAllItems(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()
	var items []models.Item
	db.Find(&items)
	json.NewEncoder(w).Encode(items)
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()
	params := mux.Vars(r)
	var item models.Item
	db.First(&item, params["id"])
	json.NewEncoder(w).Encode(item)
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()
	json.NewDecoder(r.Body).Decode(&NewItem)
	db.Create(&NewItem)
	json.NewEncoder(w).Encode(NewItem)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()
	params := mux.Vars(r)
	var item models.Item
	db.First(&item, params["id"])
	json.NewDecoder(r.Body).Decode(&item)
	db.Save(&item)
	json.NewEncoder(w).Encode(item)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()
	params := mux.Vars(r)
	var item models.Item
	db.Delete(&item, params["id"])
	json.NewEncoder(w).Encode(item)
}