package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error
	// Initializer DB SQLite
	db, err = InitDb()
	if err != nil {
		return fmt.Errorf("error initializer sqlite %v", err)
	}
	return nil
}

func GetSQLite() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	// Initializer Logger
	logger = NewLogger(p)
	return logger
}
