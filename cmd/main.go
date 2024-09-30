package main

import (
	"github.com/Axel791/contact_api/internal/handlers/v1"
	"github.com/Axel791/contact_api/internal/models"
	userRepo "github.com/Axel791/contact_api/internal/repositories/user"
	"github.com/go-chi/chi/v5"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
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

	// Repositories
	userRepository := userRepo.NewUserRepository(db)

	// Handlers
	userListHandler := v1.NewUserGetListHandler(userRepository)
	userDetailHandler := v1.NewUserDetailHandler(userRepository)

	router := chi.NewRouter()
	router.Route("/api/v1", func(router chi.Router) {
		router.Get("/users", userListHandler.GetUsers)
		router.Get("/users/{userId}", userDetailHandler.GetUserDetail)
	},
	)

	log.Println("Server is running on port 8080...")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		return
	}
}
