package emails

import (
	"accounts/internal/api/v1/refresh_tokens/domain/entities"
	"accounts/internal/core/domain"
	"accounts/internal/core/domain/criteria"
	"accounts/internal/db/postgres"
	"fmt"

	"gorm.io/gorm"
)

// --------------------------------
// INFRASTRUCTURE
// --------------------------------
// Role Postgres Repository
// --------------------------------

type LoginMethodPostgresRepository struct {
	postgres.PostgresRepository[entities.RefreshToken, RefreshTokenModel]
	connection *gorm.DB
}

func NewLoginMethodPostgresRepository(connection *gorm.DB) *LoginMethodPostgresRepository {
	return &LoginMethodPostgresRepository{connection: connection}
}

func (r *LoginMethodPostgresRepository) Save(entity entities.RefreshToken) error {
	result := domain.EntityToModel[entities.RefreshToken, RefreshTokenModel](entity)
	if result.Err != nil {
		return result.Err
	}

	entityModel := result.Data

	if err := r.connection.Create(&entityModel).Error; err != nil {
		return err
	}

	return nil
}

func (r *LoginMethodPostgresRepository) List() ([]entities.RefreshToken, error) {

	var records []RefreshTokenModel

	if err := r.connection.Find(&records).Error; err != nil {
		return nil, err
	}

	var recordsEntities []entities.RefreshToken

	for _, record := range records {

		result := domain.ModelToEntity[entities.RefreshToken, RefreshTokenModel](record)

		if result.Err != nil {
			return nil, result.Err
		}

		recordsEntities = append(recordsEntities, result.Data)
	}

	return recordsEntities, nil
}

func (r *LoginMethodPostgresRepository) Matching(cr criteria.Criteria) ([]entities.RefreshToken, error) {
	var records []RefreshTokenModel

	// Se inicia la consulta sobre el RefreshTokenModelo RefreshTokenModel.
	query := r.connection.Model(&RefreshTokenModel{})

	// Se recorren los filtros para agregarlos a la consulta.
	for _, f := range cr.Filters.Get() {
		// Construir la condición de la consulta, por ejemplo: "name = ?"
		condition := fmt.Sprintf("%s %s ?", f.Field, f.Operator)
		query = query.Where(condition, f.Value)
	}

	// Ejecuta la consulta y almacena el resultado en records.
	err := query.Find(&records).Error
	if err != nil {
		return nil, err
	}

	// Convertir cada RefreshTokenModel obtenido a la entidad Role.
	var recordsEntities []entities.RefreshToken
	for _, rm := range records {
		result := domain.ModelToEntity[entities.RefreshToken, RefreshTokenModel](rm)
		if result.Err != nil {
			return nil, result.Err
		}
		recordsEntities = append(recordsEntities, result.Data)
	}

	return recordsEntities, nil
}
