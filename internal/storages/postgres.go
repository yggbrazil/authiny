package storages

import (
	"authiny/internal/entitys"

	"github.com/jmoiron/sqlx"
	"github.com/yggbrazil/go-toolbox/database"
	"github.com/yggbrazil/go-toolbox/env"
)

type StoragePostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DataBase string
	SSLMode  string
}

type storage struct {
	db *sqlx.DB
}

func NewPostgresStorage(c StoragePostgresConfig) (Storage, error) {
	db, err := database.ConnectByConfig(database.Config{
		ID:       "postgres",
		Type:     "postgres",
		User:     env.String("AUTHINY_DB_USER"),
		Port:     env.MustInt("AUTHINY_DB_PORT"),
		Password: env.String("AUTHINY_DB_PASSWORD"),
		Host:     env.String("AUTHINY_DB_HOST"),
		DataBase: env.String("AUTHINY_DB_DATABASE"),
		SSLMode:  env.String("AUTHINY_DB_SSLMODE"),
	})

	if err != nil {
		return &storage{}, err
	}

	return &storage{
		db,
	}, nil
}

func (s *storage) CreateApplication(applicationName string) (string, error) {
	var a entitys.Application

	err := s.db.Get(&a, `INSERT INTO applications (name) VALUES ($1) RETURNING id`, applicationName)

	if err != nil {
		return "", err
	}

	return a.ID, nil
}

func (s *storage) GetApplication(applicationID string) (entitys.Application, error) {
	var a entitys.Application

	err := s.db.Get(&a, `SELECT id, name FROM applications WHERE id = $1`, applicationID)
	if err != nil {
		return entitys.Application{}, err
	}

	return a, nil
}

func (s *storage) ListApplications() ([]entitys.Application, error) {
	var applications []entitys.Application

	err := s.db.Select(&applications, `SELECT id, name FROM applications`)
	if err != nil {
		return []entitys.Application{}, err
	}

	return applications, nil
}

func (s *storage) DeleteApplication(applicationID string) error {
	_, err := s.db.NamedExec(`DELETE FROM applications WHERE id = :id`, map[string]interface{}{
		"id": applicationID,
	})

	if err != nil {
		return err
	}

	return nil
}
