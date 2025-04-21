package conf

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type postgres struct {
	Port int
	Host string
	Name string
	User string
	Pass string
}

func FromEnv() (*postgres, error) {
	p := new(postgres)
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return nil, fmt.Errorf("failed to get port from .env file: %v", err)
	}
	p.Port = port
	p.Host = os.Getenv("DB_HOST")
	p.Name = os.Getenv("DB_NAME")
	p.User = os.Getenv("DB_USER")
	p.Pass = os.Getenv("DB_PASS")

	if p.Host == "" {
		return nil, errors.New("config err: DB_HOST must not be  empty.")
	}
	if p.Name == "" {
		return nil, errors.New("config err: DB_NAME must not be  empty.")
	}
	if p.User == "" {
		return nil, errors.New("config err: DB_USER must not be  empty.")
	}
	if p.Pass == "" {
		return nil, errors.New("config err: DB_PASS must not be  empty.")
	}
	if p.Port == 0 {
		return nil, errors.New("config err: DB_PORT must not be  zero.")
	}
	return p, nil
}

func (p *postgres) DSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		p.Host, p.User, p.Pass, p.Name, p.Port)
}
