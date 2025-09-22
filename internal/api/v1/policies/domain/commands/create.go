package commands

import policies_enums "accounts/internal/api/v1/policies/domain/enums"

type CreatePolicyCommand struct {
	Name           string                      `json:"name" binding:"required"`
	Description    string                      `json:"description,omitempty"`
	Resource       string                      `json:"resource" binding:"required"`
	Action         string                      `json:"action" binding:"required"`
	Effect         policies_enums.PolicyEffect `json:"effect" binding:"required"`
	OrganizationID string                      `json:"organization_id" binding:"required"`
}
