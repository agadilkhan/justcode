package pg

import (
	"context"
	"lecture8/entity"
)

func (p *Postgres) CreateUser(ctx context.Context, u *entity.User) (int, error) {
	res := p.DB.WithContext(ctx).Create(u)
	if res.Error != nil {
		return 0, res.Error
	}
	return u.ID, nil
}

func (p *Postgres) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	user := new(entity.User)
	res := p.DB.WithContext(ctx).Where("username = ?", username).Find(&user)
	return user, res.Error
}
