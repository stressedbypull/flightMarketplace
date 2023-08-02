package database

import (
	"fmt"
	"log"
	"tui/flightmarketplace/config"
	"tui/flightmarketplace/model"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type PostgresClient struct {
	DB *gorm.DB
}

// Connect to database
// Connect to database
func Connect() (PostgresClient, error) {
	// connection string
	psqlconn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Config.Postgres.Host, config.Config.Postgres.Port,
		config.Config.Postgres.Username, config.Config.Postgres.Password, config.Config.Postgres.DBname, config.Config.Postgres.SSLMode)

	db, err := gorm.Open("postgres", psqlconn)
	if err != nil {
		return PostgresClient{}, fmt.Errorf("error during connection to database %s. Caused by %q", config.Config.Postgres.DBname, err)
	}
	log.Printf("Connected to %s", config.Config.Postgres.DBname)
	return PostgresClient{db}, nil
}

// Create tables if not present inside db
func (client *PostgresClient) ConfigureTables() error {
	// Migrate the schema
	err := client.DB.AutoMigrate(&model.User{}, &model.Flight{}, &model.Order{}).Error
	if err != nil {
		return fmt.Errorf("error when creating tables. Caused by %v", err)
	}
	log.Println("Tables created!")

	return nil
}
