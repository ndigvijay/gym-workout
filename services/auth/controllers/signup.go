package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ndigvijay/gym-workout/services/auth/db"
	"github.com/ndigvijay/gym-workout/services/auth/models"
	// "log"
	"net/http"
)

func Signup(c *gin.Context) {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	err := c.Bind(&body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	newUser := models.User{
		Username: body.Username,
		Password: body.Password,
	}

	result := db.DB.Create(&newUser)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": newUser,
	})
}
