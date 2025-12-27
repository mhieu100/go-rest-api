package service

import (
	"errors"
	"go-rest-api/internal/model"
	"go-rest-api/internal/repository"
)

type UserService interface {
	GetUser(id uint) (*model.User, error)
	GetAllUsers() ([]*model.User, error)
	CreateUser(email, name string) error
	UpdateUser(id uint, email, name string) error
	DeleteUser(id uint) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{repo: r}
}

func (s *userService) GetUser(id uint) (*model.User, error) {
	return s.repo.FindByID(id)
}

func (s *userService) GetAllUsers() ([]*model.User, error) {
	return s.repo.FindAll()
}

func (s *userService) CreateUser(email, name string) error {
	if email == "" {
		return errors.New("email required")
	}

	user := &model.User{
		Email: email,
		Name:  name,
	}
	return s.repo.Create(user)
}

func (s *userService) UpdateUser(id uint, email, name string) error {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	if email != "" {
		user.Email = email
	}
	if name != "" {
		user.Name = name
	}

	return s.repo.Update(user)
}

func (s *userService) DeleteUser(id uint) error {
	return s.repo.Delete(id)
}
