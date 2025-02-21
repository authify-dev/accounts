package entities

import "accounts/internal/core/domain"

type User struct {
	domain.Entity
	UserName  string `json:"username"`
	Name      string `json:"name"`
	Birthdate string `json:"birthdate"`
	RoleID    string `json:"role"`
}
