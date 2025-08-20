package jwt

import (
	"clean-architecture/internal/domain/service/auth"
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrUnexpectedSigningMethod = errors.New("unexpected signing method")
	ErrTokenExpired            = errors.New("token expired")
	ErrInvalidClaims           = errors.New("invalid claims")
)

type Claims struct {
	UserID uint64 `json:"user_id"`
	jwt.RegisteredClaims
}

type JWT struct {
	Secret string
}

func NewJWT(secret string) auth.TokenService {
	return &JWT{Secret: secret}
}

// GenerateToken implements repository.TokenService interface
func (j *JWT) GenerateToken(ctx context.Context, userID uint64, duration time.Duration) (string, error) {
	now := time.Now()
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(duration)),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.Secret))
}

// VerifyToken implements repository.TokenService interface
func (j *JWT) VerifyToken(ctx context.Context, tokenString string) (*auth.TokenClaims, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return []byte(j.Secret), nil
	}
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, keyFunc, jwt.WithoutClaimsValidation()) // inspect token to handle refresh token
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, ErrInvalidClaims
	}

	// check token expired
	if time.Now().After(claims.ExpiresAt.Time) {
		return nil, ErrTokenExpired
	}

	return &auth.TokenClaims{
		UserID: claims.UserID,
		Exp:    claims.ExpiresAt.Unix(),
		Iat:    claims.IssuedAt.Unix(),
	}, nil
}
