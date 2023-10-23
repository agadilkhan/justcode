package pg

import (
	"context"
	"lecture8/entity"
)

func (p *Postgres) CreateReview(ctx context.Context, r *entity.Review) (int, error) {
	res := p.DB.WithContext(ctx).Create(r)
	if res.Error != nil {
		return 0, res.Error
	}

	return r.ID, nil
}

func (p *Postgres) GetReview(ctx context.Context, id int) (*entity.Review, error) {
	var r entity.Review

	res := p.DB.WithContext(ctx).Where("id = ?", id).First(&r)

	return &r, res.Error
}

func (p *Postgres) GetAllReviews(ctx context.Context) (*[]entity.Review, error) {
	var r []entity.Review

	res := p.DB.WithContext(ctx).Find(&r)

	return &r, res.Error
}

func (p *Postgres) DeleteReview(ctx context.Context, id int) (int, error) {
	res := p.DB.WithContext(ctx).Delete(&[]entity.Review{}, id)

	return id, res.Error
}

func (p *Postgres) UpdateReview(ctx context.Context, r *entity.Review) (*entity.Review, error) {
	var oldReview *entity.Review

	res := p.DB.WithContext(ctx).Where("id = ?", r.ID).First(&oldReview)
	if res.Error != nil {
		return nil, res.Error
	}

	oldReview.Title = r.Title
	oldReview.Content = r.Content

	res = p.DB.WithContext(ctx).Save(&oldReview)

	return oldReview, res.Error
}
