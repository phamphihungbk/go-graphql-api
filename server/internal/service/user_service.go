package service

import (
	"github.com/phamphihungbk/go-graphql/internal/model"
	"github.com/phamphihungbk/go-graphql/internal/repository"
)

type UserServiceInterface interface {
	GetModel() model.User
	GetAllItems(parameters repository.ListParametersInterface) ([]model.User, error)
	GetItem(id uint) (model.User, error)
	CreateItem(item model.User) model.User
	UpdateItem(item model.User) model.User
	DeleteItem(id uint) error
}

type UserService struct {
	Repository *repository.UserRepository
}

func NewUserService(repository *repository.UserRepository) *UserService {
	return &UserService{repository}
}

func (c *UserService) GetModel() model.User {
	return c.Repository.GetModel()
}

func (c *UserService) GetAllItems(parameters repository.ListParametersInterface) ([]model.User, error) {
	data, err := c.Repository.ListAll(parameters)
	return data, err
}

func (c *UserService) GetItem(id uint) (model.User, error) {
	data, err := c.Repository.Find(id)
	return data, err
}

func (c *UserService) CreateItem(item model.User) model.User {
	return c.Repository.Create(item)
}

func (c *UserService) UpdateItem(item model.User) model.User {
	return c.Repository.Update(item)
}

func (c *UserService) DeleteItem(id uint) error {
	return c.Repository.Delete(id)
}
