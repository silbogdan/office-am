package main

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"am/office-check-in/database"
	"am/office-check-in/minio_config"
	"am/office-check-in/routes"
)

func main() {
	err := godotenv.Load()

	dsn := os.Getenv("DATABASE_URL")
	dbName := os.Getenv("DATABASE_NAME")

	if dsn == "" || dbName == "" {
		panic("DATABASE_URL is not set")
	}

	database.Connect(dsn, dbName)
	database.Migrate()

	minioEndpoint := os.Getenv("MINIO_ENDPOINT")
	minioBucket := os.Getenv("MINIO_BUCKET")
	minioAccessKeyID := os.Getenv("MINIO_ACCESS_KEY")
	minioSecretAccessKey := os.Getenv("MINIO_SECRET_KEY")

	if minioEndpoint == "" || minioBucket == "" || minioAccessKeyID == "" || minioSecretAccessKey == "" {
		panic("MINIO_ENDPOINT, MINIO_BUCKET, MINIO_ACCESS_KEY_ID, MINIO_SECRET_ACCESS_KEY are not set")
	}

	minio_config.Connect(minioEndpoint, minioBucket, minioAccessKeyID, minioSecretAccessKey)

	if err != nil {
		panic(err)
	}

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	routes.AddAuth(e)
	routes.AddTimeLogs(e)
	routes.AddFile(e)

	e.Logger.Fatal(e.Start(":1323"))
}
