package services

import (
	"github.com/phamphihungbk/go-graphql/src/abstracts"
	"github.com/phamphihungbk/go-graphql/src/repositories"
)

type UserServiceInterface interface {
	abstracts.BaseServiceInterface
}

type UserService struct {
	*abstracts.BaseService
	userRepository repositories.UserRepositoryInterface
}

// @Summary UserService constructor
func NewUserService(userRepository repositories.UserRepositoryInterface) UserServiceInterface {
	baseService := abstracts.NewBaseService(userRepository)
	return &UserService{baseService, userRepository}
}
