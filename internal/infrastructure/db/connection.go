package db

import (
	"fmt"
	"github.com/kimosapp/poc/internal/infrastructure/configuration"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnection() (*gorm.DB, error) {
	dbConfig := configuration.GetDBConfig()
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.GetDatabaseHost(),
		dbConfig.GetDatabasePort(),
		dbConfig.GetDatabaseUser(),
		dbConfig.GetDatabasePassword(),
		dbConfig.GetDatabaseName(),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
