package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var connection *gorm.DB

func Connect(dsn string, name string) {
	// Add name to the dsn
	namedDsn := fmt.Sprintf("%s dbname=%s", dsn, name)

	db, err := gorm.Open(postgres.Open(namedDsn), &gorm.Config{})

	if err != nil {
		createDatabase(dsn)
	}

	connection = db
}

func Connection() *gorm.DB {
	return connection
}

func createDatabase(dsn string) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	instance, _ := db.DB()
	defer instance.Close()

	dbname := os.Getenv("DATABASE_NAME")

	if dbname == "" {
		panic("DATABASE_NAME is not set")
	}

	createDatabaseCommand := fmt.Sprintf("CREATE DATABASE %s", dbname)
	db.Exec(createDatabaseCommand)
}
