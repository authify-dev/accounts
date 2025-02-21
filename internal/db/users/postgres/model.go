package postgres

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	UserName  string    `gorm:"type:varchar(255);uniqueIndex;not null;"`
	Name      string    `gorm:"type:varchar(255);not null;"`
	Birthdate string    `gorm:"type:date;not null;"`
	RoleID    string    `gorm:"type:uuid;not null;"`
}

func (UserModel) TableName() string {
	return "users"
}

// BeforeCreate se ejecuta antes de insertar un registro en la base de datos.
func (u *UserModel) BeforeCreate(tx *gorm.DB) (err error) {
	// Si no se proporcionó un ID, lo generamos.
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	// Si el username viene en blanco, lo generamos como "User" + primeros 8 dígitos del UUID.
	if u.UserName == "" {
		u.UserName = "User" + u.ID.String()[0:8]
	}
	return nil
}
