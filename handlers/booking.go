package handlers

import (
	"flight-search/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetBookings(c echo.Context) (err error) {
	return c.JSON(http.StatusOK, da.GetBookings())
}

func CreateBooking(c echo.Context) (err error) {
	booking := new(models.Booking)
	if err = c.Bind(booking); err != nil {
		return echo.NewHTTPError(http.StatusOK, "can't bind")
	}
	return c.JSON(http.StatusOK, da.CreateBooking(booking))
}

func GetBooking(c echo.Context) (err error) {
	bookId := c.Param("id")
	return c.JSON(http.StatusOK, da.GetBooking(bookId))
}
