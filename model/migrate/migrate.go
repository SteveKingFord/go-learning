package migrate

import (
	"gorm.io/gorm"
	"kingford-backend/model"
)

func InitMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		model.SysUser{},
		model.Collection{},
		model.CollectionItem{},
	)
}