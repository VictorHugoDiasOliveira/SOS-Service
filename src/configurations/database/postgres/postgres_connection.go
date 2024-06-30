package postgres

import (
	"fmt"
	"sosservice/src/configurations/logger"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	SSLMode  string
}

var DB *gorm.DB

func InitDatabase(config *Config) *gorm.DB {

	logger.Info("Starting database",
		zap.String("journey", "Init Database"),
	)

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.Host,
		config.User,
		config.Password,
		config.DBName,
		config.Port,
		config.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	logger.Info("Connection established",
		zap.String("journey", "Init Database"),
	)

	return db
}
