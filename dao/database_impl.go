package dao

import (
	"flight-search/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormDatabase struct {
	Gorm *gorm.DB
}

func (database *GormDatabase) GetAllFlights() *[]models.Flight {
	flights := new([]models.Flight)
	database.Gorm.Find(flights)
	return flights
}

func (database *GormDatabase) CreateNewFlight(flight *models.Flight) *models.Flight {
	flight.Id = uuid.New()
	database.Gorm.Create(flight)
	return flight
}

func (database *GormDatabase) CreateBooking(booking *models.Booking) *models.Booking {
	booking.Id = uuid.New()
	database.Gorm.Create(booking)
	return booking
}

func (database *GormDatabase) GetBookings() *[]models.Booking {
	booking := new([]models.Booking)
	database.Gorm.Find(booking)
	return booking
}

func (database *GormDatabase) GetBooking(bookId string) *models.Booking {
	booking := new(models.Booking)
	flight := new(models.Flight)
	database.Gorm.Where("id = ?", bookId).Find(booking)
	booking.Flight = *flight
	return booking
}

func (database *GormDatabase) GetFlightsBySourceAndDestination(source, destination string) *[]models.Flight {
	flights := new([]models.Flight)
	database.Gorm.Where("source=? AND destination=?", source, destination).Find(flights)
	return flights
}
