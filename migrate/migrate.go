package migrate

import (
	"gorm.io/gorm"
	collection "kingford-backend/modules/collection/model"
	user "kingford-backend/modules/user/model"
)

func InitMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		user.SysUser{},
		collection.Collection{},
		collection.CollectionCategory{},
	)
}