package main

import (
	"fmt"
	"golang-crud-rest-api/controllers"
	"golang-crud-rest-api/database"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {

	// Load Configurations from config.json using Viper
	LoadAppConfig()

	// Initialize Database
	database.Connect(AppConfig.ConnectionString)
	database.Migrate()
	
	// Initialize the router
	router := mux.NewRouter().StrictSlash(true)

	// Register Routes
	RegisterordertRoutes(router)

	// Start the server
	log.Println(fmt.Sprintf("Starting Server on port %s", AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port), router))
}

func RegisterordertRoutes(router *mux.Router) {
	router.HandleFunc("/api/orderts", controllers.Getorderts).Methods("GET")
	router.HandleFunc("/api/orderts/{id}", controllers.GetordertById).Methods("GET")
	router.HandleFunc("/api/orderts", controllers.Createordert).Methods("POST")
	router.HandleFunc("/api/orderts/{id}", controllers.Updateordert).Methods("PUT")
	router.HandleFunc("/api/orderts/{id}", controllers.Deleteordert).Methods("DELETE")
}
