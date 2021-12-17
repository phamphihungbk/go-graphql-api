package repository

import (
	"github.com/phamphihungbk/go-graphql/internal/model"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	GetModel() model.User
	Find(id uint) (model.User, error)
	Create(item model.User) model.User
	Update(item model.User) model.User
	Delete(id uint) error
}

type UserRepository struct {
	UserRepositoryInterface
	model model.User
	db    *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		model: model.User{},
		db:    db,
	}
}

func (c *UserRepository) GetModel() model.User {
	return c.model
}

func (c *UserRepository) Find(id uint) (model.User, error) {
	item := c.GetModel()
	err := c.db.First(item, id).Error
	return item, err
}

func (c *UserRepository) Create(item model.User) model.User {
	c.db.Create(item)
	return item
}

func (c *UserRepository) Update(item model.User) model.User {
	c.db.Save(item)
	return item
}

func (c *UserRepository) Delete(id uint) error {
	item, err := c.Find(id)
	if err != nil {
		return err
	}
	c.db.Delete(item)
	return nil
}
