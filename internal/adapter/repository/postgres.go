package repository

import (
	"github.com/mzfarshad/music_store_api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresConnection() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.Get().Postgres.DSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
