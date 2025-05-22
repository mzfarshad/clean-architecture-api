package repository

import (
	"github.com/mzfarshad/music_store_api/internal/domain"
	"gorm.io/gorm"
)

var migrationModels = []any{
	&User{},
}

func Migrate(db *gorm.DB) error {
	if err := beforeMigrate(db); err != nil {
		return err
	}
	if err := db.AutoMigrate(migrationModels...); err != nil {
		return err
	}
	return afterMigrate(db)
}

func beforeMigrate(db *gorm.DB) error {
	if err := createEnum(db, "user_type", string(domain.Admin), string(domain.Customer)); err != nil {
		return err
	}
	return nil
}

func afterMigrate(db *gorm.DB) error {
	return nil
}
