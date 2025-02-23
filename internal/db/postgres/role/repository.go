package postgres

import (
	"accounts/internal/api/v1/roles/domain/entities"
	"accounts/internal/core/domain"
	"accounts/internal/core/domain/criteria"
	"accounts/internal/db/postgres"
	"time"

	"github.com/google/uuid"
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

	roleModel.ID = uuid.New()
	roleModel.CreatedAt = time.Now()
	roleModel.UpdatedAt = time.Now()
	roleModel.IsRemoved = false

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

func (r *RolePostgresRepository) Matching(criteria criteria.Criteria) ([]entities.Role, error) {

	return r.List()
}
