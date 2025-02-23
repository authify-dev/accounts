package postgres

import (
	"accounts/internal/api/v1/users/domain/entities"
	"accounts/internal/core/domain"
	"accounts/internal/core/domain/criteria"
	"accounts/internal/db/postgres"

	"gorm.io/gorm"
)

// --------------------------------
// INFRASTRUCTURE
// --------------------------------
// User Postgres Repository
// --------------------------------

type UserPostgresRepository struct {
	postgres.PostgresRepository[entities.User, UserModel]
	connection *gorm.DB
}

func NewUserPostgresRepository(connection *gorm.DB) *UserPostgresRepository {
	return &UserPostgresRepository{connection: connection}
}

func (r *UserPostgresRepository) Save(entity entities.User) error {
	result := domain.EntityToModel[entities.User, UserModel](entity)
	if result.Err != nil {
		return result.Err
	}

	model := result.Data

	if err := r.connection.Create(&model).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserPostgresRepository) List() ([]entities.User, error) {

	var records []UserModel

	if err := r.connection.Find(&records).Error; err != nil {
		return nil, err
	}

	var recordsEntities []entities.User

	for _, record := range records {

		result := domain.ModelToEntity[entities.User, UserModel](record)

		if result.Err != nil {
			return nil, result.Err
		}

		recordsEntities = append(recordsEntities, result.Data)
	}

	return recordsEntities, nil
}

func (r *UserPostgresRepository) Matching(criteria criteria.Criteria) ([]entities.User, error) {

	return r.List()
}
