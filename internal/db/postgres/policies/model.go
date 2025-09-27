package policies_pg

import (
	"accounts/internal/api/v1/policies/domain/entities"
	policies_enums "accounts/internal/api/v1/policies/domain/enums"
	"accounts/internal/core/settings"
	organizations_gorm "accounts/internal/db/postgres/organinizations"
	"foundation/infrastructure/db/cgorm"
	"foundation/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	ENTITY_POLICY_CODE = "02"
)

// PolicyModel representa el modelo de datos para la entidad Policy.
type PolicyModel struct {
	cgorm.Model[entities.PolicyEntity]

	Name        string                      `json:"name"`
	Description string                      `json:"description,omitempty"`
	Resource    string                      `gorm:"type:varchar(255);not null;uniqueIndex:uniq_policy_org_resource_action_effect,priority:2" json:"resource"`
	Action      string                      `gorm:"type:varchar(255);not null;uniqueIndex:uniq_policy_org_resource_action_effect,priority:3" json:"action"`
	Effect      policies_enums.PolicyEffect `gorm:"type:varchar(10);not null;uniqueIndex:uniq_policy_org_resource_action_effect,priority:4" json:"effect"`

	OrganizationID string `gorm:"type:varchar(50);not null;uniqueIndex:uniq_policy_org_resource_action_effect,priority:1" json:"organization_id"`

	Organization *organizations_gorm.OrganizationModel `gorm:"foreignKey:OrganizationID"`
}

// TableName especifica el nombre de la tabla en la base de datos.
func (PolicyModel) TableName() string {
	if settings.Settings.DB_SCHEMA == "" {
		return "policies"
	}
	return settings.Settings.DB_SCHEMA + ".policies"
}

// GetID retorna el identificador único del modelo.
func (o PolicyModel) GetID() uuid.UUID {
	return o.ID
}

func (m *PolicyModel) BeforeCreate(tx *gorm.DB) (err error) {
	idx, err := utils.NewUUIDx(settings.Settings.UUID_MODULE, ENTITY_POLICY_CODE)
	if err != nil {
		return err
	}
	m.ID = idx.UUID()
	return m.Model.BeforeCreate(tx)
}
