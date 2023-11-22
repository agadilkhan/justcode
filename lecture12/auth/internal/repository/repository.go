package repository

import (
	"context"
	"lecture12/auth/internal/entity"
	"lecture12/pkg/database/postgres"
)

type Repository struct {
	UserRepository
	TokenRepository
}

func NewRepository(postgres *postgres.Postgres) *Repository {
	return &Repository{
		NewUserRepository(postgres),
		NewTokenRepository(postgres),
	}
}

type UserRepository interface {
	CreateUser(ctx context.Context, u *entity.User) (int, error)
	GetUserByLogin(ctx context.Context, login string) (*entity.User, error)
}

type TokenRepository interface {
	CreateToken(ctx context.Context, t *entity.Token) error
}
