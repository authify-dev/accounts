package users

import "accounts/internal/api/v1/users/domain/entities"

type UserRepository interface {
	Save(user entities.User) error
}
