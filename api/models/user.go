package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string
	Email     string
	Password  string
	Picture   string
	QrCodeId  string
	QrCodeUrl string
	TimeLogs  []TimeLog
}

type UserBody struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Picture  string `json:"picture"`
}

type UserLoginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
