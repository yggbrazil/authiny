package usecases

import (
	"authiny/internal/storages"
	"context"
)

type ApplicationService interface {
	Create(ctx context.Context, name string) (string, error)
}

type applicationService struct {
	storage storages.Storage
}

func NewApplicationService(storage storages.Storage) (ApplicationService, error) {
	return &applicationService{
		storage,
	}, nil
}

func (s *applicationService) Create(ctx context.Context, name string) (string, error) {
	return s.storage.CreateApplication(name)
}
