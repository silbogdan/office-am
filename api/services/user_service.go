package services

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/aidarkhanov/nanoid"
	"github.com/minio/minio-go/v7"

	"am/office-check-in/database"
	"am/office-check-in/minio_config"
	"am/office-check-in/models"
)

func CreateUser(user models.UserBody) (models.User, error) {
	dbConnection := database.Connection()

	// Generate QR code data for user
	alphabet := nanoid.DefaultAlphabet
	size := 6

	qrCodeData, err := nanoid.Generate(alphabet, size)

	if err != nil {
		return models.User{}, errors.New("failed to generate QR code data")
	}

	// Create QR code with generated data
	qrServiceUrl := os.Getenv("QR_SERVICE_URL")

	if qrServiceUrl == "" {
		panic("QR_SERVICE_URL is not set")
	}

	url := fmt.Sprintf("%s%s%s", qrServiceUrl, "/create?data=", qrCodeData)
	resp, err := http.Post(url, "application/json", nil)

	if err != nil || resp.StatusCode != 200 {
		return models.User{}, errors.New("failed to create QR code")
	}
	defer resp.Body.Close()

	minioClient, bucket := minio_config.Client()

	// Upload QR code to minio
	info, err := minioClient.PutObject(context.Background(), bucket, fmt.Sprintf("%s%s", qrCodeData, ".png"),
		resp.Body, -1, minio.PutObjectOptions{ContentType: "image/png"})

	if err != nil {
		return models.User{}, errors.New("failed to upload QR code to minio")
	}

	// Create user
	newUser := models.User{
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		Picture:   user.Picture,
		QrCodeId:  qrCodeData,
		QrCodeUrl: info.Key,
	}

	dbConnection.Create(&newUser)

	return newUser, nil
}
