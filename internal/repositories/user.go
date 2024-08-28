package repository

import (
	"errors"
	model "staycation/internal/models"
	database "staycation/pkg/databases"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *model.User) error
	FindByEmail(email string) (*model.User, error)
	FindByPhone(phone string) (*model.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository() UserRepository {
	return &userRepo{db: database.DB}
}

func (r *userRepo) Create(user *model.User) error {
	result := r.db.Create(&user)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return result.Error
	}

	return nil
}

func (r *userRepo) FindByEmail(email string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("email_exist")
		}

		return nil, err
	}

	return &user, nil
}

func (r *userRepo) FindByPhone(phone string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("phone = ?", phone).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("phone_exist")
		}

		return nil, err
	}

	return &user, nil
}
