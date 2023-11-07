package repository

import (
	"context"
	"fmt"
	"lecture12/auth/internal/entity"
	"lecture12/pkg/database/postgres"
)

type Token struct {
	*postgres.Postgres
}

func NewTokenRepository(postgres *postgres.Postgres) TokenRepository {
	return &Token{
		postgres,
	}
}

func (tr *Token) CreateToken(ctx context.Context, t *entity.Token) error {
	res := tr.DB.WithContext(ctx).Create(&t)
	if res.Error != nil {
		return fmt.Errorf("failed to create token err: %v", res.Error)
	}

	return nil
}
