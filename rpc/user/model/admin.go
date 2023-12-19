package model

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	Username string `json:"username" gorm:"type:varchar(100);not null"`
	Password string `json:"password" gorm:"type:varchar(200);not null"`
}
