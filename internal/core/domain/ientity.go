package domain

import (
	"accounts/internal/utils"
	"encoding/json"
	"fmt"
)

// --------------------------------
// DOMAIN
// --------------------------------
// IEntity
//--------------------------------

// Definimos una interfaz que represente a una entidad.
type IEntity interface {
	GetID() string
}

func ToJSON[E IEntity](entity E) []byte {
	jsonData, err := json.MarshalIndent(entity, "", "  ")
	if err != nil {
		fmt.Println("Error al convertir a JSON:", err)
		return nil
	}

	return jsonData
}

// Función genérica que opera sobre tipos que cumplen con IEntity.
func FromJSON[E IEntity](m map[string]interface{}) (E, error) {
	var entity E

	// Convertir el mapa a bytes JSON.
	bytes, err := json.Marshal(m)
	if err != nil {
		return entity, err
	}

	// Deserializar los bytes JSON en la entidad.
	err = json.Unmarshal(bytes, &entity)
	return entity, err
}

func EntityToModel[E IEntity, M IModel](entity IEntity) utils.Result[M] {
	var result map[string]interface{}

	// Convertir la entidad a JSON (bytes).
	data, err := json.Marshal(entity)
	if err != nil {
		return utils.Result[M]{Err: err}
	}

	// Convertir los bytes JSON a un mapa.
	err = json.Unmarshal(data, &result)
	if err != nil {
		return utils.Result[M]{Err: err}
	}

	// Convertir el mapa a modelo.
	model, err := FromJSON[M](result)
	if err != nil {
		return utils.Result[M]{Err: err}
	}

	return utils.Result[M]{Data: model}
}
