package handler

import (
	"github.com/Lisbooa16/goapi/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	Db     *gorm.DB
)

func InitializerHandler() {
	logger = config.GetLogger("handler")
	Db = config.GetSQLite()
}
