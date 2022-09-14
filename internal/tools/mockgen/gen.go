package mockgen

import (
	_ "github.com/golang/mock/mockgen/model" // for generating tests
)

//go:generate mockgen -package mocks -destination ../../../internal/mocks/mock_event_repository.go github.com/AssylzhanZharzhanov/connect/internal/event/domain EventRepository
//go:generate mockgen -package mocks -destination ../../../internal/mocks/mock_event_usecase.go github.com/AssylzhanZharzhanov/connect/internal/event/domain EventUseCase
