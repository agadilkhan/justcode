package entity

import (
	"gorm.io/gorm"
)

type Token struct {
	*gorm.Model
	ID          int    `json:"id" db:"id" gorm:"primary_key"`
	Login       string `json:"login" db:"login"`
	AccessToken string `json:"token" db:"access_token"`
}
