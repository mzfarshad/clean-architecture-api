package config

import "github.com/spf13/viper"

type postgres struct {
	DSN string `validate:"required"`
}

func (x *postgres) viper() *postgres {
	x.DSN = viper.GetString("DB_DSN")
	return x
}
