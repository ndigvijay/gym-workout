package database

import (
	// "github.com/joho/godotenv"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/ndigvijay/gym-workout/db/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var DBLock sync.Mutex 

// var DB *gorm.DB

func InitDB() *gorm.DB {
    DBLock.Lock()
    defer DBLock.Unlock()
    host := os.Getenv("DB_HOST")
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbname := os.Getenv("DB_NAME")
    port := 5432
	// log.Println(host,dbname,user,password,dbname)
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbname, port)
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    db.AutoMigrate(&models.User{},&models.WorkoutModel{})
    if err != nil {
        log.Fatalf("failed to connect database: %v", err)
    }
    return db
}