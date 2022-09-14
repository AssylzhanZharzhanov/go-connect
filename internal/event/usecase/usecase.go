package usecase

import "github.com/AssylzhanZharzhanov/connect/internal/event/domain"

type useCase struct {
	repository domain.EventRepository
}

// NewUseCase - creates a new "EventUseCase"
func NewUseCase(repository domain.EventRepository) domain.EventUseCase {
	return useCase{
		repository: repository,
	}
}

func (u useCase) CreateEvent(dto domain.CreateEventRequestDTO) error {
	//TODO implement me
	panic("implement me")
}

func (u useCase) GetEvent(eventID domain.EventID) (*domain.Event, error) {
	//TODO implement me
	panic("implement me")
}

func (u useCase) GetEvents() ([]*domain.Event, error) {
	//TODO implement me
	panic("implement me")
}

func (u useCase) UpdateEvent(eventID domain.EventID) error {
	//TODO implement me
	panic("implement me")
}

func (u useCase) DeleteEvent(eventID domain.EventID) error {
	//TODO implement me
	panic("implement me")
}
