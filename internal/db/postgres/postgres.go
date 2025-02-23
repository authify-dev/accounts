package postgres

import (
	"accounts/internal/core/domain"
	"fmt"
)

// --------------------------------
// INFRASTRUCTURE
// --------------------------------
// PostgresRepository
// --------------------------------

type PostgresRepository[E domain.IEntity] struct {
}

func (m *PostgresRepository[E]) View(data []E) {

	for _, e := range data {

		fmt.Println(string(domain.ToJSON[E](e)))
		fmt.Println("-------------------------------------------------")
	}

}
