package dao

import (
	"flight-search/db"
	"flight-search/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormDatabase struct {
	Gorm *gorm.DB
}

func (database *GormDatabase) GetAllFlights() *[]models.Flight {
	flights := new([]models.Flight)
	db.GetDbConfig().GetDB().Find(flights)

	return flights
}

func (database *GormDatabase) CreateNewFlight(flight *models.Flight) *models.Flight {
	flight.Id = uuid.New()

	db.GetDbConfig().GetDB().Create(flight)

	return flight
}

func (database *GormDatabase) CreateBooking(booking *models.Booking) *models.Booking {
	booking.Id = uuid.New()

	db.GetDbConfig().GetDB().Create(booking)

	return booking
}

func (database *GormDatabase) GetBookings() *[]models.Booking {
	booking := new([]models.Booking)
	db.GetDbConfig().GetDB().Find(booking)
	return booking
}

func (database *GormDatabase) GetBooking(bookId string) *models.Booking {
	booking := new(models.Booking)
	flight := new(models.Flight)
	db.GetDbConfig().GetDB().Where("id = ?", bookId).Find(booking)

	booking.Flight = *flight
	return booking
}

func (database *GormDatabase) GetFlightsBySourceAndDestination(source, destination string) *[]models.Flight {
	flights := new([]models.Flight)
	db.GetDbConfig().GetDB().Where("source=? AND destination=?", source, destination).Find(flights)

	return flights
}
