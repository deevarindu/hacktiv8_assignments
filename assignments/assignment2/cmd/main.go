package main

import (
	"net/http"

	"github.com/deevarindu/hacktiv8_assignments/assignment2/httpserver/controllers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/orders", controllers.GetAllOrders).Methods("GET")
	router.HandleFunc("/orders/{id}", controllers.GetOrder).Methods("GET")
	router.HandleFunc("/orders", controllers.CreateOrder).Methods("POST")
	router.HandleFunc("/orders/{id}", controllers.UpdateOrder).Methods("PUT")
	router.HandleFunc("/orders/{id}", controllers.DeleteOrder).Methods("DELETE")

	router.HandleFunc("/orders/{id}/items", controllers.GetAllItems).Methods("GET")
	router.HandleFunc("/orders/{id}/items/{item_id}", controllers.GetItem).Methods("GET")
	router.HandleFunc("/orders/{id}/items", controllers.CreateItem).Methods("POST")
	router.HandleFunc("/orders/{id}/items/{item_id}", controllers.UpdateItem).Methods("PUT")
	router.HandleFunc("/orders/{id}/items/{item_id}", controllers.DeleteItem).Methods("DELETE")

	http.Handle("/", router)
	http.ListenAndServe(":8080", router)
}
