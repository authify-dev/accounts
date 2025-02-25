package entities

import "encoding/json"

type SignUp struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func NewSingUpFromJSON(jsonData []byte) (SignUp, error) {
	var entity SignUp
	err := json.Unmarshal(jsonData, &entity)
	return entity, err
}

type SignUpResponse struct {
	JWT          string `json:"jwt"`
	RefreshToken string `json:"refresh_token"`
}
