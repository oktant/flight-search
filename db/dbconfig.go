package db

import "github.com/jinzhu/gorm"

type Config interface {
	InitializeDatabase(dbEnv interface{}) (*gorm.DB, error)
	GetDB() *gorm.DB
}
