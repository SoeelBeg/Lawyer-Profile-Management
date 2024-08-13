package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/soeel/go-bookstore-lawyer/pkg/models"
	"github.com/soeel/go-bookstore-lawyer/pkg/utils"
)

// CreateLawyer creates a new lawyer in the database
func CreateLawyer(w http.ResponseWriter, r *http.Request) {
	var newLawyer models.Lawyer
	if err := utils.ParseBody(r, &newLawyer); err != nil {
		http.Error(w, "Unable to parse request body", http.StatusBadRequest)
		return
	}
	lawyer, err := newLawyer.CreateLawyer()
	if err != nil {
		http.Error(w, "Unable to create lawyer", http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(lawyer)
	if err != nil {
		http.Error(w, "Unable to marshal response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// GetLawyers retrieves all lawyers from the database
func GetLawyers(w http.ResponseWriter, r *http.Request) {
	lawyers, err := models.GetAllLawyers()
	if err != nil {
		http.Error(w, "Unable to fetch lawyers", http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(lawyers)
	if err != nil {
		http.Error(w, "Unable to marshal response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// GetLawyerById retrieves a lawyer by its ID
func GetLawyerById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lawyerId := vars["lawyerId"]
	ID, err := strconv.ParseInt(lawyerId, 10, 64)
	if err != nil {
		http.Error(w, "Invalid lawyer ID", http.StatusBadRequest)
		return
	}
	lawyerDetails, err := models.GetLawyerById(ID)
	if err != nil {
		http.Error(w, "Lawyer not found", http.StatusNotFound)
		return
	}
	res, err := json.Marshal(lawyerDetails)
	if err != nil {
		http.Error(w, "Unable to marshal response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// DeleteLawyer deletes a lawyer by its ID
func DeleteLawyer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	lawyerId := vars["lawyerId"]
	ID, err := strconv.ParseInt(lawyerId, 10, 64)
	if err != nil {
		http.Error(w, "Invalid lawyer ID", http.StatusBadRequest)
		return
	}
	lawyer, err := models.DeleteLawyer(ID)
	if err != nil {
		http.Error(w, "Unable to delete lawyer", http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(lawyer)
	if err != nil {
		http.Error(w, "Unable to marshal response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// UpdateLawyer updates a lawyer's details
func UpdateLawyer(w http.ResponseWriter, r *http.Request) {
	var updateLawyer models.Lawyer
	if err := utils.ParseBody(r, &updateLawyer); err != nil {
		http.Error(w, "Unable to parse request body", http.StatusBadRequest)
		return
	}
	vars := mux.Vars(r)
	lawyerId := vars["lawyerId"]
	ID, err := strconv.ParseInt(lawyerId, 10, 64)
	if err != nil {
		http.Error(w, "Invalid lawyer ID", http.StatusBadRequest)
		return
	}
	lawyerDetails, err := models.UpdateLawyer(ID, &updateLawyer)
	if err != nil {
		http.Error(w, "Unable to update lawyer", http.StatusInternalServerError)
		return
	}
	res, err := json.Marshal(lawyerDetails)
	if err != nil {
		http.Error(w, "Unable to marshal response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
