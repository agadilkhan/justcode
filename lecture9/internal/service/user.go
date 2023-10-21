package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"lecture9/internal/entity"
	"lecture9/internal/service/jwttoken"
	"time"
)

func (m *Manager) Register(ctx context.Context, u *entity.User) (uint, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	u.Password = string(hashedPassword)

	userID, err := m.Repo.CreateUser(ctx, u)

	return userID, err
}

func (m *Manager) Login(ctx context.Context, username, password string) (string, error) {
	user, err := m.Repo.GetUserByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("user not found")
		}

		return "", fmt.Errorf("get user err: %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", fmt.Errorf("incorrect password: %w", err)
	}

	token := new(jwttoken.JWTToken)
	token.SecretKey = "lecture_9"

	accessToken, err := token.CreateToken(user.ID, time.Duration(15)*time.Minute)

	return accessToken, err
}
