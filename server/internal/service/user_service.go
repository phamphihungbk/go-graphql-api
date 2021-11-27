package service

import(
	"github.com/phamphihungbk/go-graphql/internal/repository"
	"github.com/phamphihungbk/go-graphql/internal/model"
)

type UserServiceInterface interface {
	GetModel() model.User
	GetItem(id uint) (model.User, error)
	Create(item model.User) model.User
	Update(item model.User) model.User
	Delete(id uint) error
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

func (c *UserService) GetItem(id uint) (model.User, error) {
	data, err := c.Repository.Find(id)
	return data, err
}

func (c *UserService) Create(item model.User) model.User {
	return c.Repository.Create(item)
}

func (c *UserService) Update(item model.User) model.User {
	return c.Repository.Update(item)
}

func (c *UserService) Delete(id uint) error {
	return c.Repository.Delete(id)
}
