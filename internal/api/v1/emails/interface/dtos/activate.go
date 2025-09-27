package dtos

import (
	"accounts/internal/api/v1/emails/domain/commands"
	"encoding/json"
)

type ActivateDTO struct {
	Code  string `json:"code" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

func (dto ActivateDTO) Validate() error {
	return nil
}

func (dto ActivateDTO) ToJson() []byte {

	data, err := json.Marshal(dto)
	if err != nil {
		return []byte{}
	}

	return data
}

func (dto ActivateDTO) ToCommand() commands.ActivateCommand {
	return commands.ActivateCommand{
		Email:          dto.Email,
		Code:           dto.Code,
	}
}
