package database

import (
	"log"
	"rest-api/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Connection *gorm.DB
}

// ConnectDB initializes the database connection
func (db *Database) ConnectDB(databaseURL string) {
	var err error
	db.Connection, err = gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	} // Perform migrations
	err = db.Connection.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

}

// GetDB returns the database connection
func (db *Database) GetDB() *gorm.DB {
	return db.Connection
}
