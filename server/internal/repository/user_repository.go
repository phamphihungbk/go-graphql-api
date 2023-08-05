package repository

import (
	"github.com/phamphihungbk/go-graphql-api/internal/model"
	"gorm.io/gorm"
)

type IUserRepository interface {
	GetAll(limit int, page int, sort string) ([]*model.User, error)
	FindByEmail(email string) (*model.User, error)
	Create(input model.CreateUserPayload) (*model.User, error)
	Update(input model.UpdateUserPayload) (*model.User, error)
	Delete(email string) (bool, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetAll(limit int, page int, sort string) (*model.UserConnection, error) {
	var users []*model.User
	query := paginate(limit, page, sort)
	err := r.db.Scopes(query).Find(&users).Error
	pageInfo := &model.PageInfo{limit, page, sort}

	if err != nil {
		return &model.UserConnection{nil, pageInfo}, err
	}

	return &model.UserConnection{users, pageInfo}, nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	var user = &model.User{Email: email}
	err := r.db.Where(&user).First(&user).Error

	return user, err
}

func (r *UserRepository) Create(input model.CreateUserPayload) (*model.User, error) {
	err := r.db.Create(input).Error
	return nil, err
}

func (r *UserRepository) Update(input model.UpdateUserPayload) (*model.User, error) {
	user, err := r.FindByEmail(input.Email)
	if err != nil {
		return &model.User{}, err
	}
	r.db.Model(user).Updates(input)

	return user, nil
}

func (r *UserRepository) Delete(email string) (bool, error) {
	item, err := r.FindByEmail(email)
	if err != nil {
		return false, err
	}
	r.db.Delete(item)

	return true, nil
}

func paginate(limit int, page int, sort string) func(db *gorm.DB) *gorm.DB {
	offset := (page - 1) * limit
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(limit).Order(sort)
	}
}
