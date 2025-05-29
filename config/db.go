package config

import (
	"alephcode-backend/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	log.Println("✅ Connected to DB, running AutoMigrate...")

	db.AutoMigrate(
		&models.Parent{},
		&models.Student{},
		&models.Teacher{},
		&models.Admin{},
		&models.Mission{},
		&models.MissionProgress{},
	)
	DB = db
}
