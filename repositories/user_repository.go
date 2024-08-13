package repositories

import (
	"anciledigital/configs"
	"anciledigital/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		DB: configs.GetDB(),
	}
}

func (ur *UserRepository) GetUserByID(id int) (*models.User, error) {
	var user models.User
	result := ur.DB.Where("id = ?", id).
		First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (ur *UserRepository) GetAllUsers() ([]models.User, error) {
	var user models.User
	var users = make([]models.User, 0)
	err := ur.DB.Model(user).Find(&users).Error
	if err != nil {
		return make([]models.User, 0), err
	}
	return users, nil
}
