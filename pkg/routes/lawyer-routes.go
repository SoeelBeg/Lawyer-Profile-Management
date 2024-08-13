package routes

import (
	"github.com/gorilla/mux"
	"github.com/soeel/go-bookstore-lawyer/pkg/controller"
)

var RegisterLawyerRoutes = func(router *mux.Router) {
	router.HandleFunc("/lawyers", controller.CreateLawyer).Methods("POST")
	router.HandleFunc("/lawyers", controller.GetLawyers).Methods("GET")
	router.HandleFunc("/lawyers/{lawyerId}", controller.GetLawyerById).Methods("GET")
	router.HandleFunc("/lawyers/{lawyerId}", controller.UpdateLawyer).Methods("PUT")
	router.HandleFunc("/lawyers/{lawyerId}", controller.DeleteLawyer).Methods("DELETE")
}
