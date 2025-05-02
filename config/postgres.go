package config

import (
	"fmt"
	"os"
	"strconv"

	apperr "github.com/mzfarshad/music_store_api/pkg/appErr"
)

type postgres struct {
	Port int
	Host string
	Name string
	User string
	Pass string
}

func fromEnv() (*postgres, error) {
	p := new(postgres)
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return nil,
			apperr.NewAppErr(apperr.StatusInternalServerError,
				"failed convert DB_PORT to int",
				apperr.TypeConfig,
				err.Error())
	}
	p.Port = port
	p.Host = os.Getenv("DB_HOST")
	p.Name = os.Getenv("DB_NAME")
	p.User = os.Getenv("DB_USER")
	p.Pass = os.Getenv("DB_PASS")

	if p.Host == "" {
		return nil, apperr.NewAppErr(
			apperr.StatusBadRequest,
			"DB_HOST must not be empty",
			apperr.TypeConfig, "check .env file")
	}
	if p.Name == "" {
		return nil, apperr.NewAppErr(
			apperr.StatusBadRequest,
			"DB_NAME must not be empty",
			apperr.TypeConfig, "check .env file")
	}
	if p.User == "" {
		return nil, apperr.NewAppErr(
			apperr.StatusBadRequest,
			"DB_USER must not be empty",
			apperr.TypeConfig, "check .env file")
	}
	if p.Pass == "" {
		return nil, apperr.NewAppErr(
			apperr.StatusBadRequest,
			"DB_PASS must not be empty",
			apperr.TypeConfig, "check .env file")
	}
	if p.Port == 0 {
		return nil, apperr.NewAppErr(
			apperr.StatusBadRequest,
			"DB_PORT must not be  zero",
			apperr.TypeConfig, "check .env file")
	}
	return p, nil
}

func (p *postgres) DSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		p.Host, p.User, p.Pass, p.Name, p.Port)
}
