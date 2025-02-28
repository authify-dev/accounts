package steps

import (
	codes_entities "accounts/internal/api/v1/codes/domain/entities"
	"accounts/internal/core/domain"
	"math/rand"
	"time"

	codes "accounts/internal/api/v1/codes/domain/repositories"
	"accounts/internal/common/logger"
	"accounts/internal/utils"
	"context"

	"github.com/google/uuid"
)

type CreateCodeStep struct {
	code_id    uuid.UUID
	user_id    uuid.UUID
	codes_repo codes.CodeRepository
}

func NewCreateCodeStep(
	codes_repo codes.CodeRepository,
	user_id uuid.UUID,
) *CreateCodeStep {
	return &CreateCodeStep{
		codes_repo: codes_repo,
		user_id:    user_id,
	}
}

func (s *CreateCodeStep) Call(ctx context.Context, payload utils.Result[any], allPayloads map[string]utils.Result[any]) utils.Result[any] {

	entry := logger.FromContext(ctx)

	id := uuid.New()

	code_entity := codes_entities.Code{
		UserID: s.user_id.String(),
		Entity: domain.Entity{
			ID: id,
		},
		Code: generateCode(6),
	}

	err := s.codes_repo.Save(code_entity)
	if err != nil {
		entry.Error("error saving the code")
		return utils.Result[any]{Err: err}
	}

	s.code_id = id

	return utils.Result[any]{
		Data: code_entity,
	}
}

func (s *CreateCodeStep) Rollback(ctx context.Context) error {
	entry := logger.FromContext(ctx)

	// Implementación de la lógica de negocio
	if s.code_id == uuid.Nil {
		entry.Error("code_id is empty")
		return nil
	}

	if err := s.codes_repo.Delete(s.code_id); err != nil {
		entry.Error("error deleting user")
		return err
	}

	entry.Info("user deleted")
	return nil
}

func (s *CreateCodeStep) Produce() string {
	return "entities.Code"
}

func generateCode(longitud int) string {
	const numeros = "0123456789"
	resultado := make([]byte, longitud)

	// Se crea un generador local de números aleatorios con semilla.
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < longitud; i++ {
		resultado[i] = numeros[r.Intn(len(numeros))]
	}
	return string(resultado)
}
