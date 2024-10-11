package utils

import (
	"github.com/nexmonyx/nexmonyx/backend/pkg/config"
)

// GetJWTSecret returns the JWT secret key
func GetJWTSecret() string {
	return config.CFG.JWTSecret
}
