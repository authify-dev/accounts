package dtos

import (
	"encoding/json"
)

type SignUpDTO struct {
	UserName string `json:"user_name"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

func (dto SignUpDTO) Validate() error {
	return nil
}

func (dto SignUpDTO) ToJson() []byte {

	data, err := json.Marshal(dto)
	if err != nil {
		return []byte{}
	}

	return data
}
