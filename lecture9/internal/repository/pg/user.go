package pg

import (
	"context"
	"lecture9/internal/entity"
)

func (p *Postgres) CreateUser(ctx context.Context, u *entity.User) (uint, error) {
	res := p.DB.WithContext(ctx).Create(u)
	if res.Error != nil {
		return 0, res.Error
	}

	return u.ID, nil
}

func (p *Postgres) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	var u entity.User

	res := p.DB.WithContext(ctx).Where("username=?", username).Find(&u)

	return &u, res.Error
}
