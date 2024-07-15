package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ndigvijay/gym-workout/services/auth/db"
	"github.com/ndigvijay/gym-workout/services/auth/models"

	// "log"
	"net/http"

	"github.com/alexedwards/argon2id"
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
	var existingUser models.User

	if existingUserError:=db.DB.Where("username=?",body.Username).First(&existingUser);existingUserError.Error==nil{
		c.JSON(402,gin.H{
			"messsage":"user already exists",
		})
		return
	}	
	hash,err:=argon2id.CreateHash(body.Password,argon2id.DefaultParams);if err!=nil{
		log.Fatal(err)
	}
	newUser:= models.User{
		Username: body.Username,
		Password: hash,
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
