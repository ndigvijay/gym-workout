package models

import "gorm.io/gorm"

type WorkoutModel struct{
	gorm.Model
	Title string `json:"title"`
	Reps int `json:"reps"`
	Load int `json:"load"`
}