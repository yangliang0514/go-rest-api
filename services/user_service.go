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
