package services

import (
	"accounts/internal/api/v1/policies/domain/entities"
	"foundation/domain/criteria"
	"foundation/domain/customctx"
	"foundation/utils"
	"foundation/utils/cerrs"
	"net/http"
)

func (s *PoliciesService) List(cc *customctx.CustomContext, organizationID string) utils.Response[entities.PolicyEntity] {

	cri := criteria.Criteria{
		Filters: *criteria.NewFilters(
			[]criteria.Filter{
				{
					Field:    "organization_id",
					Operator: criteria.OperatorEqual,
					Value:    organizationID,
				},
			},
		),
	}

	res, err := s.policies_repository.Matching(cri)

	if err != nil {
		return utils.Response[entities.PolicyEntity]{
			Error:      cc.NewError(cerrs.NewCustomError(http.StatusInternalServerError, "Error listing policies: "+err.Error(), "policies.list.error_listing_policies")),
			StatusCode: http.StatusInternalServerError,
			Success:    false,
		}
	}

	return utils.Response[entities.PolicyEntity]{
		Results:    res,
		Success:    true,
		StatusCode: http.StatusOK,
	}
}
