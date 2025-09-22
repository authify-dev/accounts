package services

import (
	"accounts/internal/api/v1/policies/domain/commands"
	"accounts/internal/api/v1/policies/domain/entities"
	"accounts/internal/common/logger"
	"accounts/internal/utils"
	"context"
	"net/http"
)

func (s *PoliciesService) Create(ctx context.Context, command commands.CreatePolicyCommand) utils.Responses[entities.PolicyEntity] {

	entry := logger.FromContext(ctx)

	entity := entities.PolicyEntity{
		Name:           command.Name,
		Description:    command.Description,
		Resource:       command.Resource,
		Action:         command.Action,
		Effect:         string(command.Effect),
		OrganizationID: command.OrganizationID,
	}

	res := s.policies_repository.SaveEntity(entity)

	if res.Err != nil {
		entry.Error("Error creating policy", "error", res.Err)
		return utils.Responses[entities.PolicyEntity]{
			Err: res.Err,
		}
	}

	return utils.Responses[entities.PolicyEntity]{
		Body:       res.Data,
		StatusCode: http.StatusCreated,
		Success:    true,
	}
}
