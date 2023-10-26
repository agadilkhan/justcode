package repository

import (
	"context"
	"lecture9/internal/entity"
)

type Repository interface {
	CreateUser(ctx context.Context, u *entity.User) (uint, error)
	GetUserByUsername(ctx context.Context, username string) (*entity.User, error)
	GetUsers(ctx context.Context) (*[]entity.User, error)
	UpdateUser(ctx context.Context, u *entity.User) (*entity.User, error)
	DeleteUser(ctx context.Context, id uint) (uint, error)
	CreateOrder(ctx context.Context, o *entity.Order) (uint, error)
	GetOrders(ctx context.Context) (*[]entity.Order, error)
	GetOrderByID(ctx context.Context, di uint) (*entity.Order, error)
	DeleteOrder(ctx context.Context, id uint) (uint, error)
	UpdateOrder(ctx context.Context, o *entity.Order) (*entity.Order, error)
	GetProducts(ctx context.Context) (*[]entity.Product, error)
	GetProductByID(ctx context.Context, id uint) (*entity.Product, error)
	CreateProduct(ctx context.Context, pr *entity.Product) (uint, error)
	UpdateProduct(ctx context.Context, pr *entity.Product) (*entity.Product, error)
	DeleteProduct(ctx context.Context, id uint) (uint, error)
}
