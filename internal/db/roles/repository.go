package roles

import "accounts/internal/api/v1/roles/domain/entities"

type RolesRepository interface {
	Save(user entities.Role) error
	List() ([]entities.Role, error)
}
