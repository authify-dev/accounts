package dtos

import (
	"accounts/internal/api/v1/emails/domain/commands"
	"encoding/json"
)

type ResendActivationCodeDTO struct {
	Email string `json:"email" binding:"required,email"`
}

func (dto ResendActivationCodeDTO) Validate() error {
	return nil
}

func (dto ResendActivationCodeDTO) ToJson() []byte {

	data, err := json.Marshal(dto)
	if err != nil {
		return []byte{}
	}

	return data
}

func (dto ResendActivationCodeDTO) ToCommand() commands.ResendActivationCodeCommand {
	return commands.ResendActivationCodeCommand{
		Email: dto.Email,
	}
}
