package postgres

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RoleModel struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name        string    `gorm:"type:varchar(255);uniqueIndex;not null;"`
	Description string    `gorm:"type:varchar(255);not null;"`
}

func (RoleModel) TableName() string {
	return "roles"
}

// BeforeCreate se ejecuta antes de insertar un registro en la base de datos.
func (u *RoleModel) BeforeCreate(tx *gorm.DB) (err error) {
	// Si no se proporcionó un ID, lo generamos.
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return nil
}
