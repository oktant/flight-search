package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"github.com/srinivasvinay/flight-search/db"
	"github.com/srinivasvinay/flight-search/models"
)

type GormDatabase struct {
	Gorm      *gorm.DB
}

func (database *GormDatabase) GetAllFlights() * []models.Flight{
	flights := new([]models.Flight)
	db.GetDbConfig().GetDB().Find(flights)

	return flights
}

func (database *GormDatabase) CreateNewFlight(flight *models.Flight) *models.Flight{
	flight.Id = uuid.NewV4().String()

	db.GetDbConfig().GetDB().Create(flight)

	return flight
}

func (database *GormDatabase) CreateBooking(booking *models.Booking) *models.Booking {
	booking.Id = uuid.NewV4().String()

	db.GetDbConfig().GetDB().Create(booking)

	return booking
}

func (database *GormDatabase) GetBookings() * []models.Booking{
	booking := new([]models.Booking)
	db.GetDbConfig().GetDB().Find(booking)
	return booking
}

func (database *GormDatabase) GetBooking(bookId string) *models.Booking{
	booking := new(models.Booking)
	flight := new(models.Flight)
	db.GetDbConfig().GetDB().Where("id = ?", bookId).Find(booking)
	db.GetDbConfig().GetDB().Model(booking).Related(flight)
	booking.Flight=flight
	return booking
}

func (database *GormDatabase) GetFlightsBySourceAndDestination(source, destination string) *[]models.Flight {
	flights := new([]models.Flight)
	db.GetDbConfig().GetDB().Where("source=? AND destination=?", source, destination).Find(flights)

	return flights
}
