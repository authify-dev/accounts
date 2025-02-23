package postgres

import "accounts/internal/core/domain"

// --------------------------------
// INFRASTRUCTURE
// --------------------------------
// PostgresRepository
// --------------------------------

type IPostgresRepository[E domain.IEntity] interface {
	View(data []E)
}
