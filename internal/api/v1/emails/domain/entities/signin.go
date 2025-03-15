package entities

import "encoding/json"

type SignIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewEntityFromJSON[E any](jsonData []byte) (E, error) {
	var entity E
	err := json.Unmarshal(jsonData, &entity)
	return entity, err
}

type SignInResponse struct {
	JWT          string `json:"jwt"`
	RefreshToken string `json:"refresh_token"`
}
