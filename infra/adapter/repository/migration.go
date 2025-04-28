package repository

import (
	"fmt"

	"github.com/mzfarshad/music_store_api/infra/domain/user"
	"gorm.io/gorm"
)

func createEnum(db *gorm.DB, values ...string) error {
	var exists bool
	enumName := "user_type"
	checkSql := `
		SELECT EXISTS (
			SELECT 1 FROM pg_type WHERE typename = ?
		)
	`
	if err := db.Raw(checkSql, enumName).Scan(&exists).Error; err != nil {
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
	creatSql := fmt.Sprintf("CREATE TYPE %s AS ENUM (%s)", enumName, valueList)
	return db.Exec(creatSql).Error
}

func CreateUserEnum(db *gorm.DB) error {
	return createEnum(db, string(user.TypeAdmin), string(user.TypeCustomer))
}
