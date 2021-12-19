package repository

import (
	"github.com/phamphihungbk/go-graphql/internal/model"
	"gorm.io/gorm"
)

type Pagination struct {
	Limit      int         `json:"limit,omitempty;query:limit"`
	Page       int         `json:"page,omitempty;query:page"`
	Sort       string      `json:"sort,omitempty;query:sort"`
	TotalRows  int64       `json:"total_rows"`
	TotalPages int         `json:"total_pages"`
	Rows       interface{} `json:"rows"`
}

type IUserRepository interface {
	ListAll(limit int, page int, sort string) ([]*model.User, error)
	Find(id int) (*model.User, error)
	Create(item *model.User) (*model.User, error)
	Update(item *model.User) *model.User
	Delete(id int) error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) ListAll(limit int, page int, sort string) ([]*model.User, error) {
	var users []*model.User
	query := ur.buildParamsQuery(limit, page, sort)
	err := ur.db.Scopes(query).Find(&users).Error
	if err != nil {
		return []*model.User{}, err
	}

	return users, nil
}

func (ur *UserRepository) Find(id int) (*model.User, error) {
	var user *model.User
	err := ur.db.First(&user, id).Error

	return user, err
}

func (ur *UserRepository) Create(item *model.User) (*model.User, error) {
	err := ur.db.Create(item).Error
	return item, err
}

func (ur *UserRepository) Update(item *model.User) *model.User {
	ur.db.Save(item)
	return item
}

func (ur *UserRepository) Delete(id int) error {
	item, err := ur.Find(id)
	if err != nil {
		return err
	}
	ur.db.Delete(item)

	return nil
}

func (ur *UserRepository) buildParamsQuery(limit int, page int, sort string) func(db *gorm.DB) *gorm.DB {
	page = (page - 1) * limit
	return func(db *gorm.DB) *gorm.DB {
		return ur.db.Offset(page).Limit(limit).Order(sort)
	}
}
