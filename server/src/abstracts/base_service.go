package abstracts

type BaseServiceInterface interface {
	GetModel() BaseModel
	GetItem(id uint) (BaseModel, error)
	Create(item BaseModel) BaseModel
	Update(item BaseModel) BaseModel
	Delete(id uint) error
}

type BaseService struct {
	Repository BaseRepositoryInterface
}

// @Summary NewBaseService constructor
func NewBaseService(repository BaseRepositoryInterface) *BaseService {
	return &BaseService{repository}
}

// @Summary Get model
func (c BaseService) GetModel() BaseModel {
	return c.Repository.GetModel()
}

func (c BaseService) GetItem(id uint) (BaseModel, error) {
	return c.Repository.Find(id)
}

// @Summary Create
func (c BaseService) Create(item BaseModel) BaseModel {
	return c.Repository.Create(item)
}

// @Summary Update
func (c BaseService) Update(item BaseModel) BaseModel {
	return c.Repository.Update(item)
}

// @Summary Delete
func (c BaseService) Delete(id uint) error {
	return c.Repository.Delete(id)
}
