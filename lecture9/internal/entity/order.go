package entity

import (
	"gorm.io/gorm"
)

type Order struct {
	*gorm.Model
	UserID   uint            `json:"user_id"`
	Products []*OrderProduct `gorm:"many2many:order_products" json:"order_products"`
}

type OrderProduct struct {
	*gorm.Model
	*Product
	Quantity int
}
