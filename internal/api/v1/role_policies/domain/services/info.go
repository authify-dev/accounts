package services

import (
	policies_entities "accounts/internal/api/v1/policies/domain/entities"
	"accounts/internal/api/v1/roles/domain/entities"
	"accounts/internal/common/logger"
	"foundation/domain/criteria"
	"foundation/domain/customctx"
	"foundation/utils"
	"foundation/utils/cerrs"
	"net/http"
)

type RolePoliciesInfoResponse struct {
	Role     entities.Role                    `json:"role"`
	Policies []policies_entities.PolicyEntity `json:"policies"`
}

func (s *RolePoliciesService) Info(cc *customctx.CustomContext, role_id string) utils.Response[RolePoliciesInfoResponse] {
	entry := logger.FromContext(cc.Context())

	entry.Info("Getting role policies info")

	role, err := s.role_repository.Search(role_id)

	if err != nil {
		entry.Error("Error getting role", "error", err)
		return utils.Response[RolePoliciesInfoResponse]{
			Error: cc.NewError(
				cerrs.NewCustomError(
					http.StatusInternalServerError, "Error getting role", "role_policies.info.error_getting_role",
				),
			),
		}
	}

	cri := criteria.Criteria{
		Filters: *criteria.NewFilters(
			[]criteria.Filter{
				{
					Field:    "role_id",
					Operator: criteria.OperatorEqual,
					Value:    role_id,
				},
			},
		),
	}

	role_policies, err := s.role_policies_repository.Matching(cri)

	if err != nil {
		entry.Error("Error getting policies", "error", err)
		return utils.Response[RolePoliciesInfoResponse]{
			Error: cc.NewError(
				cerrs.NewCustomError(
					http.StatusInternalServerError, "Error getting policies", "role_policies.info.error_getting_policies",
				),
			),
		}
	}

	var policies []policies_entities.PolicyEntity

	for _, role_policy := range role_policies {
		policy, err := s.policies_repository.Search(role_policy.PolicyID)
		if err != nil {
			entry.Error("Error getting policy", "error", err)
			return utils.Response[RolePoliciesInfoResponse]{
				Error: cc.NewError(
					cerrs.NewCustomError(
						http.StatusInternalServerError, "Error getting policy", "role_policies.info.error_getting_policy",
					),
				),
			}
		}

		policies = append(policies, policy)
	}

	return utils.Response[RolePoliciesInfoResponse]{
		Data: RolePoliciesInfoResponse{
			Role:     role,
			Policies: policies,
		},
		StatusCode: http.StatusOK,
		Success:    true,
	}

}
