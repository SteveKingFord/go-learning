package migrate

import (
	"gorm.io/gorm"
	model3 "kingford-backend/modules/collection/model"
	model2 "kingford-backend/modules/user/model"
)

func InitMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		model2.SysUser{},
		model3.Collection{},
		model3.CollectionItem{},
	)
}