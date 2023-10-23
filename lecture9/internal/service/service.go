package service

import (
	"context"
	"lecture9/internal/entity"
)

type Service interface {
	Register(ctx context.Context, user *entity.User) (uint, error)
	Login(ctx context.Context, username, password string) (string, error)
	CreateOrder(ctx context.Context, o *entity.Order) (uint, error)
	GetOrders(ctx context.Context) (*[]entity.Order, error)
	GetOrderByID(ctx context.Context, id uint) (*entity.Order, error)
	DeleteOrder(ctx context.Context, id uint) (uint, error)
	UpdateOrder(ctx context.Context, o *entity.Order) (*entity.Order, error)
	GetProducts(ctx context.Context) (*[]entity.Product, error)
}
