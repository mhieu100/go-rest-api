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

func (r *userGorm) FindAll(page, limit int) ([]*model.User, int64, error) {
	var users []*model.User
	var total int64

	offset := (page - 1) * limit

	if err := r.db.Model(&model.User{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := r.db.Offset(offset).Limit(limit).Find(&users).Error
	if err != nil {
		return nil, 0, err
	}
	return users, total, nil
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
