package jwttoken

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"time"
)

var (
	ErrInvalidToken = errors.New("invalid token err")
	ErrExpiresToken = errors.New("expired token err")
)

type JWTToken struct {
	SecretKey string
}

func (j *JWTToken) CreateToken(userID uint, duration time.Duration) (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	payload := &JWTPayload{
		ID:        id,
		UserID:    userID,
		IssuedAt:  time.Now(),
		ExpiresAt: time.Now().Add(duration),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	return jwtToken.SignedString([]byte(j.SecretKey))
}

func (j *JWTToken) ValidateToken(token string) (*JWTPayload, error) {
	keyFunc := func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if ok {
			return []byte(j.SecretKey), nil
		}

		return nil, ErrInvalidToken
	}

	jwtToken, err := jwt.ParseWithClaims(token, &JWTPayload{}, keyFunc)
	if err != nil {
		validationErr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(validationErr, ErrExpiresToken) {
			return nil, ErrExpiresToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*JWTPayload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}
