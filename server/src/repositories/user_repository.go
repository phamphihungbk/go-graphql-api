package repositories

import (
	"gorm.io/gorm"
	"github.com/phamphihungbk/go-graphql/src/abstracts"
	"github.com/phamphihungbk/go-graphql/src/models"
)

type UserRepositoryInterface interface {
	abstracts.BaseRepositoryInterface
}

type UserRepository struct {
	*abstracts.BaseRepository
	model models.User
}

// @Summary UserRepository constructor
func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
	var model models.User
	repo := abstracts.NewBaseRepository(db, &model).(*abstracts.BaseRepository)
	return &UserRepository{repo, model}
}
