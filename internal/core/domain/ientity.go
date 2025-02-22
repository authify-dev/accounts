package domain

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

// --------------------------------
// DOMAIN
// --------------------------------
// IEntity
//--------------------------------

// Definimos una interfaz que represente a una entidad.
type IEntity interface {
	GetID() uuid.UUID
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
