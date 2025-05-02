package repository

import (
	"errors"
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

func createEnum(db *gorm.DB, name string, values ...string) error {
	//func createEnum(db *gorm.DB, values ...string) error {
	//	var exists bool
	//	enumName := "user_type"
	//	checkSql := `
	//	SELECT EXISTS (
	//		SELECT 1 FROM pg_type WHERE typname = ?
	//	)`
	//	if err := db.Raw(checkSql, enumName).Scan(&exists).Error; err != nil {
	//	return err
	//}
	//	if exists {
	//	return nil
	//}
	//	valueList := ""
	//	for i, v := range values {
	//	if i > 0 {
	//	valueList += ", "
	//}
	//	valueList += fmt.Sprintf("'%s'", v)
	//}
	//	creatSql := fmt.Sprintf("CREATE TYPE %s AS ENUM (%s)", enumName, valueList)
	//	return db.Exec(creatSql).Error
	//}

	//TODO: implement me
	return errors.New("implement me")
}
