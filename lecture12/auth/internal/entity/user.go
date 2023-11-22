package entity

import "gorm.io/gorm"

type User struct {
	*gorm.Model
	ID        int    `json:"id" db:"id" gorm:"primary_key"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Login     string `json:"login" db:"login"`
	Password  string `json:"password" db:"password"`
}
