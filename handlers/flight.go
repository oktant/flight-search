package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/srinivasvinay/flight-search/dao"
	"github.com/srinivasvinay/flight-search/models"
	"net/http"
)
var da dao.GormDatabase

func GetFlights(c echo.Context) (err error){
	source := c.QueryParam("source")
	destination := c.QueryParam("destination")
	if len(source) >0 && len(destination) >0 {
		return c.JSON(http.StatusOK, da.GetFlightsBySourceAndDestination(source, destination))

	}
	return c.JSON(http.StatusOK, da.GetAllFlights())
}

func CreateFlight(c echo.Context) (err error){
	flight := new (models.Flight)
	if err=c.Bind(flight); err !=nil {
			return echo.NewHTTPError(http.StatusOK, "can't bind")
	}

	return c.JSON(http.StatusOK, da.CreateNewFlight(flight))
}

func GetBookings(c echo.Context) (err error) {

	return c.JSON(http.StatusOK, da.GetBookings())
}

func CreateBooking(c echo.Context) (err error) {
	booking := new (models.Booking)
	if err=c.Bind(booking); err !=nil {
		return echo.NewHTTPError(http.StatusOK, "can't bind")
	}
	return c.JSON(http.StatusOK, da.CreateBooking(booking))
}

func GetBooking(c echo.Context) (err error) {
	bookId := c.Param("id")
	return c.JSON(http.StatusOK, da.GetBooking(bookId))
}
