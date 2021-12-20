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
	Create(input model.CreateUserInput) (*model.User, error)
	Update(id int, input model.UpdateUserInput) (*model.User, error)
	Delete(id int) (bool, error)
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

func (ur *UserRepository) Create(input model.CreateUserInput) (*model.User, error) {
	err := ur.db.Create(input).Error
	return nil, err
}

func (ur *UserRepository) Update(id int, input model.UpdateUserInput) (*model.User, error) {
	user, err := ur.Find(id)
	if err != nil {
		return &model.User{}, err
	}
	ur.db.Model(user).Updates(input)

	return user, nil
}

func (ur *UserRepository) Delete(id int) (bool, error) {
	item, err := ur.Find(id)
	if err != nil {
		return false, err
	}
	ur.db.Delete(item)

	return true, nil
}

func (ur *UserRepository) buildParamsQuery(limit int, page int, sort string) func(db *gorm.DB) *gorm.DB {
	page = (page - 1) * limit
	return func(db *gorm.DB) *gorm.DB {
		return ur.db.Offset(page).Limit(limit).Order(sort)
	}
}
