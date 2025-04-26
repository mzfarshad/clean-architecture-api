package repository

import "gorm.io/gorm"

func createEnum(db *gorm.DB, values ...string) error {
	// TODO: implement me
	return db.Exec(`?`, values).Error
}
