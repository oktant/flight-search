package dao

import "flight-search/models"

type GormAccessDatabase interface {
	GetAllFlights() *[]models.Flight
	CreateNewFlight(results models.Flight) *models.Flight
	GetBookings() *[]models.Booking
	GetBooking(bookId string) *models.Booking
	CreateBooking(booking *models.Booking) *models.Booking
	GetFlightsBySourceAndDestination(source, destination string) *[]models.Flight
}
