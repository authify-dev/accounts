package postgres

import (
	"accounts/internal/api/v1/codes/domain/entities"
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

type CodePostgresRepository struct {
	postgres.PostgresRepository[entities.Code, CodeModel]
	connection *gorm.DB
}

func NewCodePostgresRepository(connection *gorm.DB) *CodePostgresRepository {
	return &CodePostgresRepository{connection: connection}
}

func (r *CodePostgresRepository) Matching(cr criteria.Criteria) ([]entities.Code, error) {
	var records []CodeModel

	// Se inicia la consulta sobre el modelo CodeModel.
	query := r.connection.Model(&CodeModel{})

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

	// Convertir cada CodeModel obtenido a la entidad Role.
	var recordsEntities []entities.Code
	for _, rm := range records {
		result := domain.ModelToEntity[entities.Code, CodeModel](rm)
		if result.Err != nil {
			return nil, result.Err
		}
		recordsEntities = append(recordsEntities, result.Data)
	}

	return recordsEntities, nil
}
