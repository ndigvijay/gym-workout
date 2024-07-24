package database



import (
	// "github.com/joho/godotenv"
    "gorm.io/gorm"
    "fmt"
    "os"
    "log"
    "sync"
    "gorm.io/driver/postgres"
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
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbname, port)
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect database: %v", err)
    }
    return db
}