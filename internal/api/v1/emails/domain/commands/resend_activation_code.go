package commands

import "accounts/internal/api/v1/emails/domain/entities"

type ResendActivationCodeCommand struct {
	Email          string `json:"email"`
	OrganizationID string `json:"organization_id"`
}

func (c *ResendActivationCodeCommand) ToEntity() entities.ResendActivationCode {
	return entities.ResendActivationCode{
		Email: c.Email,
	}
}
