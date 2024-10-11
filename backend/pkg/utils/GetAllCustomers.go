package utils

import (
	"github.com/nexmonyx/nexmonyx/backend/pkg/models"
)

// GetAllCustomers returns all customers from the database
func GetAllCustomers() ([]models.Customer, error) {
	var customers []models.Customer
	result := DB.Find(&customers)
	if result.Error != nil {
		return nil, result.Error
	}
	return customers, nil
}
