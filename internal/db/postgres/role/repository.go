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
	postgres.PostgresRepository[entities.Role]
	connection *gorm.DB
}

func NewRolePostgresRepository(connection *gorm.DB) *RolePostgresRepository {
	return &RolePostgresRepository{connection: connection}
}

func (r *RolePostgresRepository) Save(role entities.Role) error {
	result := domain.EntityToModel[entities.Role, RoleModel](role)
	if result.Err != nil {
		return result.Err
	}

	roleModel := result.Data

	if err := r.connection.Create(&roleModel).Error; err != nil {
		return err
	}

	return nil
}

func (r *RolePostgresRepository) List() ([]entities.Role, error) {

	var roles []RoleModel

	if err := r.connection.Find(&roles).Error; err != nil {
		return nil, err
	}

	var rolesEntities []entities.Role

	for _, role := range roles {

		result := domain.ModelToEntity[entities.Role, RoleModel](role)

		if result.Err != nil {
			return nil, result.Err
		}

		rolesEntities = append(rolesEntities, result.Data)
	}

	return rolesEntities, nil
}

func (r *RolePostgresRepository) Matching(cr criteria.Criteria) ([]entities.Role, error) {
	var roleModels []RoleModel

	// Se inicia la consulta sobre el modelo RoleModel.
	query := r.connection.Model(&RoleModel{})

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

	// Convertir cada RoleModel obtenido a la entidad Role.
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
