package db

import "gorm.io/gorm"

type Config interface {
	InitializeDatabase() (*gorm.DB, error)
}
