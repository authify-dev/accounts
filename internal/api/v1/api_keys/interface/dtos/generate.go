package dtos

import (
	"accounts/internal/context/v1/api_keys/domain/commands"
)

type GenerateAPIKeysDTO struct {
	Name           string `json:"name" binding:"required"`
	OrganizationID string `json:"organization_id" binding:"required"`
	Description    string `json:"description" binding:"omitempty"`
}

func (dto GenerateAPIKeysDTO) Validate() error {
	return nil
}

func (dto *GenerateAPIKeysDTO) ToCommand() commands.GenerateAPIKeysCommand {
	return commands.GenerateAPIKeysCommand{
		Name:           dto.Name,
		OrganizationID: dto.OrganizationID,
		Description:    dto.Description,
	}
}
