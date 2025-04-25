package userrepo

import (
	"context"
	"errors"
	"fmt"

	"github.com/mzfarshad/music_store_api/internal/api/presenter"
	"github.com/mzfarshad/music_store_api/internal/models"
	apperr "github.com/mzfarshad/music_store_api/pkg/appErr"
	"github.com/mzfarshad/music_store_api/pkg/logger"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var ErrNotFoundEmail = errors.New("not found email")

type AuthUserRepository interface {
	FindEmail(ctx context.Context, email string) (*models.User, error)
	SaveUser(ctx context.Context, user *presenter.SignUpUser) error
}

type authUserRepo struct {
	db *gorm.DB
}

func NewAuthUserRepo(db *gorm.DB) AuthUserRepository {
	return &authUserRepo{db: db}
}

func (a *authUserRepo) FindEmail(ctx context.Context, email string) (*models.User, error) {
	log := logger.GetLogger(ctx)
	var user models.User
	err := a.db.WithContext(ctx).Where("email = ?", email).Debug().First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			customErr := apperr.NewAppErr(
				apperr.StatusNotFound,
				ErrNotFoundEmail.Error(),
				apperr.TypeDatabase,
				err.Error(),
			)
			log.Error(ctx, "", customErr)
			return nil, customErr
		}
		customErr := apperr.NewAppErr(
			apperr.StatusInternalServerError,
			fmt.Sprintf("failed find email %s", email),
			apperr.TypeDatabase,
			err.Error(),
		)
		log.Error(ctx, "", err)
		return nil, customErr
	}
	return &user, nil
}

func (a *authUserRepo) SaveUser(ctx context.Context, user *presenter.SignUpUser) error {
	log := logger.GetLogger(ctx)
	newUser := new(models.User)
	pass, err := hashPass(ctx, user.Password)
	if err != nil {
		log.Error(ctx, "", err)
		return err
	}
	newUser.Email = user.Email
	newUser.Password = pass
	newUser.Name = user.Name
	err = a.db.WithContext(ctx).Debug().Save(&newUser).Error
	if err != nil {
		customErr := apperr.NewAppErr(
			apperr.StatusInternalServerError,
			fmt.Sprintf("failed save user by email: %s", user.Email),
			apperr.TypeDatabase,
			err.Error(),
		)
		log.Error(ctx, "", customErr)
		return customErr
	}
	return nil
}

func hashPass(ctx context.Context, pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), 16)
	if err != nil {
		customErr := apperr.NewAppErr(
			apperr.StatusInternalServerError,
			"failed hashing password, ",
			apperr.TypeInternal,
			err.Error(),
		)
		log := logger.GetLogger(ctx)
		log.Error(ctx, "", customErr)
		return "", customErr
	}
	return string(hash), nil
}
