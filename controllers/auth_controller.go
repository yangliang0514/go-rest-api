package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/yangliang0514/go-rest-api/models"
	"github.com/yangliang0514/go-rest-api/services"
	"golang.org/x/crypto/bcrypt"
)

// hardcoded jwt key for simplicity
const jwtKey = "ce78504d2113ccc2378ac0e6754c6f8b8f757e16fbdc22c79a1700a3706742ce"

func Signup(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.Id = uuid.New().String()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user.HashedPassword = string(hashedPassword)
	newUser, err := services.CreateUser(user)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  newUser.Id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(jwtKey))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": gin.H{
		"id":    newUser.Id,
		"email": newUser.Email,
		"name":  newUser.Name,
	}, "token": tokenString})
}

func Login(c *gin.Context) {

}
