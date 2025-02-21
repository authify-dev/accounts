package roles

import (
	"accounts/internal/api/v1/roles/domain/entities"
	"accounts/internal/common/criteria"
)

type RolesRepository interface {
	Save(user entities.Role) error
	List() ([]entities.Role, error)
	Filters(crit criteria.Criteria) ([]entities.Role, error)
}
