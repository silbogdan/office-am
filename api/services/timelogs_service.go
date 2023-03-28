package services

import (
	"am/office-check-in/database"
	"am/office-check-in/models"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func Scan(entranceType string) (models.TimeLog, error) {
	// Scan QR code (using camera)
	resp, err := http.Get(fmt.Sprintf("%s%s", os.Getenv("QR_SERVICE_URL"), "/scan"))

	if err != nil || resp.StatusCode != 200 {
		return models.TimeLog{}, err
	}
	defer resp.Body.Close()

	// Get QR code data
	qrCodeDataBytes, err := io.ReadAll(resp.Body)

	if err != nil {
		return models.TimeLog{}, err
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
		return models.TimeLog{}, errors.New("user is already checked in to the same entrance")
	}

	// Create timelog
	timelog = models.TimeLog{
		UserId: user.ID,
		Type:   entranceType,
	}

	dbConnection.Create(&timelog)

	// Return timelog
	return timelog, nil
}
