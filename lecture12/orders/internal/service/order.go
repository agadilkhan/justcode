package service

import (
	"context"
	"lecture12/orders/internal/entity"
	"lecture12/orders/internal/repository"
)

type Order struct {
	repo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) OrderService {
	return &Order{
		repo,
	}
}

func (o *Order) GetAllOrders(ctx context.Context) (*[]entity.Order, error) {
	return o.repo.GetAllOrders(ctx)
}

func (o *Order) GetOrderByID(ctx context.Context, id int) (*entity.Order, error) {
	return o.repo.GetOrderByID(ctx, id)
}
