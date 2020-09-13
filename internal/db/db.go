package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func init() {
	sqlite, err := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	db = sqlite
	db.AutoMigrate(User{}, Post{}, Comment{})
}

func TruncateAllTable() {
	for _, t := range []interface{}{User{}, Post{}, Comment{}} {
		db.Session(&gorm.Session{AllowGlobalUpdate: true}).Unscoped().Delete(t)
	}
}
