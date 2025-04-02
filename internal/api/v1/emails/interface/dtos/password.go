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

type ConfirmPasswordDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Code     string `json:"code" binding:"required,min=6,max=6"`
	Password string `json:"password" binding:"required,min=8,max=100"`
}

func (dto ConfirmPasswordDTO) Validate() error {
	return nil
}

func (dto ConfirmPasswordDTO) ToJson() []byte {
	data, err := json.Marshal(dto)
	if err != nil {
		return []byte{}
	}

	return data
}

func (dto ConfirmPasswordDTO) ToEntity() entities.ConfirmPassword {
	return entities.ConfirmPassword{
		Email:    dto.Email,
		Code:     dto.Code,
		Password: dto.Password,
	}
}
