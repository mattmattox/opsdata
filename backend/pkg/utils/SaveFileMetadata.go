package utils

import (
	"github.com/nexmonyx/nexmonyx/backend/pkg/models"
)

// SaveFileMetadata saves file metadata to the database
func SaveFileMetadata(file *models.FileMetadata) error {
	return DB.Create(file).Error
}
