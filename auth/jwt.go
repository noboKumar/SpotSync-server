package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	accessTokenExpiry  = time.Minute * 15
	refreshTokenExpiry = time.Hour * 24 * 7

	TokenTypeAccess  = "access"
	TokenTypeRefresh = "refresh"
)

type JwtClaims struct {
	UserId    uint   `json:"user_id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	TokenType string `json:"token_type"`

	jwt.RegisteredClaims
}

type JwtService interface {
	GenerateAccessToken(userId uint, name, email string, role string) (string, error)
	GenerateRefreshToken(userId uint, name, email string, role string) (string, error)
	ValidateToken(tokenString string) (*JwtClaims, error)
}

type jwtService struct {
	secretKey string
}

func NewJwtService(secretKey string) JwtService {
	if secretKey == "" {
		secretKey = "default_secret_key"
	}
	return &jwtService{secretKey: secretKey}
}

func (j *jwtService) GenerateAccessToken(userId uint, name string, email string, role string) (string, error) {
	claims := &JwtClaims{
		UserId:    userId,
		Name:      name,
		Email:     email,
		Role:      role,
		TokenType: TokenTypeAccess,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(accessTokenExpiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "SpotSync",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *jwtService) GenerateRefreshToken(userId uint, name, role string, email string) (string, error) {
	claims := &JwtClaims{
		UserId:    userId,
		Name:      name,
		Email:     email,
		Role:      role,
		TokenType: TokenTypeRefresh,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(refreshTokenExpiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "SpotSync",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *jwtService) ValidateToken(tokenString string) (*JwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(j.secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JwtClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
