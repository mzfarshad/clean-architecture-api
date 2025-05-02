package repository

import (
	"github.com/mzfarshad/music_store_api/internal/domain"
	"gorm.io/gorm"
)

func gormModelToDomainEntity(m gorm.Model) domain.Entity {
	return domain.Entity{
		Id:        m.ID,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
