package main

import (
	"embed"
	"message-saver/internal/handler"
	"message-saver/internal/migrator"
	"message-saver/internal/repository"
	"message-saver/internal/service"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

const migrationsDir = "migrations"

//go:embed migrations/*.sql
var MigrationsFS embed.FS

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

	migrator := migrator.MustGetNewMigrator(MigrationsFS, migrationsDir)

	err = migrator.ApplyMigrations(db)

	if err != nil {
		panic(err)
	}

	log.Info("Migration was completed successfully")

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	log.Info("Starting server")
	handlers.InitRoutes(log).Run("0.0.0.0:" + os.Getenv("SERVER_PORT"))
}
