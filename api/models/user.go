package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Picture   string    `json:"picture"`
	QrCodeId  string    `json:"qrCodeId"`
	QrCodeUrl string    `json:"qrCodeUrl"`
	TimeLogs  []TimeLog `json:"timeLogs" gorm:"ForeignKey:UserId"`
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
