package service

import (
	"context"
	"lecture9/internal/entity"
)

func (m *Manager) GetProducts(ctx context.Context) (*[]entity.Product, error) {
	return m.Repo.GetProducts(ctx)
}
