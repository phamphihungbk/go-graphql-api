package service

import (
	"errors"
	"github.com/phamphihungbk/go-graphql-api/internal/model"
	"github.com/phamphihungbk/go-graphql-api/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	GetAllUsers(limit int, page int, sort string) (*model.UserConnection, error)
	GetUser(email string) (*model.User, error)
	IssueToken(payload model.LoginPayload) (*model.AccessToken, error)
	CreateUser(payload model.CreateUserPayload) (*model.User, error)
	UpdateUser(payload model.UpdateUserPayload) (*model.User, error)
	DeleteUser(email string) (bool, error)
}

type UserService struct {
	repository   *repository.UserRepository
	tokenService *TokenService
}

func NewUserService(repository *repository.UserRepository, tokenService *TokenService) *UserService {
	return &UserService{repository, tokenService}
}

func (s *UserService) GetAllItems(limit int, page int, sort string) (*model.UserConnection, error) {
	return s.repository.GetAll(limit, page, sort)
}

func (s *UserService) GetUser(email string) (*model.User, error) {
	return s.repository.FindByEmail(email)
}

func (s *UserService) IssueToken(payload model.LoginPayload) (*model.AccessToken, error) {
	user, err := s.FindByEmail(payload.Email)
	if err != nil {
		return nil, err
	}

	if _, err := checkPassword(payload.Password, user.Password); err != nil {
		return nil, err
	}

	token := s.tokenService.generate(user)

	return &model.AccessToken{token}, nil
}

func (s *UserService) CreateUser(payload model.CreateUserPayload) (*model.User, error) {
	password, err := hashPassword(payload.Password)
	if err != nil {
		return nil, err
	}
	payload.Password = password

	return s.repository.Create(payload)
}

func (s *UserService) UpdateUser(payload model.UpdateUserPayload) (*model.User, error) {
	return s.repository.Update(payload)
}

func (s *UserService) DeleteUser(email string) (bool, error) {
	return s.repository.Delete(email)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPassword(password string, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	if err == nil {
		return false, errors.New("division by zero")
	}

	return true, nil
}
