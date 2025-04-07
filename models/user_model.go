package models

type User struct {
	Id             string  `json:"id" gorm:"primaryKey"`
	Email          string  `json:"email" binding:"required" gorm:"unique"`
	Name           string  `json:"name" binding:"required"`
	Password       string  `json:"password" binding:"required" gorm:"-"`
	HashedPassword string  `json:"hashed_password"`
	Events         []Event `json:"events" gorm:"foreignKey:UserId;references:Id"`
}
