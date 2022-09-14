package domain

import "context"

// EventUseCase - represents use case layer
type EventUseCase interface {
	CreateEvent(ctx context.Context, dto CreateEventRequestDTO) error
	GetEvent(ctx context.Context, eventID EventID) (*Event, error)
	GetEvents(ctx context.Context) ([]*Event, error)
	UpdateEvent(ctx context.Context, eventID EventID, dto UpdateEventRequestDTO) error
	DeleteEvent(ctx context.Context, eventID EventID) error
}
