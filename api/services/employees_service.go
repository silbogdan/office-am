package services

import (
	"am/office-check-in/database"
	"am/office-check-in/minio_config"
	"am/office-check-in/models"
	"context"
	"errors"
	"net/url"
)

func GetAllEmployees() ([]models.User, error) {
	// var users []models.User
	var users []models.User

	dbConnection := database.Connection()
	dbConnection.Preload("TimeLogs").Find(&users)

	if len(users) < 1 {
		return users, errors.New("could not find any users")
	}

	// Get presigned file url from Minio
	minioClient, bucket := minio_config.Client()

	for idx, e := range users {
		picUrl, err := minioClient.PresignedGetObject(context.Background(), bucket, e.Picture, SEVEN_DAYS_IN_NANOSECONDS, url.Values{})
		qrUrl, err2 := minioClient.PresignedGetObject(context.Background(), bucket, e.QrCodeUrl, SEVEN_DAYS_IN_NANOSECONDS, url.Values{})

		if err != nil || err2 != nil {
			return nil, errors.New("could not sign file URL for user" + e.Name)
		}

		users[idx].Picture = picUrl.String()
		users[idx].QrCodeUrl = qrUrl.String()
	}

	return users, nil
}

func GetLogsForEmployee(id uint) ([]models.TimeLog, error) {
	var logs []models.TimeLog

	dbConnection := database.Connection()
	result := dbConnection.Where(&models.TimeLog{UserId: id}).Order("created_at DESC").Find(&logs)

	if result.RowsAffected < 1 {
		return logs, errors.New("no logs found")
	}

	return logs, nil
}
