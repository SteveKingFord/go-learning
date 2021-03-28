package model

import (
	"gorm.io/gorm"
	"time"
)


type BaseModel struct {
	Id        string `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
