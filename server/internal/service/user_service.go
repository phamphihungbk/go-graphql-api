package service

import(
	"github.com/phamphihungbk/go-graphql/internal/repository"
	"github.com/phamphihungbk/go-graphql/internal/model"
)

type User model.User

type UserServiceInterface interface {
	GetModel() User
	GetItem(id uint) (User, error)
	Create(item User) User
	Update(item User) User
	Delete(id uint) error
}

type UserService struct {
	Repository repository.UserRepositoryInterface
}

func NewUserService(repository repository.UserRepositoryInterface) *UserService {
	return &UserService{repository}
}

func (c *UserService) GetModel() User {
	return c.Repository.GetModel()
}

func (c *UserService) GetItem(id uint) (User, error) {
	data, err := c.Repository.Find(id)
	return data, err
}

func (c *UserService) Create(item User) User {
	return c.Repository.Create(item)
}

func (c *UserService) Update(item User) User {
	return c.Repository.Update(item)
}

func (c *UserService) Delete(id uint) error {
	return c.Repository.Delete(id)
}
