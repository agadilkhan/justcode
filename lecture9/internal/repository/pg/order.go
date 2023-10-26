package pg

import (
	"context"
	"lecture9/internal/entity"
)

func (p *Postgres) CreateOrder(ctx context.Context, o *entity.Order) (uint, error) {
	res := p.DB.WithContext(ctx).Create(&o)

	if res.Error != nil {
		return 0, res.Error
	}

	return o.ID, nil
}

func (p *Postgres) GetOrders(ctx context.Context) (*[]entity.Order, error) {
	var orders []entity.Order

	res := p.DB.WithContext(ctx).Find(&orders)

	return &orders, res.Error
}

func (p *Postgres) DeleteOrder(ctx context.Context, id uint) (uint, error) {
	res := p.DB.WithContext(ctx).Delete(&entity.Order{}, id)

	if res.Error != nil {
		return 0, res.Error
	}

	return id, nil
}

func (p *Postgres) UpdateOrder(ctx context.Context, o *entity.Order) (*entity.Order, error) {
	var oldOrder *entity.Order

	p.DB.WithContext(ctx).Where("id=", o.ID).Find(oldOrder)

	oldOrder = o

	res := p.DB.WithContext(ctx).Save(oldOrder)

	return oldOrder, res.Error
}
