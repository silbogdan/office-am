package models

import "gorm.io/gorm"

type TimeLog struct {
	gorm.Model
	UserId uint
	Type   string
}
