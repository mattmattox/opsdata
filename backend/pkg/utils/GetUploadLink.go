package utils

import (
	"github.com/nexmonyx/nexmonyx/backend/pkg/models"
)

// GetUploadLink retrieves an upload link from the database
func GetUploadLink(linkID string) (*models.UploadLink, error) {
	var link models.UploadLink
	result := DB.Where("id = ?", linkID).First(&link)
	if result.Error != nil {
		return nil, result.Error
	}
	return &link, nil
}
