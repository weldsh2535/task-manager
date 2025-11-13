package handlers

import (
	"net/http"

	"github.com/weldsh2535/task-manager/models"
	"github.com/weldsh2535/task-manager/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Register
func Register(c *gin.Context) {
	var input UserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 10)

	user := models.User{Name: input.Name, Email: input.Email, Password: string(hash)}
	DB.Create(&user)
	c.JSON(http.StatusOK, user)
}

// Login
func Login(c *gin.Context) {
	var input UserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, _ := utils.GenerateJWT(user.ID)
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// List Users
func GetUsers(c *gin.Context) {
	var users []models.User
	DB.Preload("Projects").Find(&users)
	c.JSON(http.StatusOK, users)
}
