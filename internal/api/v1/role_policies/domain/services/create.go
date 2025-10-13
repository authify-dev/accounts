package services

import (
	"accounts/internal/api/v1/role_policies/domain/commands"
	"accounts/internal/api/v1/role_policies/domain/entities"
	"accounts/internal/common/logger"
	"foundation/domain/criteria"
	"foundation/domain/customctx"
	"foundation/utils"
	"foundation/utils/cerrs"
	"net/http"
)

func (s *RolePoliciesService) Create(cc *customctx.CustomContext, command commands.CreateRolePoliciesCommand) utils.Response[entities.RolePoliciesEntity] {
	entry := logger.FromContext(cc.Context())

	cri := criteria.Criteria{
		Filters: *criteria.NewFilters(
			[]criteria.Filter{
				{
					Field:    "organization_id",
					Operator: criteria.OperatorEqual,
					Value:    command.OrganizationID,
				},
				{
					Field:    "id",
					Operator: criteria.OperatorEqual,
					Value:    command.RoleID,
				},
			},
		),
	}

	roles, err := s.role_repository.Matching(cri)

	if err != nil {
		entry.Error("Error getting role", "error", err)
		return utils.Response[entities.RolePoliciesEntity]{
			Error: cc.NewError(
				cerrs.NewCustomError(
					http.StatusInternalServerError,
					"Error getting role",
					"role_policies.create.error_getting_role",
				),
			),
		}
	}

	if len(roles) == 0 {
		entry.Error("Role not found")
		return utils.Response[entities.RolePoliciesEntity]{
			Error: cc.NewError(
				cerrs.NewCustomError(
					http.StatusNotFound,
					"Role not found",
					"role_policies.create.role_not_found",
				),
			),
		}
	}

	role := roles[0]

	if role.OrganizationID != command.OrganizationID {
		entry.Error("Role not found in organization")
		return utils.Response[entities.RolePoliciesEntity]{
			Error: cc.NewError(
				cerrs.NewCustomError(
					http.StatusNotFound,
					"Role not found in organization",
					"role_policies.create.role_not_found_in_organization",
				),
			),
		}
	}

	cri = criteria.Criteria{
		Filters: *criteria.NewFilters(
			[]criteria.Filter{
				{
					Field:    "organization_id",
					Operator: criteria.OperatorEqual,
					Value:    command.OrganizationID,
				},
				{
					Field:    "id",
					Operator: criteria.OperatorEqual,
					Value:    command.PolicyID,
				},
			},
		),
	}

	policies, err := s.policies_repository.Matching(cri)

	if err != nil {
		entry.Error("Error getting policy", "error", err)
		return utils.Response[entities.RolePoliciesEntity]{
			Error: cc.NewError(
				cerrs.NewCustomError(
					http.StatusInternalServerError,
					"Error getting policy",
					"role_policies.create.error_getting_policy",
				),
			),
		}
	}

	if len(policies) == 0 {
		entry.Error("Policy not found")
		return utils.Response[entities.RolePoliciesEntity]{
			Error: cc.NewError(
				cerrs.NewCustomError(
					http.StatusNotFound,
					"Policy not found",
					"role_policies.create.policy_not_found",
				),
			),
		}
	}

	policy := policies[0]

	if policy.OrganizationID != command.OrganizationID {
		entry.Error("Policy not found in organization")
		return utils.Response[entities.RolePoliciesEntity]{
			Error: cc.NewError(
				cerrs.NewCustomError(
					http.StatusNotFound,
					"Policy not found in organization",
					"role_policies.create.policy_not_found_in_organization",
				),
			),
		}
	}

	entity := entities.RolePoliciesEntity{
		RoleID:         command.RoleID,
		PolicyID:       command.PolicyID,
		OrganizationID: command.OrganizationID,
	}

	res := s.role_policies_repository.Save(entity)

	if res.Err != nil {
		entry.Error("Error creating policy", "error", res.Err)
		return utils.Response[entities.RolePoliciesEntity]{
			Error: cc.NewError(
				cerrs.NewCustomError(
					http.StatusInternalServerError,
					"Error creating policy",
					"role_policies.create.error_creating_policy",
				),
			),
		}
	}

	return utils.Response[entities.RolePoliciesEntity]{
		Data:       res.Data,
		StatusCode: http.StatusCreated,
		Success:    true,
	}
}
