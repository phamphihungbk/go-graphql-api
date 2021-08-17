package abstracts

type BaseServiceInterface interface {
}

type BaseService struct {
	BaseServiceInterface
	Repository BaseRepositoryInterface
}

// constructor
func NewBaseService(repository BaseRepositoryInterface) *BaseService {
	return &BaseService{Repository: repository}
}
