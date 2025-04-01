package dtos

import (
	"accounts/internal/api/v1/emails/domain/entities"
	"encoding/json"
)

type ResetPasswordDTO struct {
	Email string `json:"email" binding:"required,email"`
}

func (dto ResetPasswordDTO) Validate() error {
	return nil
}

func (dto ResetPasswordDTO) ToJson() []byte {
	data, err := json.Marshal(dto)
	if err != nil {
		return []byte{}
	}

	return data
}

func (dto ResetPasswordDTO) ToEntity() entities.ResetPassword {
	return entities.ResetPassword{
		Email: dto.Email,
	}
}
