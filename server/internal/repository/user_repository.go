package repository

import (
	"github.com/phamphihungbk/go-graphql/internal/model"
	"gorm.io/gorm"
)

const DefaultPageSize = 20

type ListParametersInterface interface{}

type UserRepositoryInterface interface {
	GetModel() model.User
	ListAll(parameters ListParametersInterface) ([]model.User, error)
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
		model: &model.User,
		db:    db,
	}
}

func (c *UserRepository) GetModel() model.User {
	return c.model
}

func (c *UserRepository) ListAll(parameters ListParametersInterface) ([]model.User, error) {
	item := c.GetModel()
	query, err := buildParamsQuery(parameters)
	if err != nil {
		return []InterfaceEntity{}, err
	}
	result := query.Find(item)

	return result.RowsAffected, result.Error
}

func (c *UserRepository) Find(id uint) (model.User, error) {
	item := c.GetModel()
	result := c.db.First(item, id)
	return result.RowsAffected, result.Error
}

func (c *UserRepository) Create(item model.User) (model.User, error) {
	result = c.db.Create(item)
	return result.RowsAffected, result.Error
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

// TODO: enhance offset and pageSize calculation
func (c *UserRepository) buildParamsQuery(parameters ListParametersInterface) (*gorm.DB, error) {
	query := c.db
	pageSize := DefaultPageSize
	page := 0
	limit := pageSize
	offset := page * pageSize
	query = query.Offset(offset).Limit(limit)

	return query, nil
}
