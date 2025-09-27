package usecases

import (
	"accounts/internal/context/v1/api_keys/domain/entities"
	"accounts/internal/context/v1/api_keys/domain/repositories"
	"foundation/domain/criteria"
	"foundation/domain/customctx"
	"foundation/domain/logger"
	"foundation/utils"
	"foundation/utils/cerrs"
	"net/http"
)

type ValidateAPIKeyUseCase struct {
	apiKeyRepository repositories.APIKeyRepository
}

func NewValidateAPIKeyUseCase(apiKeyRepository repositories.APIKeyRepository) *ValidateAPIKeyUseCase {
	return &ValidateAPIKeyUseCase{apiKeyRepository: apiKeyRepository}
}

func (u *ValidateAPIKeyUseCase) Validate(cc *customctx.CustomContext, apiKey string) utils.Response[entities.APIKeyEntity] {

	entry := logger.FromContext(cc.Context())

	entry.Info("Validating API key")

	cri := criteria.Criteria{
		Filters: *criteria.NewFilters(
			[]criteria.Filter{
				{
					Field:    "key",
					Value:    apiKey,
					Operator: criteria.OperatorEqual,
				},
			},
		),
	}

	apiKeyEntity, err := u.apiKeyRepository.Matching(cri)
	if err != nil {
		return utils.Response[entities.APIKeyEntity]{
			Success:    false,
			StatusCode: http.StatusInternalServerError,
			Error:      cc.NewError(cerrs.NewCustomError(http.StatusInternalServerError, err.Error(), "error_getting_api_key")),
		}
	}

	if len(apiKeyEntity) == 0 {
		return utils.Response[entities.APIKeyEntity]{
			Success:    false,
			StatusCode: http.StatusUnauthorized,
			Error:      cc.NewError(cerrs.NewCustomError(http.StatusUnauthorized, "API key not found", "api_key_not_found")),
		}
	}

	return utils.Response[entities.APIKeyEntity]{
		Success:    true,
		StatusCode: http.StatusOK,
		Data:       apiKeyEntity[0],
	}
}
