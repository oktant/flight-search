package db

import "gorm.io/gorm"

type Config interface {
	InitializeDatabase(dbEnv interface{}) (*gorm.DB, error)
	GetDB() *gorm.DB
}
