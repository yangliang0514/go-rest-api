package services

import (
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
