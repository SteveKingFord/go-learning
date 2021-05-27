package model

import "gorm.io/gorm"

type  SysUser struct {
	NickName string `json:"nick_name" orm:"size(20)"`
	gorm.Model
}