package migrations

import (
	"user-service/src/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	// List of all models to migrate
	modelsList := []interface{}{
		&models.User{},
	}

	for _, model := range modelsList {
		err := db.AutoMigrate(model)
		if err != nil {
			return err
		}
	}

	return nil
}
