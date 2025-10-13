package commands

import "accounts/internal/api/v1/emails/domain/entities"

type ResetPassword struct {
	Email          string `json:"email"`
	OrganizationID string `json:"organization_id"`
}

func (c ResetPassword) ToEntity() entities.ResetPassword {
	return entities.ResetPassword{
		Email:          c.Email,
		OrganizationID: c.OrganizationID,
	}
}

type ConfirmPassword struct {
	Email          string `json:"email"`
	Code           string `json:"code"`
	Password       string `json:"password"`
	OrganizationID string `json:"organization_id"`
}

func (c ConfirmPassword) ToEntity() entities.ConfirmPassword {
	return entities.ConfirmPassword{
		Email:    c.Email,
		Code:     c.Code,
		Password: c.Password,
	}
}
