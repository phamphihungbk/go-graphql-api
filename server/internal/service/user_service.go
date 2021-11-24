package service

type UserServiceInterface interface {
	GetModel() UserModel
	GetItem(id uint) (UserModel, error)
	Create(item UserModel) UserModel
	Update(item UserModel) UserModel
	Delete(id uint) error
}

type UserService struct {
	Repository UserRepositoryInterface
}

func NewUserService(repository UserRepositoryInterface) *UserService {
	return &UserService{repository}
}

func (c *UserService) GetModel() UserModel {
	return c.Repository.GetModel()
}

func (c *UserService) GetItem(id uint) (UserModel, error) {
	data, err := c.Repository.Find(id)
	return data, err
}

func (c *UserService) Create(item UserModel) UserModel {
	return c.Repository.Create(item)
}

func (c *UserService) Update(item UserModel) UserModel {
	return c.Repository.Update(item)
}

func (c *UserService) Delete(id uint) error {
	return c.Repository.Delete(id)
}
