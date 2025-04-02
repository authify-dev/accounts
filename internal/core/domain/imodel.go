package domain

import (
	"accounts/internal/utils"
	"encoding/json"
)

// --------------------------------
// DOMAIN
// --------------------------------
// IModel
// --------------------------------
// Definimos una interfaz que represente a una entidad.
type IModel interface {
	GetID() string
}

func ModelToEntity[E IEntity, M IModel](model IModel) utils.Result[E] {
	var result map[string]interface{}

	// Convertir el struct a JSON (bytes).
	data, err := json.Marshal(model)
	if err != nil {
		return utils.Result[E]{Err: err}
	}

	// Convertir los bytes JSON a un mapa.
	err = json.Unmarshal(data, &result)
	if err != nil {
		return utils.Result[E]{Err: err}
	}

	entity, err := FromJSON[E](result)
	if err != nil {
		return utils.Result[E]{Err: err}
	}

	return utils.Result[E]{Data: entity}
}
