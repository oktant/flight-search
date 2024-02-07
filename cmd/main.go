package main

import (
	"flight-search/db"
	"flight-search/handlers"
	"flight-search/middleware"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	netHttp "net/http"
)

var (
	e       *echo.Echo
	baseUrl = "/api/v1"
)

func main() {
	db.InitializeDb()
	e := echo.New()
	generalGroup := e.Group(baseUrl)
	e.Use(echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{netHttp.MethodGet, netHttp.MethodHead, netHttp.MethodPut, netHttp.MethodPatch, netHttp.MethodPost, netHttp.MethodDelete},
	}))
	generalGroup.Use(middleware.ValidateRequestHeaderMiddleware)
	flightsGroup := generalGroup.Group("/flights")
	flightsGroup.GET("", handlers.GetFlights)
	flightsGroup.POST("", handlers.CreateFlight)

	bookingsGroup := generalGroup.Group("/bookings")
	bookingsGroup.GET("", handlers.GetBookings)
	bookingsGroup.GET("/:id", handlers.GetBooking)
	bookingsGroup.POST("", handlers.CreateBooking)

	e.Logger.Fatal(e.Start(":8080"))
}
