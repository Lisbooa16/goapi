package config

import (
	"fmt"

	"github.com/Lisbooa16/goapi/schemas"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DB_USERNAME = "root"
const DB_PASSWORD = "test"
const DB_NAME = "golang"
const DB_HOST = "127.0.0.1"
const DB_PORT = "3306"

var Db *gorm.DB

func InitDb() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	var err error
	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "parseTime=true&loc=Local"
	fmt.Println("dsn : ", dsn)

	Db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Errorf("sqlite opening error: %v", err)
		return nil, err
	}
	err = Db.AutoMigrate(&schemas.Opening{})

	if err != nil {
		logger.Errorf("sqlite automigration error: %v", err)
		return nil, err
	}

	return Db, nil
}
