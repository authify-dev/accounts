package dtos

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

type CreateUserDTO struct {
	Name     string `json:"name" validate:"required"`
	UserName string `json:"user_name" validate:"required"`
	Role     string `json:"role" validate:"required"`
}

func (dto CreateUserDTO) Validate() error {
	validate := validator.New()
	err := validate.Struct(dto)
	return err
}

func (dto CreateUserDTO) ToJson() ([]byte, error) {
	return json.Marshal(dto)
}
