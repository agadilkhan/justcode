package entity

import "gorm.io/gorm"

type Product struct {
	*gorm.Model
	ID    int     `db:"id" gorm:"primary_key"`
	Name  string  `db:"name"`
	Price float64 `db:"price"`
	Code  string  `db:"code"`
}
