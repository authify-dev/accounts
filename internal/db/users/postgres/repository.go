package postgres

import (
	"accounts/internal/api/v1/users/domain/entities"

	"gorm.io/gorm"
)

type UserPostgresRepository struct {
	Conection *gorm.DB
}

func NewUserPostgresRepository(db *gorm.DB) *UserPostgresRepository {
	return &UserPostgresRepository{
		Conection: db,
	}
}

func (u *UserPostgresRepository) Save(user entities.User) error {
	u.Conection.Create(&UserModel{
		UserName: user.UserName,
	})
	return nil
}
