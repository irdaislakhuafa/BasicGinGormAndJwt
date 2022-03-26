package database

import (
	"log"

	"github.com/irdaislakhuafa/BasicGinGormAndJwt/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type Connection struct {
	options *utils.OptionsFlags
}

func (connection *Connection) SetOptionParams(optionsParam *utils.OptionsFlags) {
	connection.options = optionsParam
}

func (connection *Connection) Setup() (*gorm.DB, error) {
	url := connection.generatedUrl()
	var err error

	db, err = gorm.Open(mysql.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln("Error :", err.Error())
		return nil, err
	} else {
		log.Printf("Success : database \"%s\" connected\n", connection.options.DbName)
		return db, nil
	}
}
func (connection *Connection) generatedUrl() (url string) {
	option := *connection.options
	url = option.DbUsername + ":" + option.DbPassword + "@tcp(" + option.DbHost + ":" + option.DbPort + ")/" + option.DbName + "?parseTime=True&charset=utf8&loc=Local"
	return url
}

func GetDB() *gorm.DB {
	return db
}
