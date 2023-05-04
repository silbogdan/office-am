package services

import (
	"am/office-check-in/database"
	"am/office-check-in/models"
	"errors"
)

func GetAllEmployees() ([]models.User, error) {
	// var users []models.User
	var users []models.User

	dbConnection := database.Connection()
	dbConnection.Preload("TimeLogs").Find(&users)

	if len(users) < 1 {
		return users, errors.New("could not find any users")
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
