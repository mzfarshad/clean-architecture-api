package repository

import (
	"context"
	"fmt"
	"github.com/mzfarshad/music_store_api/pkg/dto"
	"github.com/mzfarshad/music_store_api/pkg/errs"

	"github.com/mzfarshad/music_store_api/internal/domain/user"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func NewUserRepo(db *gorm.DB) user.Repository {
	return &userRepo{
		db: db,
	}
}

type userRepo struct {
	db *gorm.DB
}

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

func (r *userRepo) Create(ctx context.Context, params user.CreateParams) (*user.Entity, error) {
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
	err = r.db.WithContext(ctx).
		Clauses(clause.Returning{}). // TODO: Search in gorm doc if we need this.-- Yes, we need it because we can get the fields that the database fills in itself after creation.
		Create(&model).Error
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
	query = query.Where("user_type = ?", user.TypeCustomer)
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

func (r *userRepo) Update(ctx context.Context, entity *user.Entity) error {
	err := r.db.WithContext(ctx).Model(&User{}).Where("id = ?", entity.Id).Updates(map[string]interface{}{
		"name":            entity.Name,
		"email":           entity.Email,
		"inactive_reason": entity.InactiveReason,
		"status":          entity.Status,
	}).Error
	return err
}
