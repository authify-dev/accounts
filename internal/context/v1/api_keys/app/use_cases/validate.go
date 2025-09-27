package usecases

import (
	common_repositories "accounts/internal/common/domain/repositories"
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

type ValidateAPIKeyUseCase struct {
	apiKeyRepository repositories.APIKeyRepository
	verifier         common_repositories.SecretVerifier
}

func NewValidateAPIKeyUseCase(
	apiKeyRepository repositories.APIKeyRepository,
	verifier common_repositories.SecretVerifier,
) *ValidateAPIKeyUseCase {
	return &ValidateAPIKeyUseCase{
		apiKeyRepository: apiKeyRepository,
		verifier:         verifier,
	}
}

// Validate recibe la Secret API Key en formato: sk_<env>_<keyid>_<randomBase64Url>
func (u *ValidateAPIKeyUseCase) Validate(cc *customctx.CustomContext, secretKey string) utils.Response[entities.APIKeyEntity] {
	entry := logger.FromContext(cc.Context())
	entry.Info("Validating Secret API key")

	parsed, perr := ParseSecretKey(secretKey)
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

	// IMPORTANTE:
	// La entidad que retornas desde el repositorio debe incluir el SecretHash.
	// Asegúrate que entities.APIKeyEntity tenga:
	//   SecretHash string `json:"-"` // no exponer en JSON
	// o bien cambia el repo para retornar una proyección que lo contenga.
	if ak.SecretHash == "" {
		return utils.Response[entities.APIKeyEntity]{
			Success:    false,
			StatusCode: http.StatusInternalServerError,
			Error:      cc.NewError(cerrs.NewCustomError(http.StatusInternalServerError, "Secret hash not available", "secret_hash_missing")),
		}
	}

	ok := u.verifier.Verify(ak.SecretHash, secretKey)

	// Verificar hash de la secret completa (tal cual llegó, con prefijo sk_... )
	if !ok {
		return utils.Response[entities.APIKeyEntity]{
			Success:    false,
			StatusCode: http.StatusUnauthorized,
			Error:      cc.NewError(cerrs.NewCustomError(http.StatusUnauthorized, "Invalid API key", "api_key_invalid")),
		}
	}

	// (Opcional) Si quieres, aquí podrías actualizar last_used_at en otra unidad de trabajo/handler.

	return utils.Response[entities.APIKeyEntity]{
		Success:    true,
		StatusCode: http.StatusOK,
		Data:       ak,
	}
}

// ParsedSecret representa el desglose de la Secret API Key.
type ParsedSecret struct {
	Env   string // "live" | "test" (o lo que uses)
	KeyID string // hex (o base64url) que metes en la key
	Raw   string // tramo aleatorio (base64url)
}

// Espera formato: sk_<env>_<keyid>_<randomBase64Url>
func ParseSecretKey(sk string) (ParsedSecret, error) {
	parts := strings.SplitN(sk, "_", 4)
	if len(parts) != 4 || parts[0] != "sk" {
		return ParsedSecret{}, cerrs.NewCustomError(http.StatusUnauthorized, "Invalid API key format", "api_key_invalid_format")
	}
	return ParsedSecret{
		Env:   parts[1],
		KeyID: parts[2],
		Raw:   parts[3],
	}, nil
}
