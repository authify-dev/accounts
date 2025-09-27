package services

import (
	"accounts/internal/api/v1/roles/domain/entities"
	"foundation/domain/criteria"
	"foundation/domain/customctx"
	"foundation/utils"
	"foundation/utils/cerrs"
	"net/http"
)

func (u *RolesService) List(cc *customctx.CustomContext, organizationID string) utils.Response[entities.Role] {
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

	res, err := u.repository.Matching(cri)

	if err != nil {
		return utils.Response[entities.Role]{
			Error:      cc.NewError(cerrs.NewCustomError(http.StatusInternalServerError, "Error listing roles: "+err.Error(), "roles.list.error_listing_roles")),
			StatusCode: http.StatusInternalServerError,
			Success:    false,
		}
	}
	return utils.Response[entities.Role]{
		Results:    res,
		Success:    true,
		StatusCode: http.StatusOK,
	}
}
