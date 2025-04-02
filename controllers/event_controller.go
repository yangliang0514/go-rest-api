package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yangliang0514/go-rest-api/models"
	"github.com/yangliang0514/go-rest-api/services"

	"github.com/google/uuid"
)

func GetAllEvents(c *gin.Context) {
	events, err := services.GetEvents()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"events": events})
}

func GetEvent(c *gin.Context) {
	event, err := services.GetEventById(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"event": event})
}

func CreateEvent(c *gin.Context) {
	var event models.Event

	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event.Id = uuid.New().String()

	event, err := services.CreateEvent(event)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Event created successfully", "event": event})
}

func UpdateEvent(c *gin.Context) {
	id := c.Param("id")

	event, err := services.GetEventById(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedEvent, err := services.UpdateEvent(id, event)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event updated successfully", "event": updatedEvent})
}

func DeleteEvent(c *gin.Context) {
	if err := services.DeleteEvent(c.Param("id")); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}
