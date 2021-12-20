package service

import (
	"github.com/phamphihungbk/go-graphql/internal/model"
	"github.com/phamphihungbk/go-graphql/internal/repository"
)

type IUserService interface {
	GetAllItems(limit int, page int, sort string) ([]*model.User, error)
	GetItem(id int) (*model.User, error)
	CreateItem(input model.CreateUserInput) (*model.User, error)
	UpdateItem(id int, input model.UpdateUserInput) (*model.User, error)
	DeleteItem(id int) (bool, error)
}

type UserService struct {
	repository *repository.UserRepository
}

func NewUserService(repository *repository.UserRepository) *UserService {
	return &UserService{repository}
}

func (us *UserService) GetAllItems(limit int, page int, sort string) ([]*model.User, error) {
	data, err := us.repository.ListAll(limit, page, sort)
	return data, err
}

func (us *UserService) GetItem(id int) (*model.User, error) {
	data, err := us.repository.Find(id)
	return data, err
}

func (us *UserService) CreateItem(input model.CreateUserInput) (*model.User, error) {
	return us.repository.Create(input)
}

func (us *UserService) UpdateItem(id int, input model.UpdateUserInput) (*model.User, error) {
	return us.repository.Update(id, input)
}

func (us *UserService) DeleteItem(id int) (bool, error) {
	return us.repository.Delete(id)
}
