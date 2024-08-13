package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/soeel/go-bookstore-lawyer/pkg/controller"
	"github.com/soeel/go-bookstore-lawyer/pkg/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Initialize the database connection
	dsn := "admin:soeel123@tcp(127.0.0.1:3306)/soeeldb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	log.Println("Database connection established")

	// Set the global database connection
	models.SetDB(db)

	// Initialize the router
	router := mux.NewRouter()

	// Define routes
	router.HandleFunc("/lawyers", controller.GetLawyers).Methods("GET")
	router.HandleFunc("/lawyers/{lawyerId:[0-9]+}", controller.GetLawyerById).Methods("GET")
	router.HandleFunc("/lawyers", controller.CreateLawyer).Methods("POST")
	router.HandleFunc("/lawyers/{lawyerId:[0-9]+}", controller.UpdateLawyer).Methods("PUT")
	router.HandleFunc("/lawyers/{lawyerId:[0-9]+}", controller.DeleteLawyer).Methods("DELETE")

	// Start the server
	port := "8080"
	log.Printf("Server started on port %s", port)
	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
