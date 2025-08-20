package auth

import (
	"context"
	"time"
)

type TokenClaims struct {
	UserID uint64 `json:"user_id"`
	Exp    int64  `json:"exp"`
	Iat    int64  `json:"iat"`
}

type TokenService interface {
	GenerateToken(ctx context.Context, userID uint64, duration time.Duration) (string, error)
	VerifyToken(ctx context.Context, tokenString string) (*TokenClaims, error)
}
