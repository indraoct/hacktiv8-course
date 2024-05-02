package options

import (
	"gorm.io/gorm"
	"hacktiv8-course/final_assignment/commons/jwt"
)

type Options struct {
	DbGorm *gorm.DB
	JwtGen *jwt.JwtPkg
}
