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

func (r *userRepo) Update(ctx context.Context, params user.UpdateParams) (*user.Entity, error) {
	if err := domain.Validate(params); err != nil {
		return nil, err
	}
	updates := make(map[string]any)
	if params.InactiveReason != "" {
		updates["inactive_reason"] = params.InactiveReason
	}
	if params.Active.Populated() {
		updates["active"] = params.Active.Value()
	}

	if len(updates) == 0 {
		return r.First(ctx, user.Where{Id: params.Where.Id})
	}
	var model User
	err := r.db.WithContext(ctx).
		Model(&model).
		Clauses(clause.Returning{}).
		Where("id = ?", params.Where.Id).
		Updates(updates).Error
	if err != nil {
		return nil, errs.Handle(err, gormErrHandler("user"))
	}
	return mapUserToEntity(&model), nil
}
