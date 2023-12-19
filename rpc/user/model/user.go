package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username    string `json:"username" gorm:"type:varchar(100);not null"`
	Email       string `json:"email" gorm:"type:varchar(100);not null"`
	Password    string `json:"password" gorm:"type:varchar(200);not null"`
	Description string `json:"description" gorm:"type:varchar(100);not null"`
	Status      string `json:"status" gorm:"type:varchar(50);not null"`
}
