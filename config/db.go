package config

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
    "os"
    "alephcode-backend/models"
)



var DB *gorm.DB

func ConnectDB() {
    dsn := os.Getenv("DATABASE_URL")
if dsn == "" {
	log.Fatal("DATABASE_URL not set")
}
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Database connection failed: %v", err)
    }

    db.AutoMigrate(&models.Mission{})
    DB = db
}
