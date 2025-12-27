package repository

import "go-rest-api/internal/model"

type UserRepository interface {
    FindByID(id uint) (*model.User, error)
    Create(user *model.User) error
}