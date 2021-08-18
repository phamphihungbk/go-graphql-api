package services

import (
	"github.com/phamphihungbk/go-graphql/src/abstracts"
	"github.com/phamphihungbk/go-graphql/src/repositories"
)

type UserServiceInterface interface {
	abstracts.BaseServiceInterface
}

type UserService struct {
	*abstracts.BaseServiceInterface
	repository repositoies.UserRepositoryInterface
}

// @Summary UserService constructor
func NewUserService(repository repositories.UserRepositoryInterface) UserServiceInterface {
	baseService := abstracts.NewBaseService(repository).(*abstracts.BaseService)
	return &UserService{baseService, repository}
}
