package entities

import (
	jwt_controller "accounts/internal/common/controllers"
	"accounts/internal/core/domain"
	"accounts/internal/core/settings"
	"encoding/json"
)

// --------------------------------
// DOMAIN
// --------------------------------
// LoginMethods Entity
// --------------------------------

// LoginMethods embebe a Entity, por lo que automáticamente implementa domain.IEntity.
type LoginMethod struct {
	domain.Entity
	UserID     string `json:"user_id,omitempty"`
	EntityID   string `json:"entity_id,omitempty"`
	EntityType string `json:"entity_type,omitempty"`
	IsActive   bool   `json:"is_active,omitempty"`
	IsVerify   bool   ` json:"is_verify,omitempty"`
}

func (r LoginMethod) ToJSON() map[string]interface{} {
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

func (r LoginMethod) ToJWT(jwt_controller jwt_controller.JWTController) string {
	login_map := r.ToJSON()

	delete(login_map, "updated_at")
	delete(login_map, "created_at")

	jwt, err := jwt_controller.GenerateToken(login_map, settings.Settings.JWT_EXPIRE)
	if err != nil {
		return ""
	}
	return jwt
}
