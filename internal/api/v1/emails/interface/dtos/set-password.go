package dtos

import "accounts/internal/api/v1/emails/domain/commands"

type SetPasswordDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (dto SetPasswordDTO) Validate() error {
	return nil
}

func (dto SetPasswordDTO) ToCommand() commands.SetPassword {
	return commands.SetPassword{
		Email:    dto.Email,
		Password: dto.Password,
	}
}
