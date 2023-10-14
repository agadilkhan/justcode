package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"lecture8/entity"
	"lecture8/handler/dto"
	"lecture8/service/jwttoken"
	"time"
)

func (m *Manager) Register(ctx context.Context, u *entity.User) (int, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	u.Password = string(hashedPassword)
	userId, err := m.Repo.CreateUser(ctx, u)

	return userId, err
}

func (m *Manager) Login(ctx context.Context, username, password string) (*dto.LoginResponse, error) {
	user, err := m.Repo.GetUserByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("user not found")
		}

		return nil, fmt.Errorf("get user err: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("incorrect password: %w", err)
	}

	token := new(jwttoken.JWTToken)
	token.SecretKey = "lecture_8"

	accessToken, err := token.CreateToken(user.ID, time.Duration(15)*time.Minute)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{Token: accessToken}, nil
}
