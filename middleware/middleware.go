package middleware

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func ValidateRequestHeaderMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		headerVal := c.Request().Header.Get("Content-Type")
		if strings.Contains(headerVal, "application/json") {
			return next(c)
		} else {
			return echo.NewHTTPError(http.StatusBadRequest, "error no content type")
		}
	}
}
