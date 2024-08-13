package models

import (
	"gorm.io/gorm"
)

var db *gorm.DB // Initialize your database connection

type Lawyer struct {
	ID           int64  `json:"id" gorm:"primaryKey"`
	Name         string `json:"name"`
	Specialty    string `json:"specialty"`
	Email        string `json:"email"`
	MobileNumber string `json:"mobile_number"`
	Address      string `json:"address"`
}

// SetDB sets the database connection
func SetDB(database *gorm.DB) {
	db = database
}

// CreateLawyer creates a new lawyer in the database
func (lawyer *Lawyer) CreateLawyer() (*Lawyer, error) {
	result := db.Create(lawyer)
	if result.Error != nil {
		return nil, result.Error
	}
	return lawyer, nil
}

// GetAllLawyers retrieves all lawyers from the database
func GetAllLawyers() ([]Lawyer, error) {
	var lawyers []Lawyer
	result := db.Find(&lawyers)
	return lawyers, result.Error
}

// GetLawyerById retrieves a lawyer by its ID
func GetLawyerById(ID int64) (*Lawyer, error) {
	var lawyer Lawyer
	result := db.First(&lawyer, ID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &lawyer, nil
}

// DeleteLawyer deletes a lawyer by its ID
func DeleteLawyer(ID int64) (*Lawyer, error) {
	lawyer, err := GetLawyerById(ID)
	if err != nil {
		return nil, err
	}
	db.Delete(lawyer)
	return lawyer, nil
}

// UpdateLawyer updates a lawyer's details
func UpdateLawyer(ID int64, updatedLawyer *Lawyer) (*Lawyer, error) {
	lawyer, err := GetLawyerById(ID)
	if err != nil {
		return nil, err
	}
	if updatedLawyer.Name != "" {
		lawyer.Name = updatedLawyer.Name
	}
	if updatedLawyer.Specialty != "" {
		lawyer.Specialty = updatedLawyer.Specialty
	}
	if updatedLawyer.Email != "" {
		lawyer.Email = updatedLawyer.Email
	}
	if updatedLawyer.MobileNumber != "" {
		lawyer.MobileNumber = updatedLawyer.MobileNumber
	}
	if updatedLawyer.Address != "" {
		lawyer.Address = updatedLawyer.Address
	}
	db.Save(lawyer)
	return lawyer, nil
}
