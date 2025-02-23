package postgres

import (
	"accounts/internal/api/v1/roles/domain/entities"
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

type RolePostgresRepository struct {
	postgres.PostgresRepository[entities.Role, RoleModel]
}

func NewRolePostgresRepository(connection *gorm.DB) *RolePostgresRepository {
	return &RolePostgresRepository{
		PostgresRepository: postgres.PostgresRepository[entities.Role, RoleModel]{
			Connection: connection,
		},
	}
}

func (r *RolePostgresRepository) Matching(cr criteria.Criteria) ([]entities.Role, error) {
	var roleModels []RoleModel

	// Se inicia la consulta sobre el modelo model.
	query := r.Connection.Model(&RoleModel{})

	// Se recorren los filtros para agregarlos a la consulta.
	for _, f := range cr.Filters.Get() {
		// Construir la condición de la consulta, por ejemplo: "name = ?"
		condition := fmt.Sprintf("%s %s ?", f.Field, f.Operator)
		query = query.Where(condition, f.Value)
	}

	// Ejecuta la consulta y almacena el resultado en roleModels.
	err := query.Find(&roleModels).Error
	if err != nil {
		return nil, err
	}

	// Convertir cada model obtenido a la entidad Role.
	var roles []entities.Role
	for _, rm := range roleModels {
		result := domain.ModelToEntity[entities.Role, RoleModel](rm)
		if result.Err != nil {
			return nil, result.Err
		}
		roles = append(roles, result.Data)
	}

	return roles, nil
}
