package services

import (
	"github.com/yangliang0514/go-rest-api/database"
	"github.com/yangliang0514/go-rest-api/models"
)

func CreateUser(user models.User) (models.User, error) {
	if result := database.DB.Create(&user); result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}

func GetUserByEmail(email string) (models.User, error) {
	var user models.User
	if result := database.DB.Where("email = ?", email).First(&user); result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}

func GetUserById(id string) (models.User, error) {
	var user models.User
	if result := database.DB.Where("id = ?", id).First(&user); result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}
