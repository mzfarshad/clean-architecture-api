package repository

import (
	"context"
	"fmt"
	"github.com/mzfarshad/music_store_api/internal/domain/user"
	"github.com/mzfarshad/music_store_api/pkg/dto"
	"github.com/mzfarshad/music_store_api/pkg/errs"
)

func (r *userRepo) FirstByEmail(ctx context.Context, email string) (*user.Entity, error) {
	var model User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&model).Error
	if err != nil {
		return nil, errs.Handle(err, gormErrHandler("user"))
	}
	return mapUserToEntity(&model), nil
}

func (r *userRepo) FirstById(ctx context.Context, id uint) (*user.Entity, error) {
	var model User
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&model).Error
	if err != nil {
		return nil, errs.Handle(err, gormErrHandler("user"))
	}
	return mapUserToEntity(&model), nil
}

func (r *userRepo) Find(ctx context.Context, params user.SearchParams) (*user.PaginationParams, error) {
	var users []*User
	var totalData int64
	var totalPages int
	query := r.db.WithContext(ctx).Model(&User{})
	if params.Email != "" {
		query = query.Where("email ILIKE ?", fmt.Sprintf("%%%s%%", params.Email))
	}
	if params.Name != "" {
		query = query.Where("name ILIKE ?", fmt.Sprintf("%%%s%%", params.Name))
	}
	query = query.Where("type = ?", user.TypeCustomer)
	if err := query.Count(&totalData).Error; err != nil {
		return nil, errs.Handle(err, gormErrHandler("user"))
	}
	totalPages = int((totalData + int64(params.Limit) - 1) / int64(params.Limit))
	ofst := (params.Page - 1) * params.Limit
	if err := query.Limit(params.Limit).Offset(ofst).Find(&users).Error; err != nil {
		return nil, err
	}
	return &user.PaginationParams{
		TotalData:  int(totalData),
		TotalPages: totalPages,
		Result:     dto.List(users, mapUserToEntity),
	}, nil
}
