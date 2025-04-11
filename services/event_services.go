package services

import (
	"errors"

	"github.com/yangliang0514/go-rest-api/database"
	"github.com/yangliang0514/go-rest-api/models"
)

func GetEvents() ([]models.Event, error) {
	var events []models.Event
	result := database.DB.Find(&events)

	if result.Error != nil {
		return []models.Event{}, result.Error
	}

	return events, nil
}

func GetEventById(id string) (models.Event, error) {
	var event models.Event
	result := database.DB.First(&event, "id = ?", id)

	if result.Error != nil {
		return models.Event{}, result.Error
	}

	return event, nil
}

func CreateEvent(event models.Event) (models.Event, error) {
	if result := database.DB.Create(&event); result.Error != nil {
		return models.Event{}, result.Error
	}

	return event, nil
}

func UpdateEvent(id string, eventUpdates models.Event) (models.Event, error) {
	event, err := GetEventById(id)

	if err != nil {
		return models.Event{}, err
	}

	if result := database.DB.Model(&event).Updates(eventUpdates); result.Error != nil {
		return models.Event{}, result.Error
	}

	return event, nil
}

func DeleteEvent(id string) error {
	if result := database.DB.Delete(&models.Event{}, id); result.Error != nil {
		return result.Error
	}
	return nil
}

func RegisterUserToEvent(eventId string, userId string) error {
	event, err := GetEventById(eventId)

	if err != nil {
		return err
	}

	user, err := GetUserById(userId)

	if err != nil {
		return err
	}

	if isUserRegistered(&event, &user) {
		return errors.New("user already registered")
	}

	if err := database.DB.Model(&event).Association("Users").Append(&user); err != nil {
		return err
	}

	return nil
}

func UnregisterUserFromEvent(eventId string, userId string) error {
	event, err := GetEventById(eventId)

	if err != nil {
		return err
	}

	user, err := GetUserById(userId)

	if err != nil {
		return err
	}

	if !isUserRegistered(&event, &user) {
		return errors.New("user wasn't registered")
	}

	if err := database.DB.Model(&event).Association("Users").Delete(&user); err != nil {
		return err
	}

	return nil
}

func isUserRegistered(event *models.Event, user *models.User) bool {
	var existingUsers []models.User

	if err := database.DB.Model(&event).Association("Users").Find(&existingUsers, "id = ?", user.Id); err != nil {
		return false
	}

	return len(existingUsers) > 0
}
