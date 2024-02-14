package db

import (
	"flight-search/models"
	"github.com/labstack/gommon/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"strings"
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
	var csb strings.Builder
	csb.WriteString("host=")
	csb.WriteString(os.Getenv("DB_HOST"))
	csb.WriteString(" port=")
	csb.WriteString(os.Getenv("DB_PORT"))
	csb.WriteString(" user=")
	csb.WriteString(os.Getenv("DB_USER"))
	csb.WriteString(" dbname=")
	csb.WriteString(os.Getenv("DB_NAME"))
	csb.WriteString(" password=")
	csb.WriteString(os.Getenv("DB_PASSWORD"))
	csb.WriteString(" sslmode=")
	csb.WriteString(os.Getenv("DB_SSL_MODE"))
	return csb.String()
}
