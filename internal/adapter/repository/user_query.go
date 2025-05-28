package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/mzfarshad/music_store_api/internal/domain/user"
	"github.com/mzfarshad/music_store_api/pkg/dto"
	"github.com/mzfarshad/music_store_api/pkg/errs"
	"github.com/mzfarshad/music_store_api/pkg/search"
)

func (r *userRepo) First(ctx context.Context, where user.Where) (*user.Entity, error) {
	query := r.db.WithContext(ctx)
	if where.Id != 0 {
		query = query.Where("id = ?", where.Id)
	}
	if where.Type != "" {
		query = query.Where("type = ?", where.Type)
	}
	if where.Email != "" {
		query = query.Where("email = ?", where.Email)
	}
	var model User
	if err := query.First(&model).Error; err != nil {
		return nil, errs.Handle(err, gormErrHandler(model.Type.String()))
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

func (r *userRepo) Search(ctx context.Context, p *search.Pagination[user.SearchParams]) ([]*user.Entity, error) {
	if p == nil {
		return nil, errors.New("trying to search in users with nil pagination")
	}
	query := r.db.WithContext(ctx).Model(&User{})
	if p.Query.Name != "" {
		query = query.Where("name ILIKE ?", fmt.Sprintf("%%%s%%", p.Query.Name))
	}
	if p.Query.Email != "" {
		query = query.Where("email ILIKE ?", fmt.Sprintf("%%%s%%", p.Query.Email))
	}
	if p.Query.Type != "" {
		query = query.Where("type = ?", p.Query.Type)
	}
	var count int64
	if err := query.Count(&count).Error; err != nil {
		return nil, errs.Handle(err, gormErrHandler("user"))
	}
	p.WithTotal(count)
	var models []*User
	if err := query.Limit(p.Limit()).Offset(p.Offset()).Find(&models).Error; err != nil {
		return nil, errs.Handle(err, gormErrHandler("user"))
	}
	return dto.List(models, mapUserToEntity), nil
}
