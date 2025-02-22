package memory

import "accounts/internal/core/domain"

// --------------------------------
// INFRASTRUCTURE
// --------------------------------
// MemoryRepository
// --------------------------------

type IMemoryRepository[E domain.IEntity] interface {
	View(data []E)
}
