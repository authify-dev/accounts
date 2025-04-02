package entities

import "encoding/json"

type Activate struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

func NewActivateFromJSON(jsonData []byte) (Activate, error) {
	var entity Activate
	err := json.Unmarshal(jsonData, &entity)
	return entity, err
}

type ActivateResponse struct {
	JWT          string `json:"jwt"`
	RefreshToken string `json:"refresh_token"`
}
