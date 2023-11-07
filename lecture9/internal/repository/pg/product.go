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

func (p *Postgres) GetProductByID(ctx context.Context, id uint) (*entity.Product, error) {
	var pr entity.Product

	res := p.DB.WithContext(ctx).Where("id=?", id).Find(&pr)

	return &pr, res.Error
}

func (p *Postgres) CreateProduct(ctx context.Context, pr *entity.Product) (uint, error) {
	res := p.DB.WithContext(ctx).Create(&pr)

	if res.Error != nil {
		return 0, res.Error
	}

	return pr.ID, nil
}

func (p *Postgres) UpdateProduct(ctx context.Context, pr *entity.Product) (*entity.Product, error) {
	var oldProduct *entity.Product

	p.DB.WithContext(ctx).Where("id=", pr.ID).Find(oldProduct)

	oldProduct = pr

	res := p.DB.WithContext(ctx).Save(oldProduct)

	return oldProduct, res.Error
}

func (p *Postgres) DeleteProduct(ctx context.Context, id uint) (uint, error) {
	res := p.DB.WithContext(ctx).Delete(&entity.Product{}, id)

	if res.Error != nil {
		return 0, res.Error
	}

	return id, nil
}
