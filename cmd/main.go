package main

import (
	"github.com/Axel791/contact_api/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"

	"github.com/Axel791/contact_api/internal/handlers/v1"
	userRepo "github.com/Axel791/contact_api/internal/repositories/user"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")

	if dsn == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = db.AutoMigrate(&models.User{}, &models.Contact{})
	if err != nil {
		log.Fatal("Failed to migrate database schemas:", err)
	}

	userRepository := userRepo.NewUserRepository(db)
	userHandler := v1.NewUserDetailGetHandler(userRepository)
}
