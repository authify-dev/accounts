package domain

import (
	"encoding/json"
	"foundation/utils"
	"foundation/utils/cerrs"
	"net/http"

	"github.com/google/uuid"
)

// --------------------------------
// DOMAIN
// --------------------------------
// IModel
// --------------------------------
// Definimos una interfaz que represente a una entidad.
type IModel interface {
	GetID() uuid.UUID
}

func ModelToEntity[E IEntity, M IModel](model IModel) utils.Result[E] {
	var result map[string]interface{}

	// Convertir el struct a JSON (bytes).
	data, err := json.Marshal(model)
	if err != nil {
		return utils.Result[E]{Err: &cerrs.CustomError{
			Code:    http.StatusInternalServerError,
			Message: "Error in convert model to entity: " + err.Error(),
			Scope:   "model_to_entity",
		}}
	}

	// Convertir los bytes JSON a un mapa.
	err = json.Unmarshal(data, &result)
	if err != nil {
		return utils.Result[E]{Err: &cerrs.CustomError{
			Code:    http.StatusInternalServerError,
			Message: "Error in convert model to entity: " + err.Error(),
			Scope:   "model_to_entity",
		}}
	}

	entity, err := FromJSON[E](result)
	if err != nil {
		return utils.Result[E]{Err: &cerrs.CustomError{
			Code:    http.StatusInternalServerError,
			Message: "Error in convert model to entity: " + err.Error(),
			Scope:   "model_to_entity",
		}}
	}

	return utils.Result[E]{Data: entity, Err: nil}
}
