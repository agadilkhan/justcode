package service

import (
	"context"
	"lecture12/orders/internal/entity"
	"lecture12/orders/internal/repository"
)

type Service struct {
	OrderService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		NewOrderService(repo.OrderRepository),
	}
}

type OrderService interface {
	GetAllOrders(ctx context.Context) (*[]entity.Order, error)
	GetOrderByID(ctx context.Context, id int) (*entity.Order, error)
}
