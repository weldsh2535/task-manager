package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// InitDB connects to MySQL and returns a gorm.DB instance.
func InitDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
