package config

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "os"
    "log"
    "github.com/joshua468/template/internal/models"
)

func InitDB() *gorm.DB {
    dsn := os.Getenv("DATABASE_URL")
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    db.AutoMigrate(&models.Product{}, &models.Sales{}, &models.Transaction{}, &models.User{})
    return db
}
