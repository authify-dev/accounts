package dtos

import (
	"accounts/internal/api/v1/roles/domain/commands"
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

type CreateRoleDTO struct {
	Name           string `json:"name" validate:"required"`
	Description    string `json:"description"`
	OrganizationID string `json:"organization_id" validate:"required"`
}

func (dto CreateRoleDTO) Validate() error {
	validate := validator.New()
	err := validate.Struct(dto)
	return err
}

func (dto CreateRoleDTO) ToJson() ([]byte, error) {
	return json.Marshal(dto)
}

func (dto CreateRoleDTO) ToCommand() commands.CreateRoleCommand {
	return commands.CreateRoleCommand{
		Name:           dto.Name,
		Description:    dto.Description,
		OrganizationID: dto.OrganizationID,
	}
}
