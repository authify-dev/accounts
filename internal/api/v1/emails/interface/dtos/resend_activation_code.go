package dtos

import (
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
