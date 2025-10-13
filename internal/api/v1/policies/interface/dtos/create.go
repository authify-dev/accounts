package dtos

import (
	"accounts/internal/api/v1/policies/domain/commands"
	policies_enums "accounts/internal/api/v1/policies/domain/enums"
)

type CreatePolicyDTO struct {
	Name        string                      `json:"name" binding:"required"`
	Description string                      `json:"description,omitempty"`
	Resource    string                      `json:"resource" binding:"required"`
	Action      string                      `json:"action" binding:"required"`
	Effect      policies_enums.PolicyEffect `json:"effect" binding:"required"`
}

func (dto CreatePolicyDTO) Validate() error {
	return nil
}

func (dto CreatePolicyDTO) ToCommand() commands.CreatePolicyCommand {
	return commands.CreatePolicyCommand{
		Name:        dto.Name,
		Description: dto.Description,
		Resource:    dto.Resource,
		Action:      dto.Action,
		Effect:      dto.Effect,
	}
}
