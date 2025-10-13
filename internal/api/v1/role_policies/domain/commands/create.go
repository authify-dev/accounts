package commands

import "accounts/internal/api/v1/role_policies/domain/entities"

type CreateRolePoliciesCommand struct {
	RoleID         string `json:"role_id"`
	PolicyID       string `json:"policy_id"`
	OrganizationID string `json:"organization_id"`
}

func (c CreateRolePoliciesCommand) ToEntity() entities.RolePoliciesEntity {
	return entities.RolePoliciesEntity{
		RoleID:         c.RoleID,
		PolicyID:       c.PolicyID,
		OrganizationID: c.OrganizationID,
	}
}
