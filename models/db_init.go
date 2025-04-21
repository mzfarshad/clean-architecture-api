package models

import (
	"github.com/mzfarshad/music_store_api/conf"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresConnection() (*gorm.DB, error) {
	psql := conf.Get().Postgres()
	return gorm.Open(postgres.Open(psql.DSN()), &gorm.Config{})
}
