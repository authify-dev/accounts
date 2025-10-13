package usecases

import (
	"accounts/internal/common/controllers"
	common_repositories "accounts/internal/common/domain/repositories"
	"accounts/internal/common/logger"
	"accounts/internal/context/v1/api_keys/domain/commands"
	"accounts/internal/context/v1/api_keys/domain/entities"
	"accounts/internal/context/v1/api_keys/domain/repositories"
	"foundation/domain/customctx"
	"foundation/utils"
	"foundation/utils/cerrs"
	"net/http"
	"time"
)

type GenerateAPIKeysResponse struct {
	SecretKey      string    `json:"secret_key"`
	PublishableKey string    `json:"publishable_key"`
	CreatedAt      time.Time `json:"created_at"`
	OrganizationID string    `json:"organization_id"`
}

type GenerateAPIKeysUseCase struct {
	generatorAPIKey  *controllers.GeneratorAPIKey
	apiKeyRepository repositories.APIKeyRepository
	verifier         common_repositories.SecretVerifier
}

func NewGenerateAPIKeysUseCase(
	generatorAPIKey *controllers.GeneratorAPIKey,
	apiKeyRepository repositories.APIKeyRepository,
	verifier common_repositories.SecretVerifier,
) *GenerateAPIKeysUseCase {
	return &GenerateAPIKeysUseCase{
		generatorAPIKey:  generatorAPIKey,
		apiKeyRepository: apiKeyRepository,
		verifier:         verifier,
	}
}

func (u *GenerateAPIKeysUseCase) Generate(cc *customctx.CustomContext, command commands.GenerateAPIKeysCommand) utils.Response[GenerateAPIKeysResponse] {
	entry := logger.FromContext(cc.Context())

	entry.Info("Generating API key")

	km, err := u.generatorAPIKey.GeneratePair()
	if err != nil {
		return utils.Response[GenerateAPIKeysResponse]{
			Success:    false,
			StatusCode: http.StatusInternalServerError,
			Error: cc.NewError(
				cerrs.NewCustomError(
					http.StatusInternalServerError,
					"Error generating API key"+err.Error(),
					"generate_api_key_error",
				),
			),
		}
	}

	secretHash, err := u.verifier.Hash(km.SecretKey)
	if err != nil {
		return utils.Response[GenerateAPIKeysResponse]{
			Success:    false,
			StatusCode: http.StatusInternalServerError,
			Error:      cc.NewError(cerrs.NewCustomError(http.StatusInternalServerError, "Error hashing API key"+err.Error(), "hash_api_key_error")),
		}
	}

	ent := entities.APIKeyEntity{
		Name:           command.Name,
		Description:    command.Description,
		OrganizationID: command.OrganizationID,
		KeyID:          km.KeyID,
		Prefix:         km.Prefix,
		SecretKey:      km.SecretKey,      // solo para respuesta de creación
		PublishableKey: km.PublishableKey, // se puede exponer al cliente
		IsActive:       command.IsActive,
		Environment:    "live",     // o como lo definas
		Scopes:         []string{}, // si aplica
		SecretHash:     secretHash,
	}

	resultPublishable := u.apiKeyRepository.Save(ent)
	if resultPublishable.Err != nil {
		return utils.Response[GenerateAPIKeysResponse]{
			Success:    false,
			StatusCode: http.StatusInternalServerError,
			Error: cc.NewError(
				cerrs.NewCustomError(
					http.StatusInternalServerError,
					"Error saving API key"+resultPublishable.Err.Error(),
					"save_api_key_error",
				),
			),
		}
	}

	return utils.Response[GenerateAPIKeysResponse]{
		Data: GenerateAPIKeysResponse{
			SecretKey:      km.SecretKey,
			PublishableKey: resultPublishable.Data.PublishableKey,
			CreatedAt:      resultPublishable.Data.CreatedAt,
			OrganizationID: resultPublishable.Data.OrganizationID,
		},
		Success:    true,
		StatusCode: http.StatusCreated,
	}
}
