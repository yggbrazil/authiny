package storages

import "authiny/internal/entitys"

type Storage interface {
	CreateApplication(applicationName string) (string, error)

	GetApplication(applicationID string) (entitys.Application, error)

	ListApplications() ([]entitys.Application, error)

	DeleteApplication(applicationID string) error
}
