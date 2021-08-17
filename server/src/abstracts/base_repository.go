package abstracts

import (
	"gorm.io/gorm"
)

type BaseRepositoryInterface interface {
}

type BaseRepository struct {
	BaseRepositoryInterface
	Db     *gorm.DB
}

// constructor
func NewBaseRepository(db *gorm.DB) *BaseRepository {
	return &BaseRepository{Db: db}
}
