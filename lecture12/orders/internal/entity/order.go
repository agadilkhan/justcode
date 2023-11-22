package entity

import "gorm.io/gorm"

type Order struct {
	*gorm.Model
	ID         int `db:"id" gorm:"primary_key"`
	CustomerID int `db:"customer_id"`
	Customer   Customer
	Products   []OrderProduct
}

type OrderProduct struct {
	*gorm.Model
	ID        int `db:"id" gorm:"primary_key"`
	OrderID   int `db:"order_id"`
	ProductID int `db:"product_id"`
	Product   Product
	Count     int `db:"count"`
}
