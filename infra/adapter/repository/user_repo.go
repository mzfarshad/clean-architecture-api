package repository

import (
	"context"

	"github.com/mzfarshad/music_store_api/infra/domain/user"
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
		return nil, err
	}
	return mapUserToEntity(&model), nil
}

func (r *userRepo) FirstById(ctx context.Context, id uint) (*user.Entity, error) {
	var model User
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&model).Error
	if err != nil {
		return nil, err
	}
	return mapUserToEntity(&model), nil
}

func (r *userRepo) Create(ctx context.Context, params user.CreateParams) (*user.Entity, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(params.Password), 16)
	if err != nil {
		return nil, err
	}
	model := User{
		Email:        params.Email,
		PasswordHash: string(hash),
		Type:         params.Type,
	}
	err = r.db.WithContext(ctx).
		Clauses(clause.Returning{}). // TODO: Search in gorm doc if we need this.
		Create(&model).Error
	if err != nil {
		return nil, err
	}
	return mapUserToEntity(&model), nil
}

func (r *userRepo) Find(ctx context.Context, params user.SearchParams) ([]*user.Entity, error) {
	return nil, nil
}

func (r *userRepo) Update(ctx context.Context, entity *user.Entity) error {
	err := r.db.Model(&User{}).Where("id = ?", entity.Id).Updates(map[string]interface{}{
		"name":            entity.Name,
		"email":           entity.Email,
		"inactive_reason": entity.InactiveReason,
		"status":          entity.Status,
	}).Error
	return err
}
