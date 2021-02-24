package model

import (
	"gin-vue-admin/global"
)

type JwtBlacklist struct {
	global.GvaModel
	Jwt string `gorm:"type:text;comment:jwt"`
}
