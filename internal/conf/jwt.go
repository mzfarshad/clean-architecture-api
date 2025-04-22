package conf

import (
	"os"

	apperr "github.com/mzfarshad/music_store_api/pkg/appErr"
)

type jwt struct {
	SecretKey string
}

func (j *jwt) fromEnv() (*jwt, error) {
	j.SecretKey = os.Getenv("SECRET_KEY")
	if j.SecretKey == "" {
		customErr := apperr.NewAppErr(
			apperr.StatusBadRequest,
			"SECRET_KEY must not be empty!. ",
			apperr.TypeInternal,
			"check .env file",
		)
		return nil, customErr
	}
	return j, nil
}
