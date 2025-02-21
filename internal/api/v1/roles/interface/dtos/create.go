package dtos

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
)

type CreateRoleDTO struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}

func (dto CreateRoleDTO) Validate() error {
	validate := validator.New()
	err := validate.Struct(dto)
	return err
}

func (dto CreateRoleDTO) ToJson() ([]byte, error) {
	return json.Marshal(dto)
}
