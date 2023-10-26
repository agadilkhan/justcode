package service

import (
	"context"
	"lecture8/entity"
	"lecture8/handler/dto"
)

type Service interface {
	Register(ctx context.Context, u *entity.User) (int, error)
	Login(ctx context.Context, username, password string) (*dto.LoginResponse, error)
	CreateReview(ctx context.Context, r *entity.Review) (int, error)
	GetReview(ctx context.Context, id int) (*entity.Review, error)
	GetAllReviews(ctx context.Context) (*[]entity.Review, error)
	DeleteReview(ctx context.Context, id int) (int, error)
	UpdateReview(ctx context.Context, r *entity.Review) (*entity.Review, error)
}
