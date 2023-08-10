package service

import "github.com/phamphihungbk/go-graphql-api/internal/repository"

type AuthenticationService struct {
	userRepo *repository.UserRepository
}

type IAuthenticationService interface {
	IsValidUser(email string) bool
}

func NewAuthenticationService(repo *repository.UserRepository) *AuthenticationService {
	return &AuthenticationService{
		userRepo: repo,
	}
}

func (u *AuthenticationService) IsValidUser(email string) bool {
	if user, _ := u.userRepo.FindByEmail(email); user != nil {
		return true
	}

	return false
}
