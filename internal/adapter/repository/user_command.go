package repository

import (
	"context"
	"github.com/mzfarshad/music_store_api/internal/domain"
	"github.com/mzfarshad/music_store_api/internal/domain/user"
	"github.com/mzfarshad/music_store_api/pkg/errs"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm/clause"
)

func (r *userRepo) Create(ctx context.Context, params user.CreateParams) (*user.Entity, error) {
	if err := domain.Validate(params); err != nil {
		return nil, err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(params.Password), 16)
	if err != nil {
		return nil, err
	}
	model := User{
		Name:         params.Name,
		Email:        params.Email,
		PasswordHash: string(hash),
		Type:         params.Type,
	}
	err = r.db.WithContext(ctx).Clauses(clause.Returning{}).Create(&model).Error
	if err != nil {
		return nil, errs.Handle(err, gormErrHandler("user"))
	}
	return mapUserToEntity(&model), nil
}

func (r *userRepo) Update(ctx context.Context, entity *user.Entity) error {
	err := r.db.WithContext(ctx).
		Model(&User{}).
		Where("id = ?", entity.Id).
		Updates(map[string]interface{}{
			"name":            entity.Name,
			"email":           entity.Email,
			"inactive_reason": entity.InactiveReason,
			"status":          entity.Active,
		}).Error
	return err
}
