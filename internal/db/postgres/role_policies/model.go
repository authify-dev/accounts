package rolepolicies_pg

import (
	"accounts/internal/api/v1/role_policies/domain/entities"
	"accounts/internal/core/settings"
	organizations_gorm "accounts/internal/db/postgres/organinizations"
	postgres_policies "accounts/internal/db/postgres/policies"
	postgres_roles "accounts/internal/db/postgres/role"
	"foundation/infrastructure/db/cgorm"
	"foundation/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	ENTITY_ROLE_POLICIES_CODE = "04"
)

// --------------------------------
// INFRASTRUCTURE
// --------------------------------
// Role Policies Model
// --------------------------------

// RolePoliciesModel utiliza Model parametrizado con RolePolicies.
type RolePoliciesModel struct {
	cgorm.Model[entities.RolePoliciesEntity]
	RoleID         string `gorm:"type:varchar(50);not null" json:"role_id"`
	PolicyID       string `gorm:"type:varchar(50);not null" json:"policy_id"`
	OrganizationID string `gorm:"type:varchar(50);not null" json:"organization_id"`

	Organization *organizations_gorm.OrganizationModel `gorm:"foreignKey:OrganizationID"`
	RoleModel    postgres_roles.RoleModel              `gorm:"foreignKey:RoleID;references:ID" json:"role"`
	PolicyModel  postgres_policies.PolicyModel         `gorm:"foreignKey:PolicyID;references:ID" json:"policy"`
}

func (RolePoliciesModel) TableName() string {
	if settings.Settings.DB_SCHEMA == "" {
		return "role_policies"
	}
	return settings.Settings.DB_SCHEMA + ".role_policies"
}

func (c RolePoliciesModel) GetID() uuid.UUID {
	return c.ID
}

func (m *RolePoliciesModel) BeforeCreate(tx *gorm.DB) (err error) {
	idx, err := utils.NewUUIDx(settings.Settings.UUID_MODULE, ENTITY_ROLE_POLICIES_CODE)
	if err != nil {
		return err
	}
	m.ID = idx.UUID()
	return m.Model.BeforeCreate(tx)
}
