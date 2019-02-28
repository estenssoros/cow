package service

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type JWTService struct {
	ctx echo.Context
}

func NewJWT(ctx echo.Context) (*JWTService, error) {
	return &JWTService{
		ctx: ctx,
	}, nil
}

func (j *JWTService) NewToken() (string, error) {
	signingKey := []byte("{{.SecretKey}}")
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
		Issuer:    "cow",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(signingKey)
	if err != nil {
		return "", err
	}
	return ss, nil
}
