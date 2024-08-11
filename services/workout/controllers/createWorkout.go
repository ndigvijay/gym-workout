package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ndigvijay/gym-workout/services/workout/db"
	"github.com/ndigvijay/gym-workout/db/models"
)

func CreateWorkout(c *gin.Context){
	var body struct{
		Title string `json:"title"`
		Reps int `json:"reps"`
		Load int `json:"load"`
	}
	err:=c.Bind(&body);if err!=nil{
		c.JSON(400,gin.H{
			"message":"error binding body",
		})
		return 
	}
	var Workout models.WorkoutModel
	if body.Load == 0 || body.Reps == 0 || body.Title == "" {
		c.JSON(401, gin.H{
			"message": "no empty values allowed",
		})
		return
	}
	
	workout:=db.DB.Select(body.Title,body.Reps,body.Load).Create(&Workout)
	c.JSON(200,gin.H{
		"message":"workout added sucessfully",
		"workout":workout,
	})
}