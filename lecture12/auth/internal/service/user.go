package service

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"lecture12/auth/internal/config"
	"lecture12/auth/internal/entity"
	"lecture12/auth/internal/handler/http/dto"
	"lecture12/auth/internal/repository"
)

type User struct {
	repo         repository.UserRepository
	tokenService TokenService
	cfg          *config.Config
}

func NewUserService(repo repository.UserRepository, cfg *config.Config, tokenService TokenService) UserService {
	return &User{
		repo:         repo,
		cfg:          cfg,
		tokenService: tokenService,
	}
}

func (us *User) Register(ctx context.Context, u *entity.User) (int, error) {
	generatedHash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, fmt.Errorf("hashed password err: %v", err)
	}

	u.Password = string(generatedHash)

	id, err := us.repo.CreateUser(ctx, u)
	if err != nil {
		return 0, fmt.Errorf("create user err: %v", err)
	}

	return id, nil
}

func (us *User) Login(ctx context.Context, login, password string) (*dto.LoginResponse, error) {
	u, err := us.repo.GetUserByLogin(ctx, login)
	if err != nil {
		return nil, fmt.Errorf("get user err: %v", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("compare hash and password err: %v", err)
	}

	token, err := us.tokenService.GenerateToken(ctx, u.ID, u.Login)
	if err != nil {
		return nil, fmt.Errorf("GenerateToken err: %v", err)
	}

	return &dto.LoginResponse{
		Login:       token.Login,
		ID:          token.ID,
		AccessToken: token.AccessToken,
	}, nil
}
