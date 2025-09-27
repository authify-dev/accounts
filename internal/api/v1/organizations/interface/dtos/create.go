package dtos

import "accounts/internal/context/v1/organizations/domain/commands"

type CreateOrganizationDTO struct {
	Name       string `json:"name" binding:"required"`
	RootUserID string `json:"root_user_id" binding:"omitempty"`
}

func (dto CreateOrganizationDTO) Validate() error {
	return nil
}

func (dto CreateOrganizationDTO) ToCommand() commands.CreateOrganizationCommand {
	return commands.CreateOrganizationCommand{
		Name:       dto.Name,
		RootUserID: dto.RootUserID,
	}
}
