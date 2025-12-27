package repository

import "go-rest-api/internal/model"

type UserRepository interface {
	FindByID(id uint) (*model.User, error)
	FindAll(page, limit int) ([]*model.User, int64, error)
	Create(user *model.User) error
	Update(user *model.User) error
	Delete(id uint) error
}
