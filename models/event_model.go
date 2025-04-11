package models

import "time"

type Event struct {
	Id          string    `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"date_time" binding:"required"`
	Users       []*User   `gorm:"many2many:user_events"`
}
