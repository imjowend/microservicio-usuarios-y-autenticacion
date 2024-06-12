package api

import (
	"errors"
	"fmt"
	usr "github.com/imjowend/microservicio-usuarios-y-autenticacion/internal/core/user"
	"github.com/joho/godotenv"
	"os"
)

type DBConfig struct {
	Host     string
	User     string
	Password string
}

type Dependencies struct {
	Repository usr.RepositoryPort
	DBConfig   DBConfig
	RouterPort string
}

func loadEnv() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}
	return nil
}

func GetDBConfig() (DBConfig, error) {
	dbConfig := DBConfig{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	}
	if dbConfig.Host == "" || dbConfig.User == "" || dbConfig.Password == "" {
		return DBConfig{}, errors.New("incomplete database configuration")
	}
	return dbConfig, nil
}
func Config() (*Dependencies, error) {
	if err := loadEnv(); err != nil {
		return nil, err
	}

	dbConfig, err := GetDBConfig()
	if err != nil {
		return nil, err
	}

	routerPort := os.Getenv("ROUTER_PORT")
	if routerPort == "" {
		return nil, errors.New("empty router port")
	}

	return &Dependencies{
		Repository: usr.NewRepository(),
		DBConfig:   dbConfig,
		RouterPort: routerPort,
	}, nil
}
