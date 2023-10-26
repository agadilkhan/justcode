package service

import (
	"context"
	"lecture8/entity"
)

func (m *Manager) CreateReview(ctx context.Context, r *entity.Review) (int, error) {
	return m.Repo.CreateReview(ctx, r)
}

func (m *Manager) GetReview(ctx context.Context, id int) (*entity.Review, error) {
	return m.Repo.GetReview(ctx, id)
}

func (m *Manager) GetAllReviews(ctx context.Context) (*[]entity.Review, error) {
	return m.Repo.GetAllReviews(ctx)
}

func (m *Manager) DeleteReview(ctx context.Context, id int) (int, error) {
	return m.Repo.DeleteReview(ctx, id)
}

func (m *Manager) UpdateReview(ctx context.Context, r *entity.Review) (*entity.Review, error) {
	return m.Repo.UpdateReview(ctx, r)
}
