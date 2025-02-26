package services

import (
	"accounts/internal/api/v1/emails/domain/entities"
	"accounts/internal/api/v1/emails/domain/steps"
	users "accounts/internal/api/v1/users/domain/entities"
	"accounts/internal/common/controllers/saga"
	"accounts/internal/core/domain"
	"accounts/internal/utils"
	"context"

	"github.com/google/uuid"
)

func (s *EmailsService) SignUp(
	ctx context.Context,
	entity entities.SignUp,
) utils.Responses[entities.SignUpResponse] {

	id := uuid.New()

	user := users.User{
		Entity: domain.Entity{
			ID: id,
		},
		UserName: entity.UserName,
		Role:     entity.Role,
	}

	controller := saga.SAGA_Controller{
		Steps: []saga.SAGA_Step[any]{
			steps.NewCreateUserStep(
				s.user_repository,
				s.role_repository,
				user,
			),
		},
	}

	results := controller.Executed(ctx)

	res := entities.SignUpResponse{
		JWT:          "jwt",
		RefreshToken: "refresh_token",
	}

	response := utils.Responses[entities.SignUpResponse]{Body: res, StatusCode: 201}

	for _, result := range results {
		if result.Err != nil {
			response.Errors = append(response.Errors, result.Err.Error())
		}
	}

	return response
}
