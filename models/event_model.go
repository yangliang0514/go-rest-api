package models

import "time"

type Event struct {
	Id          string    `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"date_time" binding:"required"`
	UserId      string    `json:"user_id" gorm:"index"`
	User        *User     `json:"-" gorm:"foreignKey:UserId;references:Id"`
}
