package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

func Connect() {
    var err error
    connStr := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
        os.Getenv("DB_PORT"),
    )
    for i := 0; i < 10; i++ {
    DB, err = gorm.Open("postgres", connStr)
    if err == nil {
        break
    }
    log.Printf("Database not ready (attempt %d/10): %v", i+1, err)
    time.Sleep(2 * time.Second)
}
if err != nil {
    log.Fatal("Failed to connect to database after retries:", err)
}
    fmt.Println("✅ Database connected!")
}
