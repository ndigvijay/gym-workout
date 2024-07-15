package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ndigvijay/gym-workout/services/auth/db"
	// "github.com/ndigvijay/gym-workout/services/auth/controllers"
)

func main() {
	r := gin.Default()
	err:=db.Init()
	if err!=nil{
		log.Fatal(err)
	}
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// r.POST("/user/signup",controllers.Signup)
	r.Run()
}
