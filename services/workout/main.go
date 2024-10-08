package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ndigvijay/gym-workout/services/workout/controllers"
)


func main(){
	r:=gin.Default()
	r.GET("/ping",func(c *gin.Context) {
		c.JSON(200,gin.H{
			"message":"pong",
		})
	})
	r.POST("/add-workout",controllers.CreateWorkout)
	r.Run()
}