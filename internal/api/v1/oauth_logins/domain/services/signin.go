package services

import (
	"accounts/internal/api/v1/emails/domain/entities"
	"accounts/internal/utils"
	"context"
)

func (s *OAuthService) SignInGoogle(ctx context.Context, code string) utils.Responses[entities.SignInResponse] {

	// Obtener el token
	// Obtener la data del usuario

	// Verificamos que el email del user no este en uso
	// Creanos el User
	// Creamos el OAuthUser
	// Creamos el login
	// Enviamos el email de binevenida

	return utils.Responses[entities.SignInResponse]{
		StatusCode: 200,
	}
}
