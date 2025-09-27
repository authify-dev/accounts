package commands

import "accounts/internal/context/v1/api_keys/domain/entities"

type GenerateAPIKeysCommand struct {
	Name           string
	OrganizationID string
	Description    string
}

func (c GenerateAPIKeysCommand) ToEntity() entities.APIKeyEntity {
	return entities.APIKeyEntity{
		Name:           c.Name,
		OrganizationID: c.OrganizationID,
		Description:    c.Description,
	}
}
