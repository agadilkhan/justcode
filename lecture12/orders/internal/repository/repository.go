package repository

import (
	"context"
	"lecture12/orders/internal/entity"
	"lecture12/pkg/database/postgres"
)

type Repository struct {
	OrderRepository
}

func NewRepository(postgres *postgres.Postgres) *Repository {
	return &Repository{
		NewOrderRepository(postgres),
	}
}

type OrderRepository interface {
	GetAllOrders(ctx context.Context) (*[]entity.Order, error)
	GetOrderByID(ctx context.Context, id int) (*entity.Order, error)
}
