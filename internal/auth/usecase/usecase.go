package usecase

import "github.com/AssylzhanZharzhanov/connect/internal/event/domain"

type useCase struct {
}

func NewUseCase() domain.EventUseCase {
	return useCase{}
}
