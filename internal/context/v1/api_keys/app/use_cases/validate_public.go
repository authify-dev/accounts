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
	"strings"
)

type ValidatePublicAPIKeyUseCase struct {
	apiKeyRepository repositories.APIKeyRepository
}

func NewValidatePublicAPIKeyUseCase(
	apiKeyRepository repositories.APIKeyRepository,
) *ValidatePublicAPIKeyUseCase {
	return &ValidatePublicAPIKeyUseCase{
		apiKeyRepository: apiKeyRepository,
	}
}

// Validate recibe la Secret API Key en formato: sk_<env>_<keyid>_<randomBase64Url>
func (u *ValidatePublicAPIKeyUseCase) Validate(cc *customctx.CustomContext, secretKey string) utils.Response[entities.APIKeyEntity] {
	entry := logger.FromContext(cc.Context())
	entry.Info("Validating Secret API key")

	parsed, perr := ParsePublicAPIKey(secretKey)
	if perr != nil {
		return utils.Response[entities.APIKeyEntity]{
			Success:    false,
			StatusCode: http.StatusUnauthorized,
			Error:      cc.NewError(cerrs.NewCustomError(http.StatusUnauthorized, "Invalid API key format", "api_key_invalid_format")),
		}
	}

	// Criterios: key_id + environment + is_active = true
	cri := criteria.Criteria{
		Filters: *criteria.NewFilters([]criteria.Filter{
			{
				Field:    "key_id",
				Value:    parsed.KeyID,
				Operator: criteria.OperatorEqual,
			},
			{
				Field:    "environment",
				Value:    parsed.Env,
				Operator: criteria.OperatorEqual,
			},
			{
				Field:    "is_active",
				Value:    true,
				Operator: criteria.OperatorEqual,
			},
		}),
		// si manejas paginación/orden, agrégalo aquí
	}

	// Esperamos UNA fila; tu repo puede devolver slice. Tomamos la primera.
	results, err := u.apiKeyRepository.Matching(cri)
	if err != nil {
		return utils.Response[entities.APIKeyEntity]{
			Success:    false,
			StatusCode: http.StatusInternalServerError,
			Error:      cc.NewError(cerrs.NewCustomError(http.StatusInternalServerError, err.Error(), "error_getting_api_key")),
		}
	}
	if len(results) == 0 {
		return utils.Response[entities.APIKeyEntity]{
			Success:    false,
			StatusCode: http.StatusUnauthorized,
			Error:      cc.NewError(cerrs.NewCustomError(http.StatusUnauthorized, "API key not found", "api_key_not_found")),
		}
	}

	ak := results[0]

	return utils.Response[entities.APIKeyEntity]{
		Success:    true,
		StatusCode: http.StatusOK,
		Data:       ak,
	}
}

type ParsedPublicAPIKey struct {
	Env   string // "live" | "test" (o lo que uses)
	KeyID string // hex (o base64url) que metes en la key
	Raw   string // tramo aleatorio (base64url)
}

// Espera formato: sk_<env>_<keyid>_<randomBase64Url>
func ParsePublicAPIKey(sk string) (ParsedPublicAPIKey, error) {
	parts := strings.SplitN(sk, "_", 4)
	if len(parts) != 4 || parts[0] != "pk" {
		return ParsedPublicAPIKey{}, cerrs.NewCustomError(http.StatusUnauthorized, "Invalid API key format", "api_key_invalid_format")
	}
	return ParsedPublicAPIKey{
		Env:   parts[1],
		KeyID: parts[2],
		Raw:   parts[3],
	}, nil
}
