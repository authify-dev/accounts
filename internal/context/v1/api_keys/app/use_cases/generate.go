package usecases

import (
	"accounts/internal/common/controllers"
	"accounts/internal/common/logger"
	"accounts/internal/context/v1/api_keys/domain/commands"
	"accounts/internal/context/v1/api_keys/domain/entities"
	"accounts/internal/context/v1/api_keys/domain/repositories"
	"foundation/domain/customctx"
	"foundation/utils"
	"foundation/utils/cerrs"
	"net/http"
)

type GenerateAPIKeysUseCase struct {
	generatorAPIKey  *controllers.GeneratorAPIKey
	apiKeyRepository repositories.APIKeyRepository
}

func NewGenerateAPIKeysUseCase(
	generatorAPIKey *controllers.GeneratorAPIKey,
	apiKeyRepository repositories.APIKeyRepository,
) *GenerateAPIKeysUseCase {
	return &GenerateAPIKeysUseCase{
		generatorAPIKey:  generatorAPIKey,
		apiKeyRepository: apiKeyRepository,
	}
}

func (u *GenerateAPIKeysUseCase) Generate(cc *customctx.CustomContext, command commands.GenerateAPIKeysCommand) utils.Response[entities.APIKeyEntity] {
	entry := logger.FromContext(cc.Context())

	entry.Info("Generating API key")

	apiKey, err := u.generatorAPIKey.GenerateAPIKey()
	if err != nil {
		return utils.Response[entities.APIKeyEntity]{
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

	apiKeyEntity := command.ToEntity()
	apiKeyEntity.Key = apiKey

	result := u.apiKeyRepository.Save(apiKeyEntity)
	if result.Err != nil {
		return utils.Response[entities.APIKeyEntity]{
			Success:    false,
			StatusCode: http.StatusInternalServerError,
			Error: cc.NewError(
				cerrs.NewCustomError(
					http.StatusInternalServerError,
					"Error saving API key"+result.Err.Error(),
					"save_api_key_error",
				),
			),
		}
	}

	return utils.Response[entities.APIKeyEntity]{
		Data:       result.Data,
		Success:    true,
		StatusCode: http.StatusCreated,
	}
}
