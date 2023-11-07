package service

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt"
	"lecture12/auth/internal/config"
	"lecture12/auth/internal/entity"
	"lecture12/auth/internal/repository"
	"time"
)

type Token struct {
	repo repository.TokenRepository
	cfg  *config.Config
}

func NewTokenService(repo repository.TokenRepository, cfg *config.Config) TokenService {
	return &Token{
		repo: repo,
		cfg:  cfg,
	}
}

func (t *Token) GenerateToken(ctx context.Context, id int, login string) (*entity.Token, error) {
	accessTokenClaims := jwt.MapClaims{
		"id":         id,
		"login":      login,
		"expires_at": time.Now().Add(time.Hour * 1).Unix(),
	}

	accessToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), accessTokenClaims)

	accessTokenString, err := accessToken.SignedString([]byte(t.cfg.SecretKey))
	if err != nil {
		return nil, fmt.Errorf("SignedString err: %v", err)
	}

	token := &entity.Token{
		ID:          id,
		Login:       login,
		AccessToken: accessTokenString,
	}

	err = t.repo.CreateToken(ctx, token)
	if err != nil {
		return nil, fmt.Errorf("CreateToken err: %v", err)
	}

	return token, nil
}
