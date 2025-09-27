package entities

import "encoding/json"

type SignUp struct {
	UserName       string `json:"user_name"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	Role           string `json:"role"`
	OrganizationID string `json:"organization_id"`
}

func NewSingUpFromJSON(jsonData []byte) (SignUp, error) {
	var entity SignUp
	err := json.Unmarshal(jsonData, &entity)
	return entity, err
}

type SignUpResponse struct {
	Message string `json:"message"`
}
