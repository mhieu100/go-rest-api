package repository

import (
	"go-rest-api/internal/model"

	"gorm.io/gorm"
)

type userGorm struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userGorm{db: db}
}

func (r *userGorm) FindByID(id uint) (*model.User, error) {
	var user model.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userGorm) FindAll() ([]*model.User, error) {
	var users []*model.User
	err := r.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userGorm) Create(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *userGorm) Update(user *model.User) error {
	return r.db.Save(user).Error
}

func (r *userGorm) Delete(id uint) error {
	return r.db.Delete(&model.User{}, id).Error
}
