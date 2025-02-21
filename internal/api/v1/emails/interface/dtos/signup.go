package dtos

import (
	"encoding/json"
)

type SignUpDTO struct {
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Role      string `json:"role" binding:"required"`
	UserName  string `json:"user_name"`
	Name      string `json:"name"`
	Birthdate string `json:"birthdate"`
}

func (dto SignUpDTO) Validate() error {
	return nil
}

func (dto SignUpDTO) ToJson() ([]byte, error) {
	return json.Marshal(dto)
}
