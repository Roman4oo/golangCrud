package controllers

import (
	"encoding/json"
	"golangCrud/database"
	"golangCrud/entities"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var order entities.Order
	json.NewDecoder(r.Body).Decode(&order)
	database.Instance.Create(&order)
	json.NewEncoder(w).Encode(order)
}

func GetOrderById(w http.ResponseWriter, r *http.Request) {
	orderId := mux.Vars(r)["id"]
	if checkIforderExists(orderId) == false {
		json.NewEncoder(w).Encode("order Not Found!")
		return
	}
	var order entities.Order
	database.Instance.First(&order, orderId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	var orders []entities.Order
	database.Instance.Find(&orders)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(orders)
}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	orderId := mux.Vars(r)["id"]
	if checkIforderExists(orderId) == false {
		json.NewEncoder(w).Encode("order Not Found!")
		return
	}
	var order entities.Order
	database.Instance.First(&order, orderId)
	json.NewDecoder(r.Body).Decode(&order)
	database.Instance.Save(&order)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	orderId := mux.Vars(r)["id"]
	if checkIforderExists(orderId) == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("order Not Found!")
		return
	}
	var order entities.Order
	database.Instance.Delete(&order, orderId)
	json.NewEncoder(w).Encode("order Deleted Successfully!")
}

func checkIforderExists(orderId string) bool {
	var order entities.Order
	database.Instance.First(&order, orderId)
	if order.ID == 0 {
		return false
	}
	return true
}
