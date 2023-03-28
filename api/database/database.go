package database

import (
	"am/office-check-in/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var connection *gorm.DB

func Connect(dsn string, name string) {
	// Add name to the dsn
	namedDsn := fmt.Sprintf("%s dbname=%s", dsn, name)

	var err error

	connection, err = gorm.Open(postgres.Open(namedDsn), &gorm.Config{})

	if err != nil {
		connection = createDatabase(dsn)
	}

	log.Printf("Connected to database %s", name)
}

func Migrate() {
	connection.AutoMigrate(&models.User{}, &models.TimeLog{})
}

func Connection() *gorm.DB {
	return connection
}

func createDatabase(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	dbname := os.Getenv("DATABASE_NAME")

	if dbname == "" {
		panic("DATABASE_NAME is not set")
	}

	createDatabaseCommand := fmt.Sprintf("CREATE DATABASE %s", dbname)
	db.Exec(createDatabaseCommand)

	// Close the connection to the database
	connection, err := db.DB()
	if err != nil {
		panic(err)
	}
	connection.Close()

	// Reconnect to the database with the new database name
	db, err = gorm.Open(postgres.Open(fmt.Sprintf("%s dbname=%s", dsn, dbname)), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}
