package repository

import (
	"context"
	"fmt"
	"lecture12/auth/internal/entity"
	"lecture12/pkg/database/postgres"
)

type User struct {
	*postgres.Postgres
}

func NewUserRepository(pg *postgres.Postgres) UserRepository {
	return &User{pg}
}

func (ur *User) CreateUser(ctx context.Context, u *entity.User) (int, error) {
	res := ur.DB.WithContext(ctx).Create(&u)
	if res.Error != nil {
		return 0, fmt.Errorf("failed to create user err: %v", res.Error)
	}

	return u.ID, nil
}

func (ur *User) GetUserByLogin(ctx context.Context, login string) (*entity.User, error) {
	var u entity.User

	res := ur.DB.WithContext(ctx).Where("login = ?", login).First(&u)
	if res.Error != nil {
		return nil, fmt.Errorf("failed to find user err: %v", res.Error)
	}

	return &u, nil
}
