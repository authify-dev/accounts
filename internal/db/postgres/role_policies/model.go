package rolepolicies_pg

import (
	"accounts/internal/api/v1/role_policies/domain/entities"
	"accounts/internal/core/settings"
	"accounts/internal/db/postgres"
	postgres_policies "accounts/internal/db/postgres/policies"
	postgres_roles "accounts/internal/db/postgres/role"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// --------------------------------
// INFRASTRUCTURE
// --------------------------------
// Role Policies Model
// --------------------------------

// RolePoliciesModel utiliza Model parametrizado con RolePolicies.
type RolePoliciesModel struct {
	postgres.Model[entities.RolePoliciesEntity]
	RoleID   string `gorm:"type:varchar(50);not null" json:"role_id"`
	PolicyID string `gorm:"type:varchar(50);not null" json:"policy_id"`

	RoleModel   postgres_roles.RoleModel      `gorm:"foreignKey:RoleID;references:ID" json:"role"`
	PolicyModel postgres_policies.PolicyModel `gorm:"foreignKey:PolicyID;references:ID" json:"policy"`
}

func (RolePoliciesModel) TableName() string {
	if settings.Settings.DB_SCHEMA == "" {
		return "role_policies"
	}
	return settings.Settings.DB_SCHEMA + ".role_policies"
}

func (c RolePoliciesModel) GetID() string {
	return c.ID
}

func (m *RolePoliciesModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = fmt.Sprintf("%s_%s", m.TableName()[:3], uuid.New().String())
	return m.Model.BeforeCreate(tx)
}
