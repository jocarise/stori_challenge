package migrations

import (
	"log"
	"newsletter-service/src/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	// List of all models to migrate
	modelsList := []interface{}{
		&models.Newsletter{},
		&models.Category{},
		&models.Recipient{},
	}

	for _, model := range modelsList {
		err := db.AutoMigrate(model)
		if err != nil {
			log.Fatalf("Error executing migrations: %v", err)
			return err
		}
	}

	return nil
}
