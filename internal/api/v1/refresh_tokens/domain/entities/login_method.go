package entities

import (
	jwt_controller "accounts/internal/common/controllers"
	"accounts/internal/common/logger"
	"accounts/internal/core/domain"
	"accounts/internal/core/settings"
	"context"
	"encoding/json"
	"time"
)

// --------------------------------
// DOMAIN
// --------------------------------
// Refreshs Entity
// --------------------------------

// Refreshs embebe a Entity, por lo que automáticamente implementa domain.IEntity.
type RefreshToken struct {
	domain.Entity
	OrganizationID string    `json:"organization_id,omitempty"`
	UserID         string    `json:"user_id,omitempty"`
	LoginMethodID  string    `json:"login_method_id,omitempty"`
	ExternalID     string    `json:"external_id,omitempty"`
	ExpiresAt      time.Time `json:"expires_at,omitempty"`
	RemoveAt       time.Time `json:"remove_at,omitempty"`
}

func (r RefreshToken) ToJSON() map[string]interface{} {
	// Convertir el struct a JSON.
	data, err := json.Marshal(r)
	if err != nil {
		return nil
	}

	// Convertir los bytes JSON a un mapa.
	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil
	}

	return result
}

func (r RefreshToken) ToJWT(ctx context.Context, jwt_controller jwt_controller.JWTController) string {

	entry := logger.FromContext(ctx)
	entry.Info("Generating Refresh Token to JWT")

	refresh_map := r.ToJSON()

	delete(refresh_map, "updated_at")
	delete(refresh_map, "created_at")
	delete(refresh_map, "user_id")
	delete(refresh_map, "login_method_id")
	delete(refresh_map, "organization_id")

	jwt, err := jwt_controller.GenerateToken(ctx, refresh_map, settings.Settings.REFRESH_EXPIRE)
	if err != nil {
		entry.Error("Error generating Refresh Token to JWT: ", err)
		return ""
	}
	return jwt
}
