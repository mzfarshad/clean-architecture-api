package conf

import (
	"log"
	"sync"
)

type Config interface {
	Postgres() *postgres
}

type model struct {
	postgres
}

func (m model) Postgres() *postgres {
	return &m.postgres
}

var conf model
var once sync.Once

func Get() Config {
	once.Do(
		func() {
			db, err := FromEnv()
			if err != nil {
				log.Fatal(err)
			}
			conf.postgres = *db
		},
	)
	return conf
}
