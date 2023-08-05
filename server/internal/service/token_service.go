package service

import (
	jwt "github.com/golang-jwt/jwt/v5"
	"github.com/phamphihungbk/go-graphql/internal/model"
	"time"
)

const (
	Issuer          string = "my server"
	Audition        string = "api.example.com"
	ExpiredInterval int64  = 7 * 24 * 60 * 60
	PrivateKey      string = "private key"
)

type ITokenService interface {
	generate() (Token, error)
	validate() error
}

type TokenService struct {
}

func NewTokenService() *TokenService {
	return &TokenService{}
}

func (s *TokenService) generate(user *model.User) string {
	issuedTime := time.Now().Unix()
	expiredTime := issuedTime + ExpiredInterval

	token := jwt.NewWithClaims(
		jwt.SigningMethodES256,
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

func (s *TokenService) validate() error {
	return nil
}
