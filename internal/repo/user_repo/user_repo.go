package userrepo

import (
	"context"
	"errors"
	"fmt"

	"github.com/mzfarshad/music_store_api/internal/api/presenter"
	"github.com/mzfarshad/music_store_api/internal/models"
	apperr "github.com/mzfarshad/music_store_api/pkg/appErr"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var ErrNotFoundEmail = errors.New("not found email")

type AuthUserRepository interface {
	FindEmail(ctx context.Context, email string) (models.User, error)
	SaveUser(ctx context.Context, user presenter.SignUpUser) error
}

type authUserRepo struct {
	db *gorm.DB
}

func NewAuthUserRepo(db *gorm.DB) AuthUserRepository {
	return &authUserRepo{db: db}
}

func (a *authUserRepo) FindEmail(ctx context.Context, email string) (models.User, error) {
	var user models.User
	result := a.db.WithContext(ctx).Where("email = ?", email).Debug().First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			customeErr := apperr.NewAppErr(
				apperr.StatusNotFound,
				ErrNotFoundEmail.Error(),
				apperr.TypeDatabase,
				result.Error.Error(),
			)
			return models.User{}, customeErr
		}
		return models.User{}, result.Error
	}
	return user, nil
}

func (a *authUserRepo) SaveUser(ctx context.Context, user presenter.SignUpUser) error {
	newUser := new(models.User)
	pass, err := hashPass(user.Password)
	if err != nil {
		return err
	}
	newUser.Email = user.Email
	newUser.Password = pass
	newUser.Name = user.Name
	result := a.db.WithContext(ctx).Debug().Save(&newUser)
	if result.Error != nil {
		customErr := apperr.NewAppErr(
			apperr.StatusInternalServerError,
			fmt.Sprintf("failed save user by email: %s", user.Email),
			apperr.TypeDatabase,
			result.Error.Error(),
		)
		return customErr
	}
	return nil
}

func hashPass(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), 16)
	if err != nil {
		customeErr := apperr.NewAppErr(
			apperr.StatusInternalServerError,
			"failed hashing password, ",
			apperr.TypeInternal,
			err.Error(),
		)
		return "", customeErr
	}
	return string(hash), nil
}
