package domain

// EventUseCase - represents use case layer
type EventUseCase interface {
	CreateEvent(dto CreateEventRequestDTO) error
	GetEvent(eventID EventID) (*Event, error)
	GetEvents() ([]*Event, error)
	UpdateEvent(eventID EventID) error
	DeleteEvent(eventID EventID) error
}
