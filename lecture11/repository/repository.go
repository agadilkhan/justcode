package repository

import (
	"context"
	"lecture8/entity"
)

type Repository interface {
	CreateUser(ctx context.Context, u *entity.User) (int, error)
	GetUserByUsername(ctx context.Context, username string) (*entity.User, error)
	CreateReview(ctx context.Context, r *entity.Review) (int, error)
	GetReview(ctx context.Context, id int) (*entity.Review, error)
	GetAllReviews(ctx context.Context) (*[]entity.Review, error)
	DeleteReview(ctx context.Context, id int) (int, error)
	UpdateReview(ctx context.Context, r *entity.Review) (*entity.Review, error)
}
