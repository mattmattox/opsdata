package utils

import (
	"github.com/nexmonyx/nexmonyx/backend/pkg/models"
)

// SaveUploadLink saves the upload link to the database
func SaveUploadLink(link *models.UploadLink) error {
	return DB.Create(link).Error
}
