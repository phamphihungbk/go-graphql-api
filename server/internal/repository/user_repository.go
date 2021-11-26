package repository

import (
	"gorm.io/gorm"
	"github.com/phamphihungbk/go-graphql/internal/model"
)

type User model.User

type UserRepositoryInterface interface {
	GetModel() User
	Find(id uint) (User, error)
	Create(item User) User
	Update(item User) User
	Delete(id uint) error
}

type UserRepository struct {
	UserRepositoryInterface
	model User
	db    *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	var user *User
	return &UserRepository{
		model: &user,
		db:    db,
	}
}

func (c *UserRepository) GetModel() User {
	return c.model
}

func (c *UserRepository) Find(id uint) (User, error) {
	item := c.GetModel()
	err := c.db.First(item, id).Error
	return item, err
}

func (c *UserRepository) Create(item User) User {
	c.db.Create(item)
	return item
}

func (c *UserRepository) Update(item User) User {
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
