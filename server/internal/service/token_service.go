package service

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/phamphihungbk/go-graphql-api/internal/model"
	"time"
)

const (
	Issuer          string = "my server"
	Audition        string = "api.example.com"
	ExpiredInterval int64  = 7 * 24 * 60 * 60
	PrivateKey      string = "private key"
)

type ITokenService interface {
	create(user *model.User) (string, error)
	validate(tokenString string) (jwt.MapClaims, error)
}

type TokenService struct {
}

func NewTokenService() *TokenService {
	return &TokenService{}
}

func (s *TokenService) create(user *model.User) (string, error) {
	issuedTime := time.Now().Unix()
	expiredTime := issuedTime + ExpiredInterval

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"iss":   Issuer,
			"sub":   user.ID,
			"aud":   Audition,
			"exp":   expiredTime,
			"iat":   issuedTime,
			"role":  "admin",
			"name":  user.Email,
			"email": user.Email,
		})

	return token.SignedString(PrivateKey)
}

func (s *TokenService) validate(tokenString string) (jwt.MapClaims, error) {
	var hmacSampleSecret []byte
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return hmacSampleSecret, nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return claims, nil
	}

	return nil, err
}
