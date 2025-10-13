package dtos

import "accounts/internal/api/v1/pending_registrations/domain/commands"

type CreatePendingRegistrationDTO struct {
	Email    string `json:"email" binding:"required"`
	UserName string `json:"user_name,omitempty"`
	Role     string `json:"role" binding:"required"`
}

func (dto CreatePendingRegistrationDTO) Validate() error {
	return nil
}

func (dto CreatePendingRegistrationDTO) ToCommand() commands.CreatePendingRegistrationCommand {
	return commands.CreatePendingRegistrationCommand{
		Email:    dto.Email,
		UserName: dto.UserName,
		Role:     dto.Role,
	}
}
