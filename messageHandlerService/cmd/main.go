package main

import (
	"message-handler/internal/kafka"
	"message-handler/internal/repository"
	"message-handler/internal/service"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.DebugLevel)
}

func main() {
	if err := godotenv.Load("/app/.env"); err != nil {
		log.Fatalf("Error loading env variables.")
	}

	db, err := repository.ConnectDB(repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMode"),
	})

	if err != nil {
		log.Fatalf("Failed to init DB: %s", err.Error())
	}

	log.Info("Successfully connected to database")

	repo := repository.NewRepository(db)
	services := service.NewService(repo)

	log.Info("Starting listening")

	kafka.Listen(services)
}
