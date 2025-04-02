package postgres

import (
	"accounts/internal/core/domain"
	"accounts/internal/core/domain/criteria"
	"accounts/internal/utils"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
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

func (r *PostgresRepository[E, M]) Save(role E) utils.Either[string] {
	result := domain.EntityToModel[E, M](role)
	if result.Err != nil {
		return utils.Either[string]{Err: result.Err}
	}

	roleModel := result.Data

	if err := r.Connection.Create(&roleModel).Error; err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			err = fmt.Errorf("duplicate key error: a record with the same unique key (%s) already exists", pgErr.Detail)
			return utils.Either[string]{Err: err}
		}
		return utils.Either[string]{Err: err}
	}

	return utils.Either[string]{Data: roleModel.GetID()}
}

func (r *PostgresRepository[E, M]) SearchAll() ([]E, error) {

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

func (r *PostgresRepository[E, M]) MatchingLow(cr criteria.Criteria, model *M) ([]E, error) {
	var roleModels []M

	// Se inicia la consulta sobre el modelo model.
	query := r.Connection.Debug().Model(model)

	// Se recorren los filtros para agregarlos a la consulta.
	for _, f := range cr.Filters.Get() {
		// Construir la condición de la consulta, por ejemplo: "name = ?"
		condition := fmt.Sprintf("%s %s ?", f.Field, f.Operator)
		query = query.Where(condition, f.Value)
	}

	// Ejecuta la consulta y almacena el resultado en roleModels.
	err := query.Find(&roleModels).Error
	if err != nil {
		return nil, err
	}

	sql := query.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx
	})
	fmt.Println(sql)

	// Convertir cada model obtenido a la entidad Role.
	var roles []E
	for _, rm := range roleModels {
		result := domain.ModelToEntity[E, M](rm)
		if result.Err != nil {
			return nil, result.Err
		}
		roles = append(roles, result.Data)
	}

	return roles, nil
}

// Delete elimina el registro que tenga el UUID especificado.
func (r *PostgresRepository[E, M]) Delete(id string) error {
	var model M

	// Borra el registro asegurándote de pasar el ID como string en la consulta
	if err := r.Connection.Delete(&model, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}

// Search busca y retorna la entidad asociada al UUID especificado.
func (r *PostgresRepository[E, M]) Search(id string) (E, error) {
	var model M
	// Se utiliza First para obtener el primer registro que coincida con el UUID.
	if err := r.Connection.First(&model, "id = ?", id).Error; err != nil {
		var empty E
		return empty, err
	}

	// Convertir el modelo obtenido a entidad.
	result := domain.ModelToEntity[E, M](model)
	if result.Err != nil {
		var empty E
		return empty, result.Err
	}

	return result.Data, nil
}

// UpdateByFields actualiza los campos indicados en el mapa para el registro con el UUID especificado.
func (r *PostgresRepository[E, M]) UpdateByFields(uuid string, fields map[string]interface{}) error {
	var model M

	fields["updated_at"] = time.Now().UTC()

	// Se filtra por el campo "id". Si el nombre de la clave primaria es distinto, cámbialo.
	if err := r.Connection.Model(&model).
		Where("id = ?", uuid).
		Updates(fields).Error; err != nil {
		return err
	}
	return nil
}
