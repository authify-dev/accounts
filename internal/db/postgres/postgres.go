package postgres

import (
	"accounts/internal/core/domain"
	"accounts/internal/core/domain/criteria"
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

func (r *PostgresRepository[E, M]) MatchingLow(cr criteria.Criteria, model *M) ([]E, error) {
	var roleModels []M

	// Se inicia la consulta sobre el modelo model.
	query := r.Connection.Model(model)

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
func (r *PostgresRepository[E, M]) Delete(uuid string) error {
	var model M
	// GORM permite pasar el valor de la clave primaria para borrar.
	if err := r.Connection.Delete(&model, uuid).Error; err != nil {
		return err
	}
	return nil
}

// Search busca y retorna la entidad asociada al UUID especificado.
func (r *PostgresRepository[E, M]) Search(uuid string) (E, error) {
	var model M
	// Se utiliza First para obtener el primer registro que coincida con el UUID.
	if err := r.Connection.First(&model, uuid).Error; err != nil {
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
	// Se filtra por el campo "id". Si el nombre de la clave primaria es distinto, cámbialo.
	if err := r.Connection.Model(&model).
		Where("id = ?", uuid).
		Updates(fields).Error; err != nil {
		return err
	}
	return nil
}
