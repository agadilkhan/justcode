package service

import (
	"context"
	"lecture9/internal/entity"
)

func (m *Manager) CreateOrder(ctx context.Context, o *entity.Order) (uint, error) {
	return m.Repo.CreateOrder(ctx, o)
}

func (m *Manager) GetOrders(ctx context.Context) (*[]entity.Order, error) {
	return m.Repo.GetOrders(ctx)
}

func (m *Manager) DeleteOrder(ctx context.Context, id uint) (uint, error) {
	return m.Repo.DeleteOrder(ctx, id)
}

func (m *Manager) UpdateOrder(ctx context.Context, o *entity.Order) (*entity.Order, error) {
	return m.Repo.UpdateOrder(ctx, o)
}
