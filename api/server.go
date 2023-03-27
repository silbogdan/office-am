package main

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"

	"am/office-check-in/database"
)

func main() {
	err := godotenv.Load()

	dsn := os.Getenv("DATABASE_URL")
	dbName := os.Getenv("DATABASE_NAME")

	if dsn == "" || dbName == "" {
		panic("DATABASE_URL is not set")
	}

	database.Connect(dsn, dbName)

	if err != nil {
		panic(err)
	}

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
