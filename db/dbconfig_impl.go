package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/gommon/log"
	"github.com/srinivasvinay/flight-search/models"
)

var ConfigInstance Config
var DB *gorm.DB

type Database struct {
	*gorm.DB
}
func SetDbConfig(inst Config) {
	ConfigInstance = inst
}
func GetDbConfig() Config {
	return ConfigInstance
}

func (d *Database) InitializeDatabase(dbEnv interface{}) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", dbEnv)
	if err != nil {
		log.Errorf("db err: %+v", err)
		return nil, err
	}
	db.DB().SetMaxIdleConns(20)
	db.AutoMigrate(models.Flight{})
	db.AutoMigrate(models.Booking{})
	db.Model(&models.Booking{}).AddForeignKey("flight_id", "flights(id)", "CASCADE", "CASCADE")
	DB = db
	return DB, nil
}

func (d *Database) GetDB() *gorm.DB {
	if DB == nil {
		connection := models.DbConnection{}
		dbEnv := connection.New().GetConnection()
		_, err := d.InitializeDatabase(dbEnv)
		if err != nil {
			log.Errorf("db err: %+v", err)
			return nil
		}
	}
	return DB
}