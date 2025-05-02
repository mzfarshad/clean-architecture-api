package config

import (
	"github.com/spf13/viper"
)

const (
	appEnvKey string = "APP_ENV"

	EnvTesting    environment = "testing"
	EnvLocal      environment = "local"
	EnvStaging    environment = "staging"
	EnvQA         environment = "qa"
	EnvProduction environment = "production"
)

type environment string

func (x environment) String() string {
	return string(x)
}

func (x environment) Is(env environment, or ...environment) bool {
	if x == env {
		return true
	}
	for _, v := range or {
		if x == v {
			return true
		}
	}
	return false
}

type application struct {
	Name    string `validate:"required"`
	Host    string
	Port    string      `validate:"required,numeric"`
	Env     environment `validate:"required,oneof=testing local staging qa production"`
	Recover bool
	// Debug controls weather error detail trace sent to client site or not (Default: false).
	Debug bool
}

func (a *application) viper() *application {
	a.Name = viper.GetString("APP_NAME")
	a.Host = viper.GetString("APP_HOST")
	a.Port = viper.GetString("APP_PORT")
	a.Env = environment(viper.GetString(appEnvKey))
	viper.SetDefault("APP_RECOVER", true)
	a.Recover = viper.GetBool("APP_RECOVER")
	a.Debug = viper.GetBool("APP_DEBUG")
	return a
}
