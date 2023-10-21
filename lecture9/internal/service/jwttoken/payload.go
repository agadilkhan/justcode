package jwttoken

import (
	"github.com/google/uuid"
	"time"
)

type JWTPayload struct {
	ID        uuid.UUID `json:"id"`
	UserID    uint      `json:"user_id"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

func (j JWTPayload) Valid() error {
	if time.Now().After(j.ExpiresAt) {
		return ErrExpiresToken
	}

	return nil
}
