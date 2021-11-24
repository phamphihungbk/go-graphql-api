package repository

import (
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	GetModel() UserModel
	Find(id uint) (UserModel, error)
	Create(item UserModel) UserModel
	Update(item UserModel) UserModel
	Delete(id uint) error
}

type UserRepository struct {
	UserRepositoryInterface
	model UserModel
	db    *gorm.DB
}

func NewUserRepository(db *gorm.DB, model UserModel) *UserRepository {
	return &UserRepository{
		model: model,
		db:    db,
	}
}

func (c *UserRepository) GetModel() UserModel {
	return c.model
}

func (c *UserRepository) Find(id uint) (UserModel, error) {
	item := c.GetModel()
	err := c.db.First(item, id).Error
	return item, err
}

func (c *UserRepository) Create(item UserModel) UserModel {
	c.db.Create(item)
	return item
}

func (c *UserRepository) Update(item UserModel) UserModel {
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
