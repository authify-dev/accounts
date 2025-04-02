package dtos

import (
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
