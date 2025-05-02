package config

import (
	"log"
	"sync"
)

type Config interface {
	Postgres() *postgres
	JWT() *jwt
}

type model struct {
	postgres
	jwt
}

func (m model) Postgres() *postgres {
	return &m.postgres
}

func (m model) JWT() *jwt {
	return &m.jwt
}

var conf model
var once sync.Once

func Get() Config {
	once.Do(
		func() {
			db, err := fromEnv()
			if err != nil {
				log.Fatal(err)
			}
			conf.postgres = *db

			jwt, err := new(jwt).fromEnv()
			if err != nil {
				log.Fatalf(err.Error())
			}
			conf.jwt = *jwt
		},
	)
	return conf
}
