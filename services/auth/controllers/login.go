package controllers

import (
	// "log"
	"os"
	"log"
	"net/http"
	"time"
	"errors"
	"github.com/alexedwards/argon2id"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ndigvijay/gym-workout/services/auth/db"
	"github.com/ndigvijay/gym-workout/services/auth/models"
)


func ParseToken(tokenString string) (*jwt.Token, error) {
	secretKey := os.Getenv("SECRET_KEY")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token, nil
}


func Login(c *gin.Context){
	tokenString, errToken := c.Cookie("Auth")
	if errToken == nil {
		token, errToken := ParseToken(tokenString)
		if errToken == nil && token.Valid {
			c.JSON(201, gin.H{
				"message": "already logged in",
			})
			return
		}
	}
	var body struct{
		Username string `json:"username"`
		Password string `json:"password"`
	}
	err:=c.Bind(&body);if err!=nil{
		log.Fatal("could not bind body")
		c.JSON(400,gin.H{
			"message":"could not bind body",
		})
		return
	}
	var existingUser models.User

	if existingUserError:=db.DB.Where("username = ?",body.Username).First(&existingUser);existingUserError.Error !=nil{
		c.JSON(404,gin.H{
			"message":"user not found",
		})
		return
	}

	match, err := argon2id.ComparePasswordAndHash(body.Password, existingUser.Password)
	if err != nil {
		log.Println("Error comparing passwords:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error comparing passwords",
		})
		return
	}

	if !match {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid username or password",
		})
		return
	}
	claims:=jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"iss":existingUser.ID,
		"exp":time.Now().Add(time.Hour * 24).Unix(),
	})
	token, err := claims.SignedString([]byte(os.Getenv("SECRET_KEY")));if err!=nil{
		log.Fatal(err)
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Auth",token,24*30*3600,"","",false,true)

	c.JSON(200,gin.H{
		"message":"logged in sucessfully",
	})
}

