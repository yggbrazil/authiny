package main

import (
	"authiny/internal/controllers/grpc"
	"authiny/internal/postgres_embedded"
	"authiny/internal/storages"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/yggbrazil/go-toolbox/env"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {
	DB_DATABASE := os.Getenv("AUTHINY_DB_DATABASE")
	DB_USER := os.Getenv("AUTHINY_DB_USER")
	DB_PASSWORD := os.Getenv("AUTHINY_DB_PASSWORD")
	AUTHINY_DB_FOLDER := os.Getenv("AUTHINY_DB_FOLDER")

	DB_PORT, err := strconv.Atoi(os.Getenv("AUTHINY_DB_PORT"))
	if err != nil {
		log.Fatal("Erro ao buscar a porta da variavel: AUTHINY_DB_PORT ", err)
	}

	postgres := postgres_embedded.NewPostgresEmbedded(postgres_embedded.PostgresEmbeddedConfig{
		Database: DB_DATABASE,
		User:     DB_USER,
		Password: DB_PASSWORD,
		Port:     DB_PORT,
		Folder:   AUTHINY_DB_FOLDER,
	})

	err = postgres.Start()
	if err != nil {
		log.Fatal(err)
	}

	defer postgres.Stop()

	storage, err := storages.NewPostgresStorage(storages.StoragePostgresConfig{
		User:     env.String("AUTHINY_DB_USER"),
		Port:     env.MustInt("AUTHINY_DB_PORT"),
		Password: env.String("AUTHINY_DB_PASSWORD"),
		Host:     env.String("AUTHINY_DB_HOST"),
		DataBase: env.String("AUTHINY_DB_DATABASE"),
		SSLMode:  env.String("AUTHINY_DB_SSLMODE"),
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
