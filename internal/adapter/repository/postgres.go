package repository

import (
	"github.com/mzfarshad/music_store_api/conf"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresConnection() (*gorm.DB, error) {
	psql := conf.Get().Postgres()
	db, err := gorm.Open(postgres.Open(psql.DSN()), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := createUserEnum(db); err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&User{}); err != nil {
		return nil, err
	}
	if err := db.Exec(`ALTER TABLE users ALTER COLUMN type SET DEFAULT 'customer'`).Error; err != nil {
		return nil, err
	}
	return db, nil
}
