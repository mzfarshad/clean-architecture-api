package models

import (
	"github.com/mzfarshad/music_store_api/internal/conf"
	apperr "github.com/mzfarshad/music_store_api/pkg/appErr"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresConnection() (*gorm.DB, error) {
	psql := conf.Get().Postgres()
	db, err := gorm.Open(postgres.Open(psql.DSN()), &gorm.Config{})
	if err != nil {
		err = apperr.NewAppErr(
			apperr.StatusInternalServerError,
			"failed connect to postgres",
			apperr.TypeDatabase,
			err.Error())

		return nil, err
	}
	if err := db.AutoMigrate(
		&User{},
		// ...
	); err != nil {
		err = apperr.NewAppErr(
			apperr.StatusInternalServerError,
			"faield auto migrate, ",
			apperr.TypeDatabase,
			err.Error())
		return nil, err
	}
	return db, nil
}
