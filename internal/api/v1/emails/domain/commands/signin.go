package commands

import "accounts/internal/api/v1/emails/domain/entities"

type SignIn struct {
	Email          string `json:"email"`
	Password       string `json:"password"`
	OrganizationID string `json:"organization_id"`
}

func (c SignIn) ToEntity() entities.SignIn {
	return entities.SignIn{
		Email:    c.Email,
		Password: c.Password,
	}
}
