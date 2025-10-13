package services

import (
	"accounts/internal/api/v1/emails/domain/entities"
	"accounts/internal/common/logger"
	"context"
	"foundation/domain/criteria"
	"foundation/utils"
	"foundation/utils/cerrs"
	"net/http"
)

func (s EmailsService) ActivateV2(ctx context.Context, entity entities.Activate) utils.Response[entities.ActivateV2Response] {

	entry := logger.FromContext(ctx)

	entry.Info("Activating user", "email", entity.Email)

	// buscamos el pending registration por el email

	criteria_email := criteria.Criteria{
		Filters: *criteria.NewFilters(
			[]criteria.Filter{
				{
					Field:    "email",
					Value:    entity.Email,
					Operator: criteria.OperatorEqual,
				},
			},
		),
	}

	pending_registrations, err := s.pending_registrations_repository.Matching(criteria_email)
	if err != nil {
		return utils.Response[entities.ActivateV2Response]{
			Error:      cerrs.NewCustomError(http.StatusInternalServerError, "error getting email", "emails.activate_v2.error_getting_email"),
			StatusCode: 500,
		}
	}

	if len(pending_registrations) == 0 {
		return utils.Response[entities.ActivateV2Response]{
			Error:      cerrs.NewCustomError(http.StatusNotFound, "pending registration not found", "emails.activate_v2.error_pending_registration_not_found"),
			StatusCode: 404,
		}
	}

	pending_registration := pending_registrations[0]

	// buscamos el code por el code_id

	criteria_code := criteria.Criteria{
		Filters: *criteria.NewFilters(
			[]criteria.Filter{
				{
					Field:    "id",
					Value:    pending_registration.CodeID,
					Operator: criteria.OperatorEqual,
				},
			},
		),
	}

	codes, err := s.codes_repository.Matching(criteria_code)
	if err != nil {
		return utils.Response[entities.ActivateV2Response]{
			Error:      cerrs.NewCustomError(http.StatusInternalServerError, "error getting codes", "emails.activate_v2.error_getting_codes"),
			StatusCode: 500,
		}
	}

	if len(codes) == 0 {
		return utils.Response[entities.ActivateV2Response]{
			Error:      cerrs.NewCustomError(http.StatusNotFound, "code not found", "emails.activate_v2.error_code_not_found"),
			StatusCode: 404,
		}
	}

	code := codes[0]

	// verificamos que sean iguales
	if code.Code != entity.Code {
		return utils.Response[entities.ActivateV2Response]{
			Error:      cerrs.NewCustomError(http.StatusBadRequest, "code not valid", "emails.activate_v2.error_code_not_valid"),
			StatusCode: 400,
		}
	}

	// crear token con expiracion de 15 minutos
	token, err := s.jwt_controller.GenerateToken(
		ctx,
		map[string]interface{}{
			"user_id": code.UserID,
			"id":      pending_registration.ID,
		},
		15*60,
	)

	if err != nil {
		return utils.Response[entities.ActivateV2Response]{
			Error:      cerrs.NewCustomError(http.StatusInternalServerError, "error generating token", "emails.activate_v2.error_generating_token"),
			StatusCode: 500,
		}
	}

	return utils.Response[entities.ActivateV2Response]{
		StatusCode: 200,
		Data:       entities.ActivateV2Response{JWT: token},
		Success:    true,
	}

}
