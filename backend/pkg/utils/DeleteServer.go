package utils

import (
	"github.com/nexmonyx/nexmonyx/backend/pkg/models"
)

// DeleteServer deletes a server from the database
func DeleteServer(id uint) error {
	return DB.Delete(&models.Server{}, id).Error
}
