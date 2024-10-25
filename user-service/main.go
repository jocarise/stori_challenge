package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"user-service/src/handlers"
	"user-service/src/middlewares"
	"user-service/src/migrations"
	"user-service/src/repositories"
	"user-service/src/services"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	API_VERSION := os.Getenv("API_VERSION")
	PORT := os.Getenv("PORT")
	DSN := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	if err := migrations.Migrate(db); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	userRepo := repositories.NewGORMUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middlewares.CORSMiddleware)
	r.Post(API_VERSION+"/users/authenticate", userHandler.AuthenticateUser)
	r.Post(API_VERSION+"/users/register", userHandler.RegisterUser)

	http.ListenAndServe(PORT, r)
}
