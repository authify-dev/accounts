package policies_pg

import (
	"accounts/internal/api/v1/policies/domain/entities"
	policies_enums "accounts/internal/api/v1/policies/domain/enums"
	"accounts/internal/core/settings"
	"accounts/internal/db/postgres"
	organizations_gorm "accounts/internal/db/postgres/organizations"
	"foundation/types/uuidx"
	"time"

	"gorm.io/gorm"
)

const (
	ENTITY_POLICY_CODE = "02"
)

// PolicyModel representa el modelo de datos para la entidad Policy.
type PolicyModel struct {
	// Se asume que postgres.Model es un struct genérico que contiene campos comunes (como ID).
	postgres.Model[entities.PolicyEntity]

	Name           string                      `json:"name"`
	Description    string                      `json:"description,omitempty"`
	Resource       string                      `json:"resource"` // e.g. "user", "chat", "document"
	Action         string                      `json:"action"`   // e.g. "create", "read", "update", "delete"
	Effect         policies_enums.PolicyEffect `json:"effect"`   // "allow" | "deny"
	OrganizationID string                      `json:"organization_id" gorm:"type:varchar(50);not null"`

	Organization *organizations_gorm.OrganizationModel `gorm:"foreignKey:OrganizationID"`
}

// TableName especifica el nombre de la tabla en la base de datos.
func (PolicyModel) TableName() string {
	return "policies"
}

// GetID retorna el identificador único del modelo.
func (o PolicyModel) GetID() string {
	return o.ID
}

func (m *PolicyModel) BeforeCreate(tx *gorm.DB) (err error) {
	idx, err := uuidx.NewUUIDx(settings.Settings.UUID_MODULE, ENTITY_POLICY_CODE)
	if err != nil {
		return err
	}
	m.ID = idx.UUID().String()
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return m.Model.BeforeCreate(tx)
}
