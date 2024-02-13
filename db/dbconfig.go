package db

import (
	"bytes"
	"flight-search/models"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func InitializeDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(connectionString()), &gorm.Config{})
	if err != nil {
		log.Errorf("db err: %+v", err)
		return nil, err
	}
	db.AutoMigrate(models.Flight{})
	db.AutoMigrate(models.Booking{})
	DB = db
	return DB, nil
}

func GetDB() *gorm.DB {
	if DB == nil {
		_, err := InitializeDatabase()
		if err != nil {
			log.Errorf("db err: %+v", err)
			return nil
		}
	}
	return DB
}

func connectionString() string {
	var buffer bytes.Buffer
	buffer.WriteString("host=")
	buffer.WriteString(os.Getenv("DB_HOST"))
	buffer.WriteString(" port=")
	buffer.WriteString(os.Getenv("DB_PORT"))
	buffer.WriteString(" user=")
	buffer.WriteString(os.Getenv("DB_USER"))
	buffer.WriteString(" dbname=")
	buffer.WriteString(os.Getenv("DB_NAME"))
	buffer.WriteString(" password=")
	buffer.WriteString(os.Getenv("DB_PASSWORD"))
	buffer.WriteString(" sslmode=")
	buffer.WriteString(os.Getenv("DB_SSL_MODE"))
	return buffer.String()
}
