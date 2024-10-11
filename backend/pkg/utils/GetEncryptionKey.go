package utils

import (
	"github.com/nexmonyx/nexmonyx/backend/pkg/config"
)

// GetEncryptionKey returns the encryption key
func GetEncryptionKey() string {
	return config.CFG.EncryptionKey
}
