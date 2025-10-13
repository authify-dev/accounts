package commands

import "accounts/internal/api/v1/roles/domain/entities"

type CreateRoleCommand struct {
	Name           string
	Description    string
	OrganizationID string
}

func (c CreateRoleCommand) ToEntity() entities.Role {
	return entities.Role{
		Name:           c.Name,
		Description:    c.Description,
		OrganizationID: c.OrganizationID,
	}
}
