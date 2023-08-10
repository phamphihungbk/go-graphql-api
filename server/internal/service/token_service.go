package service

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/phamphihungbk/go-graphql-api/internal/model"
)

const (
	Issuer          string = "my server"
	Audition        string = "api.example.com"
	ExpiredInterval int64  = 7 * 24 * 60 * 60
	PrivateKey      string = "private key"
)

func CreateToken(user *model.User) (string, error) {
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

func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	var hmacSampleSecret []byte
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return hmacSampleSecret, nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	if !isExpiredToken(claims) {
		return nil, err
	}

	if !isValidAudition(claims) {
		return nil, err
	}

	return claims, nil
}

func isExpiredToken(claims jwt.MapClaims) bool {
	iat := claims["iat"].(int64)
	exp := claims["exp"].(int64)
	now := time.Now().Unix()

	if iat > exp {
		return false
	}

	if now > exp {
		return false
	}

	return true
}

func isValidAudition(claims jwt.MapClaims) bool {
	aud := claims["aud"].(string)

	if aud == Audition {
		return true
	}

	return false
}
