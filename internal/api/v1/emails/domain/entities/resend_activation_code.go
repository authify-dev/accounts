package entities

import "encoding/json"

type ResendActivationCode struct {
	Email string `json:"email"`
}

func NewResendActivationCodeFromJSON(jsonData []byte) (ResendActivationCode, error) {
	var entity ResendActivationCode
	err := json.Unmarshal(jsonData, &entity)
	return entity, err
}

type ResendActivationCodeResponse struct {
	Message string `json:"message"`
}
