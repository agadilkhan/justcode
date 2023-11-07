package repository

import (
	"context"
	"fmt"
	"lecture12/orders/internal/entity"
	"lecture12/pkg/database/postgres"
)

type Order struct {
	*postgres.Postgres
}

func NewOrderRepository(postgres *postgres.Postgres) OrderRepository {
	return &Order{
		postgres,
	}
}

func (o *Order) GetAllOrders(ctx context.Context) (*[]entity.Order, error) {
	var orders []entity.Order

	res := o.DB.WithContext(ctx).Find(&orders)
	if res.Error != nil {
		return nil, fmt.Errorf("failed to find all orders err: %v", res.Error)
	}

	return &orders, nil
}

func (o *Order) GetOrderByID(ctx context.Context, id int) (*entity.Order, error) {
	var order entity.Order

	res := o.DB.WithContext(ctx).Where("id = ?", id).First(&order)
	if res.Error != nil {
		return nil, fmt.Errorf("failed to find order by id err: %v", res.Error)
	}

	return &order, nil
}
