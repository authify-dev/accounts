package memory

import (
	"accounts/internal/core/domain"
	"fmt"
)

// --------------------------------
// INFRASTRUCTURE
// --------------------------------
// MemoryRepository
// --------------------------------

type MemoryRepository[E domain.IEntity] struct {
}

func (m *MemoryRepository[E]) View(data []E) {

	for _, e := range data {

		fmt.Println(string(domain.ToJSON[E](e)))
		fmt.Println("-------------------------------------------------")
	}

}
