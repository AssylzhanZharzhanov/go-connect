package usecase

import (
	"context"
	"github.com/AssylzhanZharzhanov/connect/internal/event/domain"
)

type useCase struct {
	repository domain.EventRepository
}

// NewUseCase - creates a new "EventUseCase"
func NewUseCase(repository domain.EventRepository) domain.EventUseCase {
	return useCase{
		repository: repository,
	}
}

func (u useCase) CreateEvent(ctx context.Context, dto domain.CreateEventRequestDTO) error {
	//TODO implement me
	return nil
}

func (u useCase) GetEvent(ctx context.Context, eventID domain.EventID) (*domain.Event, error) {
	//TODO implement me
	panic("implement me")
}

func (u useCase) GetEvents(ctx context.Context) ([]*domain.Event, error) {
	//TODO implement me
	panic("implement me")
}

func (u useCase) UpdateEvent(ctx context.Context, eventID domain.EventID, dto domain.UpdateEventRequestDTO) error {
	//TODO implement me
	panic("implement me")
}

func (u useCase) DeleteEvent(ctx context.Context, eventID domain.EventID) error {
	//TODO implement me
	panic("implement me")
}
