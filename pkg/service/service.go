package service

import (
	"github.com/AssylzhanZharzhanov/connect/models/user/dto"
	"github.com/AssylzhanZharzhanov/connect/pkg/repository"
)

type Auth interface {
	CreateUser(dto.UserCreateRequestBody) (dto.UserResponseBody, error)
	AuthenticateUser(body dto.UserAuthenticateRequestBody) (dto.UserAuthenticateResponseBody, error)
}

type Service struct {
	Auth
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Auth: NewAuthService(repository),
	}
}
