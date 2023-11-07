package service

import (
	"context"
	"lecture12/auth/internal/config"
	"lecture12/auth/internal/entity"
	"lecture12/auth/internal/handler/http/dto"
	"lecture12/auth/internal/repository"
)

type Service struct {
	UserService
	TokenService
}

func NewService(repo *repository.Repository, cfg *config.Config) *Service {
	return &Service{
		NewUserService(repo.UserRepository, cfg, NewTokenService(repo, cfg)),
		NewTokenService(repo.TokenRepository, cfg),
	}
}

type UserService interface {
	Register(ctx context.Context, u *entity.User) (int, error)
	Login(ctx context.Context, login, password string) (*dto.LoginResponse, error)
}

type TokenService interface {
	GenerateToken(ctx context.Context, id int, login string) (*entity.Token, error)
}
