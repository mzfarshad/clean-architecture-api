package config

import (
	"github.com/spf13/viper"
	"time"
)

type token struct {
	Secret string        `validate:"required"`
	TTL    time.Duration `validate:"gt=0"`
}

// jwt is a struct that holds the configuration for the JWT manager.
type jwt struct {
	Access token
	//Refresh token
	// Admin is a struct that holds the configuration for the JWT manager for the back office.
	//Admin struct {
	//	Access  token
	//	Refresh token
	//}
}

// viper is a function that reads the configuration from the os env variables.
func (x *jwt) viper() *jwt {
	x.Access.Secret = viper.GetString("JWT_ACCESS_SECRET")
	x.Access.TTL = viper.GetDuration("JWT_ACCESS_TTL")
	//x.Refresh.Secret = viper.GetString("JWT_REFRESH_SECRET")
	//x.Refresh.TTL = viper.GetDuration("JWT_REFRESH_TTL")
	//x.Admin.Access.Secret = viper.GetString("JWT_ADMIN_ACCESS_SECRET")
	//x.Admin.Access.TTL = viper.GetDuration("JWT_ADMIN_ACCESS_TTL")
	//x.Admin.Refresh.Secret = viper.GetString("JWT_ADMIN_REFRESH_SECRET")
	//x.Admin.Refresh.TTL = viper.GetDuration("JWT_ADMIN_REFRESH_TTL")
	return x
}
