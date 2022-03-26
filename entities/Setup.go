package entities

import (
	"github.com/irdaislakhuafa/BasicGinGormAndJwt/database"
)

func Setup(list ...interface{}) {
	database.GetDB().AutoMigrate(list...)
}
