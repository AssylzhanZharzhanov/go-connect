package repository

import (
	"github.com/AssylzhanZharzhanov/connect/models/user/dto"
	"github.com/AssylzhanZharzhanov/connect/models/user/entity"
	"github.com/jmoiron/sqlx"
)

type Auth interface {
	CreateUser(dto.UserCreateRequestBody) (entity.Users, error)
	GetUser(body dto.UserAuthenticateRequestBody) (entity.Users, error)
}

type Repository struct {
	Auth
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Auth: NewAuthRepository(db),
	}
}
