package services

import (
	"accounts/internal/api/v1/policies/domain/commands"
	"accounts/internal/api/v1/policies/domain/entities"
	"accounts/internal/common/logger"
	"context"
	"foundation/utils"
	"foundation/utils/cerrs"
	"net/http"
)

func (s *PoliciesService) Create(ctx context.Context, command commands.CreatePolicyCommand) utils.Response[entities.PolicyEntity] {

	entry := logger.FromContext(ctx)

	entity := entities.PolicyEntity{
		Name:           command.Name,
		Description:    command.Description,
		Resource:       command.Resource,
		Action:         command.Action,
		Effect:         string(command.Effect),
		OrganizationID: command.OrganizationID,
	}

	res := s.policies_repository.Save(entity)

	if res.Err != nil {
		entry.Error("Error creating policy", "error", res.Err)
		return utils.Response[entities.PolicyEntity]{
			Error: cerrs.NewCustomError(http.StatusInternalServerError, "Error creating policy", "policies.create.error_creating_policy"),
		}
	}

	return utils.Response[entities.PolicyEntity]{
		Data:       res.Data,
		StatusCode: http.StatusCreated,
		Success:    true,
	}
}
