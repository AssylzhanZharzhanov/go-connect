package repository

import (
	"github.com/AssylzhanZharzhanov/connect/models/user/dto"
	"github.com/AssylzhanZharzhanov/connect/models/user/entity"
	"github.com/jmoiron/sqlx"
)

type AuthRepository struct {
	db *sqlx.DB
}

func (a AuthRepository) CreateUser(body dto.UserCreateRequestBody) (entity.Users, error) {
	panic("implement me")
}

func (a AuthRepository) GetUser(body dto.UserAuthenticateRequestBody) (entity.Users, error) {
	panic("implement me")
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}
