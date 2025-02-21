package entities

import "accounts/internal/core/domain"

type Role struct {
	domain.Entity
	Name        string `json:"name"`
	Description string `json:"description"`
}
