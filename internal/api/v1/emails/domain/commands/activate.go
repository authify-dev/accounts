package commands

import "accounts/internal/api/v1/emails/domain/entities"

type ActivateCommand struct {
	Email          string `json:"email"`
	Code           string `json:"code"`
	OrganizationID string `json:"organization_id"`
}

func (c *ActivateCommand) ToEntity() entities.Activate {
	return entities.Activate{
		Email:          c.Email,
		Code:           c.Code,
		OrganizationID: c.OrganizationID,
	}
}
