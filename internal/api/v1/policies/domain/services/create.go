package services

import (
	"accounts/internal/api/v1/policies/domain/commands"
	"accounts/internal/api/v1/policies/domain/entities"
	"accounts/internal/common/logger"
	"foundation/domain/customctx"
	"foundation/utils"
	"foundation/utils/cerrs"
	"net/http"
)

func (s *PoliciesService) Create(cc *customctx.CustomContext, command commands.CreatePolicyCommand) utils.Response[entities.PolicyEntity] {

	entry := logger.FromContext(cc.Context())

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
			Error: cc.NewError(
				cerrs.NewCustomError(
					http.StatusInternalServerError,
					"Error creating policy: "+res.Err.Error(),
					"policies.create.error_creating_policy",
				),
			),
			StatusCode: res.Err.GetCode(),
			Success:    false,
		}
	}

	return utils.Response[entities.PolicyEntity]{
		Data:       res.Data,
		StatusCode: http.StatusCreated,
		Success:    true,
	}
}
