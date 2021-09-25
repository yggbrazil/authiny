package postgres_embedded

import (
	"time"

	embeddedpostgres "github.com/fergusstrange/embedded-postgres"
)

type PostgresEmbeddedConfig struct {
	Database string
	User     string
	Password string
	Port     int
	Folder   string
}

type PostgresEmbedded struct {
	postgres *embeddedpostgres.EmbeddedPostgres
}

func NewPostgresEmbedded(p PostgresEmbeddedConfig) *PostgresEmbedded {
	postgres := embeddedpostgres.NewDatabase(embeddedpostgres.DefaultConfig().
		Version(embeddedpostgres.V13).
		StartTimeout(45 * time.Second).
		RuntimePath(p.Folder).
		Username(p.User).
		Password(p.Password).
		Database(p.Database).
		Port(uint32(p.Port)),
	)

	return &PostgresEmbedded{
		postgres,
	}
}

func (p *PostgresEmbedded) Start() error {
	return p.postgres.Start()
}

func (p *PostgresEmbedded) Stop() error {
	return p.postgres.Stop()
}
