package service

import (
	"github.com/AssylzhanZharzhanov/connect/models/user/dto"
	"github.com/AssylzhanZharzhanov/connect/pkg/repository"
)

type AuthService struct {
	repos repository.Auth
}

func (s *AuthService) CreateUser(body dto.UserCreateRequestBody) (dto.UserResponseBody, error) {
	panic("implement me")
}

func (s *AuthService) AuthenticateUser(body dto.UserAuthenticateRequestBody) (dto.UserAuthenticateResponseBody, error) {
	panic("implement me")
}

func NewAuthService(repository *repository.Repository) *AuthService {
	return &AuthService{
		repos: repository.Auth,
	}
}

