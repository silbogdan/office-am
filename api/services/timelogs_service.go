package services

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"am/office-check-in/database"
	"am/office-check-in/minio_config"
	"am/office-check-in/models"
)

func Scan(entranceType string) (models.TimeLogResponse, error) {
	// Scan QR code (using camera)
	resp, err := http.Get(fmt.Sprintf("%s%s", os.Getenv("QR_SERVICE_URL"), "/scan"))

	if err != nil || resp.StatusCode != 200 {
		return models.TimeLogResponse{}, err
	}
	defer resp.Body.Close()

	// Get QR code data
	qrCodeDataBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		return models.TimeLogResponse{}, err
	}

	qrCodeData := string(qrCodeDataBytes)

	// Check if user exists
	dbConnection := database.Connection()
	var user models.User
	dbConnection.Where("qr_code_id = ?", qrCodeData).First(&user)

	// Check if user is already checked in (based on entranceType)
	// 1. Get last timelog
	var timelog models.TimeLog
	dbConnection.Where("user_id = ?", user.ID).Last(&timelog)

	// 2. Check if timelog type is the same as entranceType
	if timelog.Type == entranceType {
		return models.TimeLogResponse{}, errors.New("user is already checked in to the same entrance")
	}

	// Create timelog
	timelog = models.TimeLog{
		UserId: user.ID,
		Type:   entranceType,
	}

	dbConnection.Create(&timelog)

	minioClient, bucket := minio_config.Client()

	picUrl, err := minioClient.PresignedGetObject(context.Background(), bucket, user.Picture, SEVEN_DAYS_IN_NANOSECONDS, url.Values{})

	if err != nil {
		return models.TimeLogResponse{}, errors.New("failed to sign url")
	}

	// Return timelog
	return models.TimeLogResponse{
		Type:    timelog.Type,
		Name:    user.Name,
		Picture: picUrl.String(),
	}, nil
}

func GetAllTimeLogs() ([]models.TimeLog, error) {
	var timeLogs []models.TimeLog

	dbConnection := database.Connection()
	dbConnection.Order("created_at desc").Find(&timeLogs)

	if len(timeLogs) < 1 {
		return timeLogs, errors.New("did not find any logs")
	}

	return timeLogs, nil
}
