package utils

import (
	"github.com/nexmonyx/nexmonyx/backend/pkg/models"
)

// DeleteUser deletes a user from the database
func DeleteUser(id uint) error {
	return DB.Delete(&models.User{}, id).Error
}
