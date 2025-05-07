package repository

import (
	"errors"
	"fmt"
	"github.com/mzfarshad/music_store_api/internal/domain"
	"github.com/mzfarshad/music_store_api/pkg/errs"
	"gorm.io/gorm"
)

func gormErrHandler(subject string) errs.Handler {
	return func(err error) errs.Error {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errs.New(errs.NotFound, subject+" is not found")
		}
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return errs.New(errs.Duplication, subject+" is duplicated")
		}
		if errors.Is(err, gorm.ErrForeignKeyViolated) {
			return errs.New(errs.Unprocessable, "unprocessable "+subject)
		}
		return errs.New(errs.Internal, "something went wrong")
	}
}

func gormModelToDomainEntity(m gorm.Model) domain.Entity {
	return domain.Entity{
		Id:        m.ID,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func createEnum(db *gorm.DB, enumName string, values ...string) error {
	var exists bool
	checkSQL := `
		SELECT EXISTS (
			SELECT 1 FROM pg_type WHERE typname = ?
		)`
	if err := db.Raw(checkSQL, enumName).Scan(&exists).Error; err != nil {
		return err
	}
	if exists {
		return nil
	}

	valueList := ""
	for i, v := range values {
		if i > 0 {
			valueList += ", "
		}
		valueList += fmt.Sprintf("'%s'", v)
	}
	createSQL := fmt.Sprintf("CREATE TYPE %s AS ENUM (%s)", enumName, valueList)
	return db.Exec(createSQL).Error
}
