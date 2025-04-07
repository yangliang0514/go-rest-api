package models

type User struct {
	Id             string  `json:"id" gorm:"primaryKey"`
	Email          string  `json:"email" binding:"required"`
	Name           string  `json:"name" binding:"required"`
	HashedPassword string  `json:"hashed_password" binding:"required"`
	Events         []Event `json:"events" gorm:"foreignKey:UserId;references:Id"`
}
