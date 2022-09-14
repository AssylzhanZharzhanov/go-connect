package usecase

import "github.com/AssylzhanZharzhanov/connect/internal/authorization/domain"

type useCase struct {
}

// NewUseCase - creates a new service with necessary dependencies.
func NewUseCase() domain.AuthorizationUseCase {
	return useCase{}
}
