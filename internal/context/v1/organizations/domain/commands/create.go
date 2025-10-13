package commands

import "accounts/internal/context/v1/organizations/domain/entities"

type CreateOrganizationCommand struct {
	Name       string
	RootUserID string
}

func (c CreateOrganizationCommand) ToEntity() entities.Organization {
	return entities.Organization{
		Name:       c.Name,
		RootUserID: c.RootUserID,
	}
}
