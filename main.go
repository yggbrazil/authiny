package main

import (
	"authiny/internal/controllers/grpc"
	"authiny/internal/postgres_embedded"
	"authiny/internal/storages"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/yggbrazil/go-toolbox/env"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {
	DB_HOST := os.Getenv("AUTHINY_DB_HOST")
	DB_PORT := env.MustInt("AUTHINY_DB_PORT")
	DB_USER := os.Getenv("AUTHINY_DB_USER")
	DB_PASSWORD := os.Getenv("AUTHINY_DB_PASSWORD")
	DB_DATABASE := os.Getenv("AUTHINY_DB_DATABASE")
	DB_SSLMODE := os.Getenv("AUTHINY_DB_SSLMODE")

	AUTHINY_DB_FOLDER := os.Getenv("AUTHINY_DB_FOLDER")

	postgres := postgres_embedded.NewPostgresEmbedded(postgres_embedded.PostgresEmbeddedConfig{
		Database: DB_DATABASE,
		User:     DB_USER,
		Password: DB_PASSWORD,
		Port:     DB_PORT,
		Folder:   AUTHINY_DB_FOLDER,
	})

	err := postgres.Start()
	if err != nil {
		log.Fatal(err)
	}

	defer postgres.Stop()

	storage, err := storages.NewPostgresStorage(storages.StoragePostgresConfig{
		User:     DB_USER,
		Port:     DB_PORT,
		Password: DB_PASSWORD,
		Host:     DB_HOST,
		DataBase: DB_DATABASE,
		SSLMode:  DB_SSLMODE,
	})
	if err != nil {
		log.Fatal(err)
	}

	grpcServer, err := grpc.NewGrpcServer(storage)
	if err != nil {
		log.Fatal(err)
	}

	grpcServer.Run()
}
