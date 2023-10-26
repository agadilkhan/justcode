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

func (p *Postgres) GetUsers(ctx context.Context) (*[]entity.User, error) {
	var u *[]entity.User

	res := p.DB.WithContext(ctx).Find(&u)

	return u, res.Error
}

func (p *Postgres) UpdateUser(ctx context.Context, u *entity.User) (*entity.User, error) {
	var oldUser *entity.User

	p.DB.WithContext(ctx).Where("id=", u.ID).Find(oldUser)

	oldUser = u

	res := p.DB.WithContext(ctx).Save(oldUser)

	return oldUser, res.Error
}

func (p *Postgres) DeleteUser(ctx context.Context, id uint) (uint, error) {
	res := p.DB.WithContext(ctx).Delete(&entity.User{}, id)

	if res.Error != nil {
		return 0, res.Error
	}

	return id, nil
}
