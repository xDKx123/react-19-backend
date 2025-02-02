package migrations

import (
	"gorm.io/gorm"
	"testProject/src/models"
)

func MigrateModels(db *gorm.DB) (err error) {
	modelsToMigrate := []interface{}{
		&models.User{},
	}

	err = db.AutoMigrate(modelsToMigrate...)

	return err
}
