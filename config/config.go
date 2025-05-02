package config

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"sync"

	"github.com/spf13/viper"

	"github.com/go-playground/validator/v10"
)

var (
	conf     *Config
	once     = new(sync.Once)
	validate = validator.New()
)

func Get() *Config {
	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			log.Printf("Failed to load env file >>> error: %s", err.Error())
		} else {
			log.Printf("Successfully loaded .env file!")
		}
		viper.AutomaticEnv()
		conf = new(Config)
		if err := viper.Unmarshal(conf); err != nil {
			panic(err)
		}
		conf.viper()
		// Validate config
		if err := conf.Validate(); err != nil {
			panic(err)
		}
		log.Printf("Successfully configured in %q environment...", conf.App.Env)
	})
	return conf
}

type Config struct {
	App      *application `validate:"required"`
	Jwt      *jwt         `validate:"required"`
	Postgres *postgres    `validate:"required"`
}

func (c *Config) Validate() error {
	err := validate.Struct(c)
	if err == nil {
		return nil
	}
	// Check error type
	var validationErrs validator.ValidationErrors
	if !errors.As(err, &validationErrs) {
		// Invalid validation error
		return err
	}
	// This is a valid validation error...
	msg := "Invalid configuration: \n"
	for _, ve := range validationErrs {
		switch ve.ActualTag() {
		case "required":
			msg += fmt.Sprintf("%q is required.\n", ve.Namespace())
		case "oneof":
			msg += fmt.Sprintf("%q should be one of %q, got: %q.\n", ve.Namespace(), ve.Param(), ve.Value())

		default:
			msg += fmt.Sprintf("%q failed on {%s", ve.Namespace(), ve.ActualTag())
			if ve.Param() != "" {
				msg += fmt.Sprintf("=%s} ", ve.Param())
			}
			msg += fmt.Sprintf("tag, got:%v.", ve.Value())
		}
	}
	return errors.New(msg)
}

func (c *Config) viper() {
	conf.App = new(application).viper()
	conf.Jwt = new(jwt).viper()
	conf.Postgres = new(postgres).viper()
}
