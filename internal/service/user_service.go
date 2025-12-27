package service

import (
	"errors"
	"go-rest-api/internal/model"
	"go-rest-api/internal/repository"
)

type UserService interface {
    GetUser(id uint) (*model.User, error)
    CreateUser(email, name string) error
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
