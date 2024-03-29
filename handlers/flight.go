package handlers

import (
	"flight-search/dao"
	"flight-search/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

var da dao.GormDatabase

func GetFlights(c echo.Context) (err error) {
	source := c.QueryParam("source")
	destination := c.QueryParam("destination")
	if len(source) > 0 && len(destination) > 0 {
		return c.JSON(http.StatusOK, da.GetFlightsBySourceAndDestination(source, destination))
	}
	return c.JSON(http.StatusOK, da.GetAllFlights())
}

func CreateFlight(c echo.Context) (err error) {
	flight := new(models.Flight)
	if err = c.Bind(flight); err != nil {
		return echo.NewHTTPError(http.StatusOK, "can't bind")
	}
	return c.JSON(http.StatusOK, da.CreateNewFlight(flight))
}
