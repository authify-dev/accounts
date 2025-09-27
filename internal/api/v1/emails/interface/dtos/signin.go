package dtos

import (
	"accounts/internal/api/v1/emails/domain/commands"
	"encoding/json"
)

type SignInDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (dto SignInDTO) Validate() error {
	return nil
}

func (dto SignInDTO) ToJson() []byte {

	data, err := json.Marshal(dto)
	if err != nil {
		return []byte{}
	}

	return data
}

func (dto SignInDTO) ToCommand() commands.SignIn {
	return commands.SignIn{
		Email:    dto.Email,
		Password: dto.Password,
	}
}
