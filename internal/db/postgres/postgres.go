package postgres

import (
	"accounts/internal/core/domain"
	"fmt"

	"gorm.io/gorm"
)

// --------------------------------
// INFRASTRUCTURE
// --------------------------------
// PostgresRepository
// --------------------------------

type PostgresRepository[E domain.IEntity, M domain.IModel] struct {
	Connection *gorm.DB
}

func (m *PostgresRepository[E, M]) View(data []E) {

	for _, e := range data {

		fmt.Println(string(domain.ToJSON[E](e)))
		fmt.Println("-------------------------------------------------")
	}

}

func (r *PostgresRepository[E, M]) Save(role E) error {
	result := domain.EntityToModel[E, M](role)
	if result.Err != nil {
		return result.Err
	}

	roleModel := result.Data

	if err := r.Connection.Create(&roleModel).Error; err != nil {
		return err
	}

	return nil
}

func (r *PostgresRepository[E, M]) List() ([]E, error) {

	var roles []M

	if err := r.Connection.Find(&roles).Error; err != nil {
		return nil, err
	}

	var rolesEntities []E

	for _, role := range roles {

		result := domain.ModelToEntity[E, M](role)

		if result.Err != nil {
			return nil, result.Err
		}

		rolesEntities = append(rolesEntities, result.Data)
	}

	return rolesEntities, nil
}
