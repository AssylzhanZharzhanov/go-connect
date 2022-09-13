package repository

import (
	"github.com/AssylzhanZharzhanov/connect/internal/event/domain"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) domain.EventRepository {
	return repository{
		db: db,
	}
}
