package service

import (
	"github.com/phamphihungbk/go-graphql/internal/model"
	"github.com/phamphihungbk/go-graphql/internal/repository"
)

type IUserService interface {
	GetAllItems(limit int, page int, sort string) ([]*model.User, error)
	GetItem(id int) (*model.User, error)
	CreateItem(item *model.User) (*model.User, error)
	UpdateItem(item *model.User) *model.User
	DeleteItem(id int) error
}

type UserService struct {
	repository repository.IUserRepository
}

func NewUserService(repository repository.IUserRepository) *UserService {
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

func (us *UserService) CreateItem(item *model.User) (*model.User, error) {
	data, err := us.repository.Create(item)
	return data, err
}

func (us *UserService) UpdateItem(item *model.User) *model.User {
	return us.repository.Update(item)
}

func (us *UserService) DeleteItem(id int) error {
	return us.repository.Delete(id)
}
