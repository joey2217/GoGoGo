package model

import (
	"github.com/jinzhu/gorm"
)

type Role uint8

const (
	Admin Role = iota // 0
	User              // 1
)

type UserDO struct {
	gorm.Model
	Username string
	Password string
	Role     Role
}
