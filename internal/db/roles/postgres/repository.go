package postgres

import (
	"accounts/internal/api/v1/roles/domain/entities"
	base_entities "accounts/internal/core/domain"

	"gorm.io/gorm"
)

type RolePostgresRepository struct {
	Conection *gorm.DB
}

func NewRolePostgresRepository(db *gorm.DB) *RolePostgresRepository {
	return &RolePostgresRepository{
		Conection: db,
	}
}

func (u *RolePostgresRepository) Save(user entities.Role) error {
	u.Conection.Create(&RoleModel{
		Name:        user.Name,
		Description: user.Description,
	})
	return nil
}

func (u *RolePostgresRepository) List() ([]entities.Role, error) {
	var roles []RoleModel
	u.Conection.Find(&roles)
	var rolesEntities []entities.Role
	for _, role := range roles {
		rolesEntities = append(rolesEntities, entities.Role{
			Name:        role.Name,
			Description: role.Description,
			Entity: base_entities.Entity{
				ID:        role.ID,
				CreatedAt: role.CreatedAt,
			},
		})
	}
	return rolesEntities, nil
}
