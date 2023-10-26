package pg

import (
	"context"
	"lecture9/internal/entity"
)

func (p *Postgres) GetProducts(ctx context.Context) (*[]entity.Product, error) {
	var products []entity.Product

	res := p.DB.WithContext(ctx).Find(&products)

	return &products, res.Error
}
