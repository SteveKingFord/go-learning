package model

import "gorm.io/gorm"

type  SysUser struct {
	gorm.Model
	NickName string `json:"nick_name" orm:"size(20)"`
}