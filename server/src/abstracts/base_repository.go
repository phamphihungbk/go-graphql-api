package abstracts

import (
	"gorm.io/gorm"
)

type BaseRepositoryInterface interface {
	GetModel() BaseModel
	Find(id uint) (BaseModel, error)
	Create(item BaseModel) BaseModel
	Update(item BaseModel) BaseModel
	Delete(id uint) error
}

type BaseRepository struct {
	BaseRepositoryInterface
	model BaseModel
	db    *gorm.DB
}

// @Summary BaseRepository constructor
func NewBaseRepository(db *gorm.DB, model BaseModel) *BaseRepository {
	return &BaseRepository{
		model: model,
		db:    db,
	}
}

// @Summary Get model
func (c BaseRepository) GetModel() BaseModel {
	return c.model
}

// @Summary find item with id
func (c BaseRepository) Find(id uint) (BaseModel, error) {
	item := c.GetModel()
	err := c.db.First(item, id).Error
	return item, err
}

// @Summary create item with given data
func (c BaseRepository) Create(item BaseModel) BaseModel {
	c.db.Create(item)
	return item
}

// @Summary update item with given data
func (c BaseRepository) Update(item BaseModel) BaseModel {
	c.db.Save(item)
	return item
}

// @Summary delete item with given id
func (c BaseRepository) Delete(id uint) error {
	item, err := c.Find(id)
	if err != nil {
		return err
	}
	c.db.Delete(item)
	return nil
}
