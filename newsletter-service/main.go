package main

import (
	"log"
	"net/http"
	"newsletter-service/src/handlers"
	"newsletter-service/src/middlewares"
	"newsletter-service/src/migrations"
	"newsletter-service/src/repositories"
	"newsletter-service/src/services"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")
	SECRET := []byte(os.Getenv("JWT_SECRET"))
	API_VERSION := os.Getenv("API_VERSION")
	DSN := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	if err := migrations.Migrate(db); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	repo := repositories.NewGORMRepository(db)

	newsletterService := services.NewNewsletterService(repo)
	newsletterHandler := handlers.NewNewsletterHandler(newsletterService)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middlewares.CORSMiddleware)

	r.Get(API_VERSION+"/newsletters/unsuscribe", newsletterHandler.UnsuscribeRecipient)
	r.Group(func(r chi.Router) {
		r.Use(middlewares.JWTMiddleware(SECRET))

		r.Get(API_VERSION+"/newsletters/categories", newsletterHandler.GetCategories)
		r.Get(API_VERSION+"/newsletters", newsletterHandler.GetNewsletters)
		r.Post(API_VERSION+"/newsletters", newsletterHandler.CreateNewsletter)
		r.Post(API_VERSION+"/newsletters/{newsletterId}/send", newsletterHandler.SendNewsletters)
		r.Put(API_VERSION+"/newsletters/{newsletterId}", newsletterHandler.AddRecipient)

		//TODO: Stadistics
	})

	http.ListenAndServe(PORT, r)
}
